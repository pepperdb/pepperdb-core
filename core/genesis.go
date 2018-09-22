package core

import (
	"fmt"
	"io/ioutil"

	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"

	"github.com/gogo/protobuf/proto"
	"github.com/pepperdb/pepperdb-core/common/dag"
	"github.com/pepperdb/pepperdb-core/common/util"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/pepperdb/pepperdb-core/consensus/pb"
	"github.com/pepperdb/pepperdb-core/core/pb"
	"github.com/pepperdb/pepperdb-core/core/state"
	"github.com/sirupsen/logrus"
)

// Genesis Block Hash
var (
	GenesisHash        = make([]byte, BlockHashLength)
	GenesisTimestamp   = int64(0)
	GenesisCoinbase, _ = NewAddressFromPublicKey(make([]byte, PublicKeyDataLength))
)

// LoadGenesisConf load genesis conf for file
func LoadGenesisConf(filePath string) (*corepb.Genesis, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to read the genesis config file.")
		return nil, err
	}
	content := string(b)

	genesis := new(corepb.Genesis)
	if err := proto.UnmarshalText(content, genesis); err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err": err,
		}).Error("Failed to parse genesis file.")
		return nil, err
	}

	return genesis, nil
}

// NewGenesisBlock create genesis @Block from file.
func NewGenesisBlock(conf *corepb.Genesis, chain *BlockChain) (*Block, error) {
	if conf == nil || chain == nil {
		return nil, ErrNilArgument
	}

	worldState, err := state.NewWorldState(chain.ConsensusHandler(), chain.storage)
	if err != nil {
		return nil, err
	}
	genesisBlock := &Block{
		header: &BlockHeader{
			hash:          GenesisHash,
			parentHash:    GenesisHash,
			chainID:       conf.Meta.ChainId,
			coinbase:      GenesisCoinbase,
			timestamp:     GenesisTimestamp,
			consensusRoot: &consensuspb.ConsensusRoot{},
			alg:           keystore.SECP256K1,
		},
		transactions: make(Transactions, 0),
		dependency:   dag.NewDag(),
		worldState:   worldState,
		txPool:       chain.txPool,
		storage:      chain.storage,
		eventEmitter: chain.eventEmitter,
		nvm:          chain.nvm,
		height:       1,
		sealed:       false,
	}

	consensusState, err := chain.ConsensusHandler().GenesisConsensusState(chain, conf)
	if err != nil {
		return nil, err
	}
	genesisBlock.worldState.SetConsensusState(consensusState)

	if err := genesisBlock.Begin(); err != nil {
		return nil, err
	}
	// add token distribution for genesis
	for _, v := range conf.TokenDistribution {
		addr, err := AddressParse(v.Address)
		if err != nil {
			logging.CLog().WithFields(logrus.Fields{
				"address": v.Address,
				"err":     err,
			}).Error("Found invalid address in genesis token distribution.")
			genesisBlock.RollBack()
			return nil, err
		}
		acc, err := genesisBlock.worldState.GetOrCreateUserAccount(addr.address)
		if err != nil {
			genesisBlock.RollBack()
			return nil, err
		}
		txsBalance, err := util.NewUint128FromString(v.Value)
		if err != nil {
			genesisBlock.RollBack()
			return nil, err
		}
		err = acc.AddBalance(txsBalance)
		if err != nil {
			genesisBlock.RollBack()
			return nil, err
		}
	}

	// genesis transaction
	declaration := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n\n\n%s",
		"Pepperdb Manifesto",
		"Contract the world..."

		"by Pepperdb (pepperdb.org)",
	)
	declarationTx, err := NewTransaction(
		chain.ChainID(),
		GenesisCoinbase, GenesisCoinbase,
		util.Uint128Zero(), 1,
		TxPayloadBinaryType,
		[]byte(declaration),
		TransactionGasPrice,
		MinGasCountPerTransaction,
	)
	if err != nil {
		return nil, err
	}
	declarationTx.timestamp = 0
	hash, err := declarationTx.calHash()
	if err != nil {
		return nil, err
	}
	declarationTx.hash = hash
	declarationTx.alg = keystore.SECP256K1
	pbTx, err := declarationTx.ToProto()
	if err != nil {
		return nil, err
	}
	txBytes, err := proto.Marshal(pbTx)
	if err != nil {
		return nil, err
	}
	genesisBlock.transactions = append(genesisBlock.transactions, declarationTx)
	if err := genesisBlock.worldState.PutTx(declarationTx.hash, txBytes); err != nil {
		return nil, err
	}

	genesisBlock.Commit()

	genesisBlock.header.stateRoot = genesisBlock.WorldState().AccountsRoot()
	genesisBlock.header.txsRoot = genesisBlock.WorldState().TxsRoot()
	genesisBlock.header.eventsRoot = genesisBlock.WorldState().EventsRoot()
	genesisBlock.header.consensusRoot = genesisBlock.WorldState().ConsensusRoot()

	genesisBlock.sealed = true

	return genesisBlock, nil
}

// CheckGenesisBlock if a block is a genesis block
func CheckGenesisBlock(block *Block) bool {
	if block == nil {
		return false
	}
	if block.Hash().Equals(GenesisHash) {
		return true
	}
	return false
}

// CheckGenesisTransaction if a tx is a genesis transaction
func CheckGenesisTransaction(tx *Transaction) bool {
	if tx == nil {
		return false
	}
	if tx.from.Equals(GenesisCoinbase) {
		return true
	}
	return false
}

// DumpGenesis return the configuration of the genesis block in the storage
func DumpGenesis(chain *BlockChain) (*corepb.Genesis, error) {
	genesis, err := LoadBlockFromStorage(GenesisHash, chain)
	if err != nil {
		return nil, err
	}
	dynasty, err := genesis.worldState.Dynasty()
	if err != nil {
		return nil, err
	}
	bootstrap := []string{}
	for _, v := range dynasty {
		addr, err := AddressParseFromBytes(v)
		if err != nil {
			return nil, err
		}
		bootstrap = append(bootstrap, addr.String())
	}
	distribution := []*corepb.GenesisTokenDistribution{}
	accounts, err := genesis.worldState.Accounts()
	if err != nil {
		return nil, err
	}
	for _, v := range accounts {
		balance := v.Balance()
		if v.Address().Equals(genesis.Coinbase().Bytes()) {
			continue
		}
		addr, err := AddressParseFromBytes(v.Address())
		if err != nil {
			return nil, err
		}
		distribution = append(distribution, &corepb.GenesisTokenDistribution{
			Address: addr.String(),
			Value:   balance.String(),
		})
	}
	return &corepb.Genesis{
		Meta: &corepb.GenesisMeta{ChainId: genesis.ChainID()},
		Consensus: &corepb.GenesisConsensus{
			Dpos: &corepb.GenesisConsensusDpos{Dynasty: bootstrap},
		},
		TokenDistribution: distribution,
	}, nil
}

//CheckGenesisConfByDB check mem and genesis.conf if equal return nil
func CheckGenesisConfByDB(pGenesisDB *corepb.Genesis, pGenesis *corepb.Genesis) error {
	//private function [Empty parameters are checked by the caller]
	if pGenesisDB != nil {
		if pGenesis.Meta.ChainId != pGenesisDB.Meta.ChainId {
			return ErrGenesisNotEqualChainIDInDB
		}

		if len(pGenesis.Consensus.Dpos.Dynasty) != len(pGenesisDB.Consensus.Dpos.Dynasty) {
			return ErrGenesisNotEqualDynastyLenInDB
		}

		if len(pGenesis.TokenDistribution) != len(pGenesisDB.TokenDistribution) {
			return ErrGenesisNotEqualTokenLenInDB
		}

		// check dpos equal
		for _, confDposAddr := range pGenesis.Consensus.Dpos.Dynasty {
			contains := false
			for _, dposAddr := range pGenesisDB.Consensus.Dpos.Dynasty {
				if dposAddr == confDposAddr {
					contains = true
					break
				}
			}
			if !contains {
				return ErrGenesisNotEqualDynastyInDB
			}

		}

		// check distribution equal
		for _, confDistribution := range pGenesis.TokenDistribution {
			contains := false
			for _, distribution := range pGenesisDB.TokenDistribution {
				if distribution.Address == confDistribution.Address &&
					distribution.Value == confDistribution.Value {
					contains = true
					break
				}
			}
			if !contains {
				return ErrGenesisNotEqualTokenInDB
			}
		}
	}
	return nil
}

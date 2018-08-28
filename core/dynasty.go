package core

import (
	"io/ioutil"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/pepperdb/pepperdb-core/common/trie"
	"github.com/pepperdb/pepperdb-core/core/pb"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

var (
	// DynastyConf define dynasty addresses
	DynastyConf *corepb.Dynasty
	// DynastyTrie ..
	DynastyTrie *trie.Trie
	// GenesisDynastyTrie dynastyTrie of genesis block
	GenesisDynastyTrie *trie.Trie
	// GenesisRealTimestamp ..
	GenesisRealTimestamp int64
	// InitialDynastyKeepTime ..
	InitialDynastyKeepTime int64
	once                   sync.Once
)

// LoadDynastyConf ..
func LoadDynastyConf(filePath string, genesis *corepb.Genesis) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err":      err,
			"filePath": filePath,
		}).Fatal("File doesn't exist.")
	}
	content := string(b)

	DynastyConf = new(corepb.Dynasty)
	if err = proto.UnmarshalText(content, DynastyConf); err != nil {
		logging.CLog().WithFields(logrus.Fields{
			"err":      err,
			"filePath": filePath,
		}).Fatal("Failed to parse dynasty.conf.")
	}

	if DynastyConf.Meta.ChainId != genesis.Meta.ChainId || len(DynastyConf.Candidate) != 1 {
		logging.CLog().WithFields(logrus.Fields{
			"GenesisChainId":      genesis.Meta.ChainId,
			"DynastyChainId":      DynastyConf.Meta.ChainId,
			"DynastyCandidateLen": len(DynastyConf.Candidate),
		}).Fatal("ChainId in dynasty.conf differs from that in genesis.conf.")
	}

	if len(DynastyConf.Candidate[0].Dynasty) != len(genesis.Consensus.Dpos.Dynasty) {
		logging.CLog().WithFields(logrus.Fields{
			"DynastySize": len(DynastyConf.Candidate[0].Dynasty),
		}).Fatal("Miners count in dynasty.conf differs from that in genesis.conf.")
	}
}

// InitDynastyFromConf Fatal when initialization failed
func InitDynastyFromConf(chain *BlockChain, BlockIntervalInSecond, DynastyIntervalInSecond int64) {
	once.Do(func() {
		d, err := trie.NewTrie(nil, chain.Storage(), false)
		if err != nil {
			logging.VLog().WithFields(logrus.Fields{
				"err": err,
			}).Fatal("Failed to new trie.")
		}

		candidate := DynastyConf.Candidate[0]
		for i := 0; i < len(candidate.Dynasty); i++ {
			addr := candidate.Dynasty[i]
			miner, err := AddressParse(addr)
			if err != nil {
				logging.VLog().WithFields(logrus.Fields{
					"err": err,
				}).Fatal("Failed to parse address.")
			}
			v := miner.Bytes()
			if _, err = d.Put(v, v); err != nil {
				logging.VLog().WithFields(logrus.Fields{
					"err": err,
				}).Fatal("Failed to put value.")
			}
		}

		DynastyTrie = d

		InitialDynastyKeepTime = int64(candidate.Serial) * DynastyIntervalInSecond

		logging.VLog().WithFields(logrus.Fields{
			"chainId":                DynastyConf.Meta.ChainId,
			"serial":                 DynastyConf.Candidate[0].Serial,
			"InitialDynastyKeepTime": InitialDynastyKeepTime,
		}).Debug("Init dynasty.conf done.")
	})

	if GenesisRealTimestamp == 0 {
		b := chain.GetBlockOnCanonicalChainByHeight(2)
		if b == nil {
			logging.VLog().Debug("Nil block found at height 2.")
			return
		}
		GenesisRealTimestamp = b.Timestamp() - BlockIntervalInSecond
		logging.VLog().WithFields(logrus.Fields{
			"GenesisRealTimestamp": GenesisRealTimestamp,
		}).Debug("Init GenesisRealTimestamp.")
	}
}

package nvm

import (
	"github.com/gogo/protobuf/proto"
	"github.com/pepperdb/pepperdb-core/core"
	"github.com/pepperdb/pepperdb-core/core/pb"
)

// SerializableAccount serializable account state
type SerializableAccount struct {
	Nonce   uint64 `json:"nonce"`
	Balance string `json:"balance"`
}

// SerializableBlock serializable block
type SerializableBlock struct {
	Timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
	Height    uint64 `json:"height"`
	Seed      string `json:"seed,omitempty"`
}

// SerializableTransaction serializable transaction
type SerializableTransaction struct {
	Hash      string `json:"hash"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Nonce     uint64 `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	GasPrice  string `json:"gasPrice"`
	GasLimit  string `json:"gasLimit"`
}

// Context nvm engine context
type Context struct {
	block    Block
	tx       Transaction
	contract Account
	state    WorldState
}

// NewContext create a engine context
func NewContext(block Block, tx Transaction, contract Account, state WorldState) (*Context, error) {
	if block == nil || tx == nil || contract == nil || state == nil {
		return nil, ErrContextConstructArrEmpty
	}
	ctx := &Context{
		block:    block,
		tx:       tx,
		contract: contract,
		state:    state,
	}
	return ctx, nil
}

func toSerializableAccount(acc Account) *SerializableAccount {
	sAcc := &SerializableAccount{
		Nonce:   acc.Nonce(),
		Balance: acc.Balance().String(),
	}
	return sAcc
}

func toSerializableBlock(block Block) *SerializableBlock {
	sBlock := &SerializableBlock{
		Timestamp: block.Timestamp(),
		Hash:      "",
		Height:    block.Height(),
	}
	if block.RandomAvailable() {
		sBlock.Seed = block.RandomSeed()
	}
	return sBlock
}

func toSerializableTransaction(tx Transaction) *SerializableTransaction {
	return &SerializableTransaction{
		From:      tx.From().String(),
		To:        tx.To().String(),
		Value:     tx.Value().String(),
		Timestamp: tx.Timestamp(),
		Nonce:     tx.Nonce(),
		Hash:      tx.Hash().String(),
		GasPrice:  tx.GasPrice().String(),
		GasLimit:  tx.GasLimit().String(),
	}
}

func toSerializableTransactionFromBytes(txBytes []byte) (*SerializableTransaction, error) {
	pbTx := new(corepb.Transaction)
	if err := proto.Unmarshal(txBytes, pbTx); err != nil {
		return nil, err
	}
	tx := new(core.Transaction)
	if err := tx.FromProto(pbTx); err != nil {
		return nil, err
	}

	return &SerializableTransaction{
		From:      tx.From().String(),
		To:        tx.To().String(),
		Value:     tx.Value().String(),
		Timestamp: tx.Timestamp(),
		Nonce:     tx.Nonce(),
		Hash:      tx.Hash().String(),
		GasPrice:  tx.GasPrice().String(),
		GasLimit:  tx.GasLimit().String(),
	}, nil
}

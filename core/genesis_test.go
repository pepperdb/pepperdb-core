package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenesisBlock(t *testing.T) {
	neb := testNeb(t)
	chain := neb.chain
	genesis := neb.chain.genesisBlock
	conf := MockGenesisConf()

	for _, v := range conf.TokenDistribution {
		addr, _ := AddressParse(v.Address)
		acc, err := genesis.worldState.GetOrCreateUserAccount(addr.Bytes())
		assert.Nil(t, err)
		assert.Equal(t, acc.Balance().String(), v.Value)
	}

	dumpConf, err := DumpGenesis(chain)
	assert.Nil(t, err)
	assert.Equal(t, dumpConf.Meta.ChainId, conf.Meta.ChainId)
	assert.Equal(t, dumpConf.TokenDistribution, conf.TokenDistribution)
}

func TestInvalidAddressInTokenDistribution(t *testing.T) {
	mockConf := MockGenesisConf()
	mockConf.TokenDistribution[0].Address = "n1UZtMgi94oE913L2Sa2C9XwvAzNTQ82v64121"
	chain := testNeb(t).chain
	_, err := NewGenesisBlock(mockConf, chain)
	assert.Equal(t, err, ErrInvalidAddressFormat)
}

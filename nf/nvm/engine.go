package nvm

import (
	"github.com/pepperdb/pepperdb-core/core"
	"github.com/pepperdb/pepperdb-core/core/state"
)

// NebulasVM type of NebulasVM
type NebulasVM struct{}

// NewNebulasVM create new NebulasVM
func NewNebulasVM() core.NVM {
	return &NebulasVM{}
}

// CreateEngine start engine
func (nvm *NebulasVM) CreateEngine(block *core.Block, tx *core.Transaction, contract state.Account, state core.WorldState) (core.SmartContractEngine, error) {
	ctx, err := NewContext(block, tx, contract, state)
	if err != nil {
		return nil, err
	}
	return NewV8Engine(ctx), nil
}

// CheckV8Run to check V8 env is OK
func (nvm *NebulasVM) CheckV8Run() error {
	engine := NewV8Engine(&Context{
		block:    core.MockBlock(nil, 1),
		contract: state.MockAccount("1.0.0"),
		tx:       nil,
		state:    nil,
	})
	_, err := engine.RunScriptSource("", 0)
	engine.Dispose()
	return err
}

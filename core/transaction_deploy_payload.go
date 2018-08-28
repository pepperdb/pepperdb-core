package core

import (
	"encoding/json"
	"fmt"

	"github.com/pepperdb/pepperdb-core/core/pb"
	"github.com/pepperdb/pepperdb-core/core/state"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"

	"github.com/pepperdb/pepperdb-core/common/util"
)

// DeployPayload carry contract deploy information
type DeployPayload struct {
	SourceType string
	Source     string
	Args       string
}

// CheckContractArgs check contract args
func CheckContractArgs(args string) error {
	if len(args) > 0 {
		var argsObj []interface{}
		if err := json.Unmarshal([]byte(args), &argsObj); err != nil {
			return err
		}
	}
	return nil
}

// LoadDeployPayload from bytes
func LoadDeployPayload(bytes []byte) (*DeployPayload, error) {
	payload := &DeployPayload{}
	if err := json.Unmarshal(bytes, payload); err != nil {
		return nil, ErrInvalidArgument
	}
	return NewDeployPayload(payload.Source, payload.SourceType, payload.Args)
}

// NewDeployPayload with source & args
func NewDeployPayload(source, sourceType, args string) (*DeployPayload, error) {
	if len(source) == 0 {
		return nil, ErrInvalidDeploySource
	}

	if sourceType != SourceTypeTypeScript && sourceType != SourceTypeJavaScript {
		return nil, ErrInvalidDeploySourceType
	}

	if err := CheckContractArgs(args); err != nil {
		return nil, ErrInvalidArgument
	}

	return &DeployPayload{
		Source:     source,
		SourceType: sourceType,
		Args:       args,
	}, nil
}

// ToBytes serialize payload
func (payload *DeployPayload) ToBytes() ([]byte, error) {
	return json.Marshal(payload)
}

// BaseGasCount returns base gas count
func (payload *DeployPayload) BaseGasCount() *util.Uint128 {
	base, _ := util.NewUint128FromInt(60)
	return base
}

// Execute deploy payload in tx, deploy a new contract
func (payload *DeployPayload) Execute(limitedGas *util.Uint128, tx *Transaction, block *Block, ws WorldState) (*util.Uint128, string, error) {
	if block == nil || tx == nil {
		return util.NewUint128(), "", ErrNilArgument
	}

	if !tx.From().Equals(tx.To()) {
		return util.NewUint128(), "", ErrContractTransactionAddressNotEqual
	}

	// payloadGasLimit <= 0, v8 engine not limit the execution instructions
	if limitedGas.Cmp(util.NewUint128()) <= 0 {
		return util.NewUint128(), "", ErrOutOfGasLimit
	}

	addr, err := tx.GenerateContractAddress()
	if err != nil {
		return util.NewUint128(), "", err
	}
	/* // disable useless owner.
	owner, err := ws.GetOrCreateUserAccount(tx.from.Bytes())
	if err != nil {
		return util.NewUint128(), "", err
	} */
	var contract state.Account
	v := GetMaxV8JSLibVersionAtHeight(block.Height())
	if len(v) > 0 {
		contract, err = ws.CreateContractAccount(addr.Bytes(), tx.Hash(), &corepb.ContractMeta{Version: v})
	} else {
		contract, err = ws.CreateContractAccount(addr.Bytes(), tx.Hash(), nil)
	}
	if err != nil {
		return util.NewUint128(), "", err
	}

	engine, err := block.nvm.CreateEngine(block, tx, contract, ws)
	if err != nil {
		return util.NewUint128(), "", err
	}
	defer engine.Dispose()

	if err := engine.SetExecutionLimits(limitedGas.Uint64(), DefaultLimitsOfTotalMemorySize); err != nil {
		return util.NewUint128(), "", err
	}

	// Deploy and Init.
	result, exeErr := engine.DeployAndInit(payload.Source, payload.SourceType, payload.Args)
	gasCount := engine.ExecutionInstructions()
	instructions, err := util.NewUint128FromInt(int64(gasCount))
	if err != nil || exeErr == ErrUnexpected {
		logging.VLog().WithFields(logrus.Fields{
			"err":      err,
			"exeErr":   exeErr,
			"gasCount": gasCount,
		}).Error("Unexpected error when executing deploy ")
		return util.NewUint128(), "", ErrUnexpected
	}
	if exeErr != nil && exeErr == ErrExecutionFailed && len(result) > 0 {
		exeErr = fmt.Errorf("Deploy: %s", result)
	}
	return instructions, result, exeErr
}

package core

import (
	"fmt"

	"github.com/pepperdb/pepperdb-core/common/util"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

// BinaryPayload carry some data
type BinaryPayload struct {
	Data []byte
}

// LoadBinaryPayload from bytes
func LoadBinaryPayload(bytes []byte) (*BinaryPayload, error) {
	return NewBinaryPayload(bytes), nil
}

// NewBinaryPayload with data
func NewBinaryPayload(data []byte) *BinaryPayload {
	return &BinaryPayload{
		Data: data,
	}
}

// ToBytes serialize payload
func (payload *BinaryPayload) ToBytes() ([]byte, error) {
	return payload.Data, nil
}

// BaseGasCount returns base gas count
func (payload *BinaryPayload) BaseGasCount() *util.Uint128 {
	return util.NewUint128()
}

// Execute the payload in tx
func (payload *BinaryPayload) Execute(limitedGas *util.Uint128, tx *Transaction, block *Block, ws WorldState) (*util.Uint128, string, error) {
	if block == nil || tx == nil || tx.to == nil {
		return util.NewUint128(), "", ErrNilArgument
	}

	// transfer to contract
	if tx.to.Type() == ContractAddress && block.height >= AcceptFuncAvailableHeight {
		// payloadGasLimit <= 0, v8 engine not limit the execution instructions
		if limitedGas.Cmp(util.NewUint128()) <= 0 {
			return util.NewUint128(), "", ErrOutOfGasLimit
		}

		// contract address is tx.to.
		contract, err := CheckContract(tx.to, ws)
		if err != nil {
			return util.NewUint128(), "", err
		}

		birthTx, err := GetTransaction(contract.BirthPlace(), ws)
		if err != nil {
			return util.NewUint128(), "", err
		}
		deploy, err := LoadDeployPayload(birthTx.data.Payload) // ToConfirm: move deploy payload in ctx.
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

		result, exeErr := engine.Call(deploy.Source, deploy.SourceType, ContractAcceptFunc, "")
		gasCount := engine.ExecutionInstructions()
		instructions, err := util.NewUint128FromInt(int64(gasCount))
		if err != nil || exeErr == ErrUnexpected {
			logging.VLog().WithFields(logrus.Fields{
				"err":      err,
				"exeErr":   exeErr,
				"gasCount": gasCount,
			}).Error("Unexpected error when executing binary")
			return util.NewUint128(), "", ErrUnexpected
		}

		if exeErr == ErrExecutionFailed && len(result) > 0 {
			exeErr = fmt.Errorf("Binary: %s", result)
		}
		if exeErr != nil {
			return instructions, "", exeErr
		}
	}

	return util.NewUint128(), "", nil
}

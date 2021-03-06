package core

import (
	"encoding/json"
	"fmt"

	"github.com/pepperdb/pepperdb-core/common/util"
	"github.com/pepperdb/pepperdb-core/common/util/byteutils"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

// CallPayload carry function call information
type CallPayload struct {
	Function string
	Args     string
}

// LoadCallPayload from bytes
func LoadCallPayload(bytes []byte) (*CallPayload, error) {
	payload := &CallPayload{}
	if err := json.Unmarshal(bytes, payload); err != nil {
		return nil, ErrInvalidArgument
	}
	return NewCallPayload(payload.Function, payload.Args)
}

// NewCallPayload with function & args
func NewCallPayload(function, args string) (*CallPayload, error) {

	if PublicFuncNameChecker.MatchString(function) == false {
		return nil, ErrInvalidCallFunction
	}

	if err := CheckContractArgs(args); err != nil {
		return nil, ErrInvalidArgument
	}

	return &CallPayload{
		Function: function,
		Args:     args,
	}, nil
}

// ToBytes serialize payload
func (payload *CallPayload) ToBytes() ([]byte, error) {
	return json.Marshal(payload)
}

// BaseGasCount returns base gas count
func (payload *CallPayload) BaseGasCount() *util.Uint128 {
	base, _ := util.NewUint128FromInt(60)
	return base
}

var (
	TestCompatArr = []string{"5b6a9ed8a48cfb0e6415f0df9f79cbbdac565dd139779c7972069b37c99a3913",
		"918d116f5d42b253e84497d65d2a6508fb5c4c1dbc5c1c2a1718ab718a50a509"}
	MainCompatArr = []string{"ee90d2cc5f930fe627363e9e05f1e98ea20025898201c849125659d6c0079242",
		"3db72f0d02daa26407d13ca9efc820ec618407d10d55ac15433784aaef93c659"}
)

// IsCompatibleStack return if compatible stack
func IsCompatibleStack(chainID uint32, hash byteutils.Hash) bool {
	if chainID == MainNetID {
		for i := 0; i < len(MainCompatArr); i++ {
			compatStr := MainCompatArr[i]
			if compatStr == hash.String() {
				return true
			}
		}

	} else if chainID == TestNetID {
		for i := 0; i < len(TestCompatArr); i++ {
			compatStr := TestCompatArr[i]
			if compatStr == hash.String() {
				return true
			}
		}
	}
	return false
}

// Execute the call payload in tx, call a function
func (payload *CallPayload) Execute(limitedGas *util.Uint128, tx *Transaction, block *Block, ws WorldState) (*util.Uint128, string, error) {
	if block == nil || tx == nil {
		return util.NewUint128(), "", ErrNilArgument
	}

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
	/* // useless owner.
	owner, err := ws.GetOrCreateUserAccount(birthTx.from.Bytes())
	if err != nil {
		return util.NewUint128(), "", err
	} */
	deploy, err := LoadDeployPayload(birthTx.data.Payload) // ToConfirm: move deploy payload in ctx.
	if err != nil {
		return util.NewUint128(), "", err
	}

	engine, err := block.nvm.CreateEngine(block, tx, contract, ws)
	if err != nil {
		return util.NewUint128(), "", err
	}
	defer engine.Dispose()

	if IsCompatibleStack(block.header.chainID, tx.hash) == true {
		if err := engine.SetExecutionLimits(2000, DefaultLimitsOfTotalMemorySize); err != nil {
			return util.NewUint128(), "", err
		}
	} else {
		if err := engine.SetExecutionLimits(limitedGas.Uint64(), DefaultLimitsOfTotalMemorySize); err != nil {
			return util.NewUint128(), "", err
		}
	}

	result, exeErr := engine.Call(deploy.Source, deploy.SourceType, payload.Function, payload.Args)
	gasCount := engine.ExecutionInstructions()
	instructions, err := util.NewUint128FromInt(int64(gasCount))

	if err != nil || exeErr == ErrUnexpected {
		logging.VLog().WithFields(logrus.Fields{
			"err":      err,
			"exeErr":   exeErr,
			"gasCount": gasCount,
		}).Error("Unexpected error when executing call")
		return util.NewUint128(), "", ErrUnexpected
	}
	if IsCompatibleStack(block.header.chainID, tx.hash) {
		instructions = limitedGas
	}
	if exeErr == ErrExecutionFailed && len(result) > 0 {
		exeErr = fmt.Errorf("Call: %s", result)
	}
	return instructions, result, exeErr
}

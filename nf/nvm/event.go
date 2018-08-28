package nvm

import "C"
import (
	"unsafe"

	"github.com/pepperdb/pepperdb-core/core/state"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

const (
	// EventBaseGasCount the gas count of a new event
	EventBaseGasCount = 20
)

const (
	// InnerTransferFailed failed status for transaction execute result.
	InnerTransferFailed = 0

	// InnerTransferExecutionSuccess success status for transaction execute result.
	InnerTransferExecutionSuccess = 1
)

// TransferFromContractEvent event for transfer in contract
type TransferFromContractEvent struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

// TransferFromContractFailureEvent event for transfer in contract
type TransferFromContractFailureEvent struct {
	Amount string `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
	Status uint8  `json:"status"`
	Error  string `json:"error"`
}

// EventTriggerFunc export EventTriggerFunc
//export EventTriggerFunc
func EventTriggerFunc(handler unsafe.Pointer, topic, data *C.char, gasCnt *C.size_t) {
	gTopic := C.GoString(topic)
	gData := C.GoString(data)

	e := getEngineByEngineHandler(handler)
	if e == nil {
		logging.VLog().WithFields(logrus.Fields{
			"category": 0, // ChainEventCategory.
			"topic":    gTopic,
			"data":     gData,
		}).Error("Event.Trigger delegate handler does not found.")
		return
	}

	// calculate Gas.
	*gasCnt = C.size_t(EventBaseGasCount + len(gTopic) + len(gData))

	contractTopic := EventNameSpaceContract + "." + gTopic
	event := &state.Event{Topic: contractTopic, Data: gData}
	e.ctx.state.RecordEvent(e.ctx.tx.Hash(), event)
}

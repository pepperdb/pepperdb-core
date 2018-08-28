package nvm

import "C"

import (
	"github.com/pepperdb/pepperdb-core/common/util/logging"
)

// V8Log export V8Log
//export V8Log
func V8Log(level int, msg *C.char) {
	s := C.GoString(msg)

	switch level {
	case 1:
		logging.CLog().Debug(s)
	case 2:
		logging.CLog().Warn(s)
	case 3:
		logging.CLog().Info(s)
	case 4:
		logging.CLog().Error(s)
	default:
		logging.CLog().Error(s)
	}
}

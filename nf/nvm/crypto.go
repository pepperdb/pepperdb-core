package nvm

import "C"

import (
	"crypto/md5"

	"github.com/pepperdb/pepperdb-core/core"
	"github.com/pepperdb/pepperdb-core/common/crypto/hash"
	"github.com/pepperdb/pepperdb-core/common/crypto/keystore"
	"github.com/pepperdb/pepperdb-core/common/util/byteutils"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

// Sha256Func ..
//export Sha256Func
func Sha256Func(data *C.char, gasCnt *C.size_t) *C.char {
	s := C.GoString(data)
	*gasCnt = C.size_t(len(s) + CryptoSha256GasBase)

	r := hash.Sha256([]byte(s))
	return C.CString(byteutils.Hex(r))
}

// Sha3256Func ..
//export Sha3256Func
func Sha3256Func(data *C.char, gasCnt *C.size_t) *C.char {
	s := C.GoString(data)
	*gasCnt = C.size_t(len(s) + CryptoSha3256GasBase)

	r := hash.Sha3256([]byte(s))
	return C.CString(byteutils.Hex(r))
}

// Ripemd160Func ..
//export Ripemd160Func
func Ripemd160Func(data *C.char, gasCnt *C.size_t) *C.char {
	s := C.GoString(data)
	*gasCnt = C.size_t(len(s) + CryptoRipemd160GasBase)

	r := hash.Ripemd160([]byte(s))
	return C.CString(byteutils.Hex(r))
}

// RecoverAddressFunc ..
//export RecoverAddressFunc
func RecoverAddressFunc(alg int, data, sign *C.char, gasCnt *C.size_t) *C.char {
	d := C.GoString(data)
	s := C.GoString(sign)

	*gasCnt = C.size_t(CryptoRecoverAddressGasBase)

	plain, err := byteutils.FromHex(d)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"hash": d,
			"sign": s,
			"alg":  alg,
			"err":  err,
		}).Debug("convert hash to byte array error.")
		return nil
	}
	cipher, err := byteutils.FromHex(s)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"data": d,
			"sign": s,
			"alg":  alg,
			"err":  err,
		}).Debug("convert sign to byte array error.")
		return nil
	}
	addr, err := core.RecoverSignerFromSignature(keystore.Algorithm(alg), plain, cipher)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"data": d,
			"sign": s,
			"alg":  alg,
			"err":  err,
		}).Debug("recover address error.")
		return nil
	}

	return C.CString(addr.String())
}

// Md5Func ..
//export Md5Func
func Md5Func(data *C.char, gasCnt *C.size_t) *C.char {
	s := C.GoString(data)
	*gasCnt = C.size_t(len(s) + CryptoMd5GasBase)

	r := md5.Sum([]byte(s))
	return C.CString(byteutils.Hex(r[:]))
}

// Base64Func ..
//export Base64Func
func Base64Func(data *C.char, gasCnt *C.size_t) *C.char {
	s := C.GoString(data)
	*gasCnt = C.size_t(len(s) + CryptoBase64GasBase)

	r := hash.Base64Encode([]byte(s))
	return C.CString(string(r))
}

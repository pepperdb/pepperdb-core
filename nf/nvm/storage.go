package nvm

import "C"

import (
	"errors"
	"regexp"
	"unsafe"

	"github.com/pepperdb/pepperdb-core/common/trie"
	"github.com/pepperdb/pepperdb-core/common/util/logging"
	"github.com/sirupsen/logrus"
)

var (
	// StorageKeyPattern the pattern of varible key stored in stateDB
	/*
		const fieldNameRe = /^[a-zA-Z_$][a-zA-Z0-9_]+$/;
		var combineStorageMapKey = function (fieldName, key) {
			return "@" + fieldName + "[" + key + "]";
		};
	*/
	StorageKeyPattern = regexp.MustCompile("^@([a-zA-Z_$][a-zA-Z0-9_]+?)\\[(.*?)\\]$")
	// DefaultDomainKey the default domain key
	DefaultDomainKey = "_"
	// ErrInvalidStorageKey invalid storage key error
	ErrInvalidStorageKey = errors.New("invalid storage key")
)

// hashStorageKey return the key hash.
// There are two kinds of key, the one is ItemKey, the other is Map-ItemKey.
// ItemKey in SmartContract is used for object storage.
// For example, the ItemKey for the statement "token.totalSupply = 1000" is "totalSupply".
// Map-ItemKey in SmartContrat is used for Map storage.
// For example, the Map-ItemKey for the statement "token.balances.set('addr1', 100)" is "@balances[addr1]".
func parseStorageKey(key string) (string, string, error) {
	matches := StorageKeyPattern.FindAllStringSubmatch(key, -1)
	if matches == nil {
		return DefaultDomainKey, key, nil
	}

	return matches[0][1], matches[0][2], nil
}

// StorageGetFunc export StorageGetFunc
//export StorageGetFunc
func StorageGetFunc(handler unsafe.Pointer, key *C.char, gasCnt *C.size_t) *C.char {
	_, storage := getEngineByStorageHandler(uint64(uintptr(handler)))
	if storage == nil {
		logging.VLog().Error("Failed to get storage handler.")
		return nil
	}

	k := C.GoString(key)

	// calculate Gas.
	*gasCnt = C.size_t(0)

	domainKey, itemKey, err := parseStorageKey(k)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"handler": uint64(uintptr(handler)),
			"key":     k,
			"err":     err,
		}).Debug("Invalid storage key.")
		return nil
	}

	val, err := storage.Get(trie.HashDomains(domainKey, itemKey))
	if err != nil {
		if err != ErrKeyNotFound {
			logging.VLog().WithFields(logrus.Fields{
				"handler": uint64(uintptr(handler)),
				"key":     k,
				"err":     err,
			}).Debug("StorageGetFunc get key failed.")
		}
		return nil
	}

	return C.CString(string(val))
}

// StoragePutFunc export StoragePutFunc
//export StoragePutFunc
func StoragePutFunc(handler unsafe.Pointer, key *C.char, value *C.char, gasCnt *C.size_t) int {
	_, storage := getEngineByStorageHandler(uint64(uintptr(handler)))
	if storage == nil {
		logging.VLog().Error("Failed to get storage handler.")
		return 1
	}

	k := C.GoString(key)
	v := []byte(C.GoString(value))

	// calculate Gas.
	*gasCnt = C.size_t(len(k) + len(v))

	domainKey, itemKey, err := parseStorageKey(k)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"handler": uint64(uintptr(handler)),
			"key":     k,
			"err":     err,
		}).Debug("Invalid storage key.")
		return 1
	}

	err = storage.Put(trie.HashDomains(domainKey, itemKey), v)
	if err != nil && err != ErrKeyNotFound {
		logging.VLog().WithFields(logrus.Fields{
			"handler": uint64(uintptr(handler)),
			"key":     k,
			"err":     err,
		}).Debug("StoragePutFunc put key failed.")
		return 1
	}

	return 0
}

// StorageDelFunc export StorageDelFunc
//export StorageDelFunc
func StorageDelFunc(handler unsafe.Pointer, key *C.char, gasCnt *C.size_t) int {
	_, storage := getEngineByStorageHandler(uint64(uintptr(handler)))
	if storage == nil {
		logging.VLog().Error("Failed to get storage handler.")
		return 1
	}

	k := C.GoString(key)

	// calculate Gas.
	*gasCnt = C.size_t(0)

	domainKey, itemKey, err := parseStorageKey(k)
	if err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"handler": uint64(uintptr(handler)),
			"key":     k,
			"err":     err,
		}).Debug("invalid storage key.")
		return 1
	}

	err = storage.Del(trie.HashDomains(domainKey, itemKey))
	if err != nil && err != ErrKeyNotFound {
		logging.VLog().WithFields(logrus.Fields{
			"handler": uint64(uintptr(handler)),
			"key":     k,
			"err":     err,
		}).Debug("StorageDelFunc del key failed.")
		return 1
	}

	return 0
}

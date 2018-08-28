package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	// ErrFileExists file exists
	ErrFileExists = errors.New("file exists")
)

// FileExists check file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// FileWrite write file to path
func FileWrite(file string, content []byte, overwrite bool) error {
	// Create the keystore directory with appropriate permissions
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return err
	}
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()

	if overwrite {
		if exist, _ := FileExists(file); exist {
			if err := os.Remove(file); err != nil {
				os.Remove(f.Name())
				return err
			}
		}
	}

	return os.Rename(f.Name(), file)
}

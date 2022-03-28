// Package dir provides additional directory related functions
// on top of os file functions functions.
package dir

import (
	"errors"
	"io"
	"os"
)

// IsEmpty checks if a directory is empty or not.
// It returns true if a directory is empty.
func IsEmpty(path string) (bool, error) {

	fInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if !fInfo.IsDir() {
		return false, errors.New("not a directory")
	}

	f, err := os.Open(path)
	if err != nil {
		return false, err
	}

	dirs, err := f.Readdirnames(0)
	if err != nil && err != io.EOF {
		if err == io.EOF {
			return true, nil
		}
		return false, err
	}

	if len(dirs) == 0 {
		return true, nil
	}
	return false, nil
}

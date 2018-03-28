// Package file provides convenience methods for file access for the parent Package bufioc
package file

import (
	"os"

	"github.com/curlymon/bufioc"
)

// NewReader opens the named file for reading. If successful, the associated file
// descriptor has mode O_RDONLY.
// It returns returns a new *bufioc.Reader whose buffer has the default size.
// If there is an error, it will be of type *os.PathError.
func NewReader(path string) (*bufioc.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufioc.NewReader(f), nil
}

// NewReaderSize opens the named file for reading. If successful, the associated file
// descriptor has mode O_RDONLY.
// It returns a new *bufioc.Reader whose buffer has at least the specified size.
// If there is an error, it will be of type *os.PathError.
func NewReaderSize(path string, size int) (*bufioc.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufioc.NewReaderSize(f, size), nil
}

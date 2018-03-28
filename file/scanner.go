package file

import (
	"os"

	"github.com/curlymon/bufioc"
)

// NewScanner opens the named file for reading. If successful, the associated file
// descriptor has mode O_RDONLY.
// It returns returns a new *bufioc.Scanner.
// If there is an error, it will be of type *os.PathError.
func NewScanner(path string) (*bufioc.Scanner, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return bufioc.NewScanner(f), nil
}

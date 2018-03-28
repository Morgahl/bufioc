package file

import (
	"os"

	"github.com/curlymon/bufioc"
)

// NewWriter opens the named file with specified flag (O_RDONLY etc.) and
// perm (before umask), if applicable.
// It returns a new *bufioc.Writer whose buffer has the default size.
// If there is an error, it will be of type *PathError.
func NewWriter(path string, flag int, perm os.FileMode) (*bufioc.Writer, error) {
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}

	return bufioc.NewWriter(f), nil
}

// NewWriterSize opens the named file with specified flag (O_RDONLY etc.) and
// perm (before umask), if applicable.
// It returns a new *bufioc.Writer whose buffer has at least the specified size.
// If there is an error, it will be of type *PathError.
func NewWriterSize(path string, flag int, perm os.FileMode, size int) (*bufioc.Writer, error) {
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}

	return bufioc.NewWriterSize(f, size), nil
}

// NewWriterCreate creates the named file with mode 0666 (before umask),
// truncating it if it already exists. If successful, the associated file
// descriptor has mode O_RDWR.
// It returns a new *bufioc.Writer whose buffer has the default size.
// If there is an error, it will be of type *os.PathError.
func NewWriterCreate(path string) (*bufioc.Writer, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return bufioc.NewWriter(f), nil
}

// NewWriterCreateSize creates the named file with mode 0666 (before umask),
// truncating it if it already exists. If successful, the associated file
// descriptor has mode O_RDWR.
// It returns a new *bufioc.Writer whose buffer has at least the specified size.
// If there is an error, it will be of type *os.PathError.
func NewWriterCreateSize(path string, size int) (*bufioc.Writer, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return bufioc.NewWriterSize(f, size), nil
}

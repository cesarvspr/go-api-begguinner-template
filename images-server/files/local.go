package files

import "path/filepath"

// Local is an implementation fo the storage interface which works with the
// local disk on the currente machine
type Local struct {
	maxFileSize int // maximum numbber of bytes for files
	basePath    string
}

// NewLocal creates a new local filesystem with the given basepath
// basePath is the base directory to save files to
// masSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil
}

package handlers

import (
	"github.com/cesarvspr/images-server/files"
	"github.com/hashicorp/go-hclog"
)

//Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

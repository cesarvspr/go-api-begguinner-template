package handlers

import "github.com/hashicorp/go-hclog"

//Files is a handler for reading and writing files
type Files struct {
	log hclog.Logger
	store files.Storage 
}


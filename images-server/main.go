package main

import (
	"github.com/cesarvspr/golang-modules/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

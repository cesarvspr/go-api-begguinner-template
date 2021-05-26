package main

import (
	"net/http"
	"os"

	"github.com/cesarvspr/images-server/env"
	"github.com/gorilla/mux"

	"github.com/cesarvspr/images-server/files"

	hclog "github.com/hashicorp/go-hclog"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
var logLevel = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
var basePath = env.String("BASE_PATH", false, "./imagestore", "Base path to save images")

func main() {
	env.Parse()

	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "product-images",
			Level: hclog.LevelFromString(*logLevel),
		},
	)
	//create logger for the server from thge default logger
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true})

	//create the storage classs, use local storage
	// max filesize 5MB
	stor, err := files.NewLocal(*basePath, 1024*1000*5)
	if err != nil {
		l.Error("unable to create storagete", "error", err)
		os.Exit(1)
	}

	// create the handlers
	fh := handlers.NewFiles(stor, l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// filename regex: {filename:[a-zA-Z]+\\.[a-z]{3}}
	// problem with FileServer is that it is dumb
	ph := sm.Methods(http.MethodPost).Subrouter()
	ph.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", fh.ServeHTTP)

}

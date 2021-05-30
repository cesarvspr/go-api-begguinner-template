package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/cesarvspr/images-server/files_sdk"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

//Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files_sdk.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files_sdk.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

//UploadRestimplements the http.Handler interface
func (f *Files) UploadRest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]
	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalis id or filename as the mux router will not send requests
	// here unless they have the correct parameters

	f.saveFile(id, fn, rw, r)
}

func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	f.log.Info("Process form for id", "id", id)

}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("save file for a product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("Unable to save", "error", err)
		http.Error(rw, "unable to save file", http.StatusInternalServerError)
	}
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Error("invalid", "path", uri)
	http.Error(rw, "invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

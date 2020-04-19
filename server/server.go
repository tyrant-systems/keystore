package server

import (
	"errors"
	"log"
	"net/http"
	"path"
)

// Exported errors.
var (
	ErrNotAKeyFile = errors.New("not a keyfile")
)

// restrictedFileSystem allows reading only files named "key".
type restrictedFileSystem struct {
	fs http.FileSystem
}

// Open implements the http.FileSystem interface. It serves only files named "key".
func (rfs restrictedFileSystem) Open(resource string) (http.File, error) {
	log.Printf("GET %s", resource)

	_, f := path.Split(resource)
	if f != "keys" {
		return nil, ErrNotAKeyFile
	}

	fp, err := rfs.fs.Open(resource)
	if err != nil {
		return nil, err
	}

	s, err := fp.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		return nil, ErrNotAKeyFile
	}

	return fp, nil
}

// KeyServer wraps an http.FileSystem. It serves only files named "key".
type KeyServer struct {
	mux *http.ServeMux
	fs  http.FileSystem
}

// New instantiates a new key server and will serve keys from "rootDir".
func New(rootDir string) KeyServer {
	return KeyServer{
		mux: http.NewServeMux(),
		fs:  restrictedFileSystem{http.Dir(rootDir)},
	}
}

// ListenAndServeKeyFiles starts the key server.
func (ks KeyServer) ListenAndServeKeyFiles(addr string) error {
	fs := http.FileServer(ks.fs)
	ks.mux.Handle("/user/", http.StripPrefix("/user", fs))
	return http.ListenAndServe(addr, ks.mux)
}

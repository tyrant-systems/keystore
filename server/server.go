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

// withLogging is a simple http middleware handler for logging all inbound requests.
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

// restrictedFileSystem allows reading only files named "key".
type restrictedFileSystem struct {
	fs http.FileSystem
}

// Open implements the http.FileSystem interface. It serves only files named "key".
func (rfs restrictedFileSystem) Open(resource string) (http.File, error) {
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
	ks.mux.Handle(
		"/",
		withLogging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "500 Internal Server Error", 500)
		})),
	)

	ks.mux.Handle(
		"/user/",
		withLogging(http.StripPrefix("/user", http.FileServer(ks.fs))),
	)

	return http.ListenAndServe(addr, ks.mux)
}

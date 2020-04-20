package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/tyrant-systems/keystore/server"
)

var (
	flagConfigFilepath string
)

func init() {
	flag.StringVar(
		&flagConfigFilepath,
		"config",
		"config.json",
		"filepath to a configuration json file",
	)

	flag.Parse()
}

func main() {
	raw, err := ioutil.ReadFile(flagConfigFilepath)
	if err != nil {
		log.Fatal(err)
	}

	fh := struct{ Configuration map[string]string }{}

	if err := json.Unmarshal(raw, &fh); err != nil {
		log.Fatal(err)
	}

	var (
		dir  = fh.Configuration["user_dir_abs_path"]
		addr = fh.Configuration["listen_addr"]
	)

	ks := server.New(dir)

	log.Printf("serving keys from %s on %s", dir, addr)

	if err := ks.ListenAndServeKeyFiles(addr); err != nil {
		log.Fatal(err)
	}
}

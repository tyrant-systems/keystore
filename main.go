package main

import (
	"log"

	"github.com/tyrant-systems/keystore/server"
)

func main() {
	var (
		dir  = "./user"
		addr = ":3030"
	)

	ks := server.New(dir)

	log.Printf("serving keys from %s on %s", dir, addr)

	if err := ks.ListenAndServeKeyFiles(addr); err != nil {
		log.Fatal(err)
	}
}

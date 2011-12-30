package main

import (
	"log"
	"github.com/davecheney/nfs"
)

func main() {
	pm, err := nfs.NewPortmapper("tcp", "stora.local")
	if err != nil {
		log.Fatalf("unable to contact portmapper: %v", err)
	}
	pm.Close()	
}

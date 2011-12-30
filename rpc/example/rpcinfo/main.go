package main

import (
	"log"
	_ "fmt"
	"github.com/davecheney/nfs/rpc"
)

func main() {
	pm, err := rpc.DialPortmapper("tcp", "stora.local")
	if err != nil {
		log.Fatalf("unable to connect to portmapper: %v", err)
	}
	// pm.Dump()
	pm.Close()
}

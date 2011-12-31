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
	if _, err := pm.Dump() ; err != nil {
		log.Fatalf("failed to call PORTMAP.DUMP: %v", err)
	}
	pm.Close()
}

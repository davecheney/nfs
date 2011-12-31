package main

import (
	"fmt"
	"github.com/davecheney/nfs/rpc"
	"log"
)

func main() {
	pm, err := rpc.DialPortmapper("tcp", "stora.local")
	if err != nil {
		log.Fatalf("unable to connect to portmapper: %v", err)
	}
	mappings, err := pm.Dump(); 
	if err != nil {
		log.Fatalf("failed to call PORTMAP.DUMP: %v", err)
	}
	fmt.Println("program\tvers\tproto\tport")
	for _, m := range mappings {
		fmt.Printf("%d\t%d\t%d\t%d", m.Prog, m.Vers, m.Prot, m.Port)
	}
	pm.Close()
}

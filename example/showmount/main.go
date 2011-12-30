package main

import (
	"log"
	"github.com/davecheney/nfs"
	"github.com/davecheney/nfs/rpc"
)

func main() {
	pm, err := rpc.DialPortmapper("tcp", "stora.local")
	if err != nil {
		log.Fatalf("unable to contact portmapper: %v", err)
	}
	// get MOUNT port

	m := rpc.Mapping {
		Prog: nfs.MOUNT_PROG,
		Vers: nfs.MOUNT_VERS,
	 	Prot: rpc.IPPROTO_TCP,
		Port: 0,		 	
	}
	port, err := pm.Getport(m)
	if err != nil {
		log.Fatalf("unable to get MOUNT port: %v", err)
	}
	log.Printf("MOUNT port: %d", port)
	pm.Close()	
}

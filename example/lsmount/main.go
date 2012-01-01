package main

import (
	"fmt"
	"github.com/davecheney/nfs"
	"github.com/davecheney/nfs/rpc"
	"log"
)

func main() {
	pm, err := rpc.DialPortmapper("tcp", "stora.local")
	if err != nil {
		log.Fatalf("unable to contact portmapper: %v", err)
	}
	// get MOUNT port
	m := rpc.Mapping{
		Prog: nfs.MOUNT_PROG,
		Vers: nfs.MOUNT_VERS,
		Prot: rpc.IPPROTO_TCP,
		Port: 0,
	}
	port, err := pm.Getport(m)
	if err != nil {
		log.Fatalf("unable to get MOUNT port: %v", err)
	}
	log.Println("MOUNT", port)
	defer pm.Close()
	mount, err := nfs.DialMount("tcp", fmt.Sprintf("stora.local:%d", port))
	if err != nil {
		log.Fatal("unable to dial MOUNT service: %v", err)
	}
	defer mount.Close()
	auth := &rpc.AUTH_UNIX {
		Stamp: 1,
		Machinename: "localhost",
		Uid: 0,
		Gid: 0,
	}
	v, err := mount.Mount("/export",auth.Auth())
	if err != nil {
		log.Fatalf("unable to mount volume: %v", err)
	}
	v.Unmount()
}

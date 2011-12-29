package main

import (
	"log"
	_ "fmt"
	"github.com/davecheney/nfs/rpc"
)

type DUMP struct {
	rpc.Header
}

func main() {
	c, err := rpc.DialTCP("tcp", "stora.local:111")
	if err != nil {
		log.Fatalf("unable to connect to portmapper: %v", err)
	}
	
	dump := &DUMP{
		rpc.Header {
			Rpcvers: 2,
			Prog: 100000,
			Vers: 2,
			Proc: 4,			
			Cred: rpc.AUTH_NULL,
			Verf: rpc.AUTH_NULL,
		},
	}

	if err := c.Call(dump, nil); err != nil {
		log.Fatalf("%v\n", err)
	}	
	c.Close()
}

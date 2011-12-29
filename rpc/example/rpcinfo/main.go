package main

import (
	"log"
	_ "fmt"
	"github.com/davecheney/nfs/rpc"
)

const (
	PROGRAM = 100000 // portmap
	VERSION = 2
	PROC = 0 // NULL
)

func main() {
	c, err := rpc.DialTCP("tcp", "stora.local:111")
	if err != nil {
		log.Fatalf("unable to connect to portmapper: %v", err)
	}
	_, err = c.Call(PROGRAM, VERSION, PROC).Send()
	if err != nil {
		log.Fatalf("%v\n", err)
	}	
	c.Close()
}

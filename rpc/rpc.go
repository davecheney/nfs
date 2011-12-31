package rpc

import (
	"io"
)

type transport interface {
	send([]byte) error
	recv() ([]byte, error)
	io.Closer
}

type mismatch_info struct {
	low  uint32
	high uint32
}

type Header struct {
        Rpcvers uint32
        Prog    uint32
        Vers    uint32
        Proc    uint32
        Cred    Auth
        Verf    Auth
}

type message struct {
        Xid     uint32
        Msgtype uint32
        Body    interface{}
}

type Auth struct {
	Flavor uint32
	Body []byte
}

var AUTH_NULL = Auth {
	0, 
	[]byte { },
}



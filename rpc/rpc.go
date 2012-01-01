package rpc

import (
	"bytes"
	"github.com/davecheney/nfs/xdr"
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
	Body   []byte
}

var AUTH_NULL = Auth{
	0,
	[]byte{},
}

type AUTH_UNIX struct {
	Stamp       uint32
	Machinename string
	Uid         uint32
	Gid         uint32
	Gids        uint32
}

// Auth converts a into an Auth opaque struct
func (a AUTH_UNIX) Auth() Auth {
	w := new(bytes.Buffer)
	xdr.Write(w, a)
	return Auth{
		1,
		w.Bytes(),
	}
}

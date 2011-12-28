package nfs

import (
	"github.com/davecheney/nfs/rpc"
)

func DialPortmapper(addr string) (*Portmapper, error) {
	return newPortmapper(rpc.DialTCP("tcp", addr+":111"))
}

func newPortmapper(conn rpc.Conn, err error) (*Portmapper, error) {
	if err != nil {
		return nil, err
	}
	return &Portmapper { conn }, nil
}

type Portmapper struct {
	conn rpc.Conn
}

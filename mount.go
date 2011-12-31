package nfs

// MOUNT
// RFC 1813 Section 5.0

import (
	"github.com/davecheney/nfs/rpc"
)

const (
	MOUNT_PROG = 100005
	MOUNT_VERS = 3

	MOUNTPROC3_EXPORT = 5
)

type Export struct {
	Dir    string
	Groups []Group
}

type Group struct {
	Name string
}

type Mount struct {
	*rpc.Client
}

func (m *Mount) Exports() ([]Export, error) {
	type export struct {
		rpc.Header
	}
	msg := &export{
		rpc.Header{
			Rpcvers: 2,
			Prog:    MOUNT_PROG,
			Vers:    MOUNT_VERS,
			Proc:    MOUNTPROC3_EXPORT,
			Cred:    rpc.AUTH_NULL,
			Verf:    rpc.AUTH_NULL,
		},
	}
	_, err := m.Call(msg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func DialMount(net, addr string) (*Mount, error) {
	client, err := rpc.DialTCP(net, addr)
	if err != nil {
		return nil, err
	}
	return &Mount{client}, nil
}

package nfs

import (
	"github.com/davecheney/nfs/rpc"
)

// PORTMAP 

type Portmapper struct {
	client *rpc.Client
}

func (p *Portmapper) Getport(

func (p *Portmapper) Close() error {
	return p.client.Close()
}

func NewPortmapper(net, host string) (*Portmapper, error) {
	client, err := rpc.DialTCP(net, host+":111")
	if err != nil {
		return nil, err
	}
	return &Portmapper{client}, nil
}

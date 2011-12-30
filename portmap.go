package main

import (
	"github.com/davecheney/nfs/rpc"
)

// PORTMAP 

type Portmapper struct {
	client rpc.Client
}

func (p *Portmapper) Close() error {
	return client.Close()
}

func NewPortmapper(net, host string) (*Portmapper, error) {
	client, err := rpc.DialTCP(network, host+":111")
	if err != nil {
		return nil, err
	}
	return &Portmapper{client}
}

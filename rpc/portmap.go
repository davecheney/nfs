package rpc

import (
	"fmt"
)

// PORTMAP
// RFC 1057 Section A.1

const (
	PMAP_PORT = 111
	PMAP_PROG = 10000
	PMAP_VERS = 2

	PMAPPROC_GETPORT = 3
	PMAPPROC_DUMP = 4

	IPPROTO_TCP = 6
	IPPROTO_UDP = 17
)

type Mapping struct {
	Prog uint32
	Vers uint32
	Prot uint32
	Port uint32
}

type Portmapper struct {
	client *Client
}

func (p *Portmapper) Getport(mapping Mapping) (uint16, error) {
	type getport struct {
		Header
		Mapping
	}
	
	msg := &getport {
		Header {
                        Rpcvers: 2,
                        Prog: PMAP_PROG,
                        Vers: PMAP_VERS,
                        Proc: PMAPPROC_GETPORT,                        
                        Cred: AUTH_NULL,
                        Verf: AUTH_NULL,
                },
		mapping,
	}
	err := p.client.Call(msg, nil)
	return 0, err
}

func (p *Portmapper) Close() error {
	return p.client.Close()
}

func DialPortmapper(net, host string) (*Portmapper, error) {
	client, err := DialTCP(net, fmt.Sprintf("%s:%d", host, PMAP_PORT))
	if err != nil {
		return nil, err
	}
	return &Portmapper{client}, nil
}

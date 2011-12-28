package rpc

import (
	"testing"

	"github.com/davecheney/nfs/xdr"
)

func TestMarshalPortmap(t *testing.T) {
	msg := Call {
		msg{
			xid: 1,
			mtype: CALL,
		},
		call{
			rpcvers: 2,
			prog: 100000,
			vers: 2,
			proc: 0,
		},
	}
	_, err := xdr.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}
}

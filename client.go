package nfs

import (
	"github.com/davecheney/nfs/rpc"
)

type Client struct {
	*rpc.Client
}

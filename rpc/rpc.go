package rpc

import (
	"io"
)

type transport interface {
        send([]byte) error
        recv() ([]byte, error)
        io.Closer
}


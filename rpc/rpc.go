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
	low uint32
	high uint32
}


package rpc

import (
	"io"
)

type transport interface {
	io.ReadWriteCloser
}

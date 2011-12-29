package xdr

import (
	"io"
	"encoding/binary"
)

type Reader struct {
	io.Reader
}

func (r *Reader) ReadUint32(v *uint32) error {
	return binary.Read(r, binary.BigEndian, v)
}

package xdr

import (
	"io"
	"encoding/binary"
)

type Writer struct {
	io.Writer
}

func (w *Writer) WriteUint32(v uint32) error {
	return binary.Write(w, binary.BigEndian, v)
}

func (w *Writer) WriteInt32(v int32) error {
	return binary.Write(w, binary.BigEndian, v)
}

func (w *Writer) WriteUint64(v uint64) error {
	return binary.Write(w, binary.BigEndian, v)
}

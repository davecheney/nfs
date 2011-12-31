package xdr

import (
	"encoding/binary"
)

func Uint32(b []byte) (uint32, []byte) {
	return binary.BigEndian.Uint32(b[0:4]), b[4:]
}

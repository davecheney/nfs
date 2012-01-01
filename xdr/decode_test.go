package xdr

import (
	"bytes"
	"testing"
)

func TestRead(t *testing.T) {
	type X struct {
		A, B, C, D uint32
	}
	x := new(X)
	b := []byte{ 
		0, 0, 0, 1,
		0, 0, 0, 2,
		0, 0, 0, 3,
		0, 0, 0, 4,
		1,
	}
	buf := bytes.NewBuffer(b)
	Read(buf, x)
}	

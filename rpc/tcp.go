package rpc

import (
	"io"
	"sync"
	"encoding/binary"
)

type tcpTransport struct {
	io.Reader
	rlock sync.Mutex
	io.WriteCloser
	wlock sync.Mutex
}

func (t *tcpTransport) Read(buf []byte) (int, error) {
	t.rlock.Lock()
	defer t.rlock.Unlock()
	return t.Reader.Read(buf)
}	

func (t *tcpTransport) Write(buf []byte) (int, error) {
	t.wlock.Lock()
	defer t.wlock.Unlock()
	var hdr uint32 = uint32(len(buf)) | 0x80000000
	if err := binary.Write(t.WriteCloser, binary.BigEndian, &hdr) ; err != nil {
		return 0, err
	}
	return t.WriteCloser.Write(buf)	
}

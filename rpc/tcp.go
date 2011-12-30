package rpc

import (
	"io"
	"sync"
	"encoding/binary"
)

type tcpTransport struct {
	io.Reader
	io.WriteCloser
	rlock, wlock sync.Mutex
}

func (t *tcpTransport) recv() ([]byte, error) {
	t.rlock.Lock()
	defer t.rlock.Unlock()
	var hdr uint32
	if err := binary.Read(t, binary.BigEndian, &hdr); err != nil {
		return nil, err
	}
	buf := make([]byte, hdr & 0x7fffffff)
	if _, err := io.ReadFull(t, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func (t *tcpTransport) send(buf []byte) error {
	t.wlock.Lock()
	defer t.wlock.Unlock()
	var hdr uint32 = uint32(len(buf)) | 0x80000000
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, hdr)
	_, err := t.WriteCloser.Write(append(b, buf...))	
	return err
}

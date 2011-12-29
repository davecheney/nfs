package rpc

import (
	"io"
	"sync"
	"encoding/binary"
	"log"
)

type transport interface {
	io.WriteCloser
	run() chan []byte
}

type tcpTransport struct {
	io.Reader
	io.WriteCloser
	wlock sync.Mutex
}

func (t *tcpTransport) run() chan []byte {
	c := make(chan []byte, 16)
	go func() {
		for {
			var hdr uint32
			if err := binary.Read(t, binary.BigEndian, &hdr); err != nil {
				if err != io.EOF {
					log.Print(err)
				}
				return	
			}		
			buf := make([]byte, hdr & 0x7fffffff)
			if _, err := io.ReadFull(t, buf); err != nil {
				log.Print(err)
				return
			}
			c <- buf
		}
	}()
	return c
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

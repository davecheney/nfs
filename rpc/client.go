package rpc

import (
	"bufio"
	"bytes"
	"net"

	"github.com/davecheney/nfs/xdr"
)

type Client struct {
	transport
}

func DialTCP(network, addr string) (*Client, error) {
        a, err := net.ResolveTCPAddr(network, addr) 
        if err != nil {
                return nil, err
        }
        conn, err := net.DialTCP(a.Network(), nil, a) 
        if err != nil {
                return nil, err
        }
        return &Client{ 	
		transport: &tcpTransport {
			Reader: bufio.NewReader(conn),
                        WriteCloser: conn,
                },
        }, nil
}

func (c *Client) Call(prog, vers, proc uint32) *Writer {
	buf := new(bytes.Buffer)
	w := &Writer{
		c,
		buf,
		xdr.Writer {
			buf,
		},
	}
	w.WriteUint32(0xcafebabe) // XID
	w.WriteUint32(0) // CALL
	w.WriteUint32(2) // RPC version 2
	w.WriteUint32(prog)
	w.WriteUint32(vers)
	w.WriteUint32(proc)
	w.WriteUint64(0) // AUTH_NULL
	w.WriteUint64(0) // AUTH_NULL
	return w 
}

type Reader struct {
	xdr.Reader
}

type Writer struct {
	*Client
	buf *bytes.Buffer
	xdr.Writer
}

func (w *Writer) Send() (*Reader, error) {
	_, err := w.Client.transport.Write(w.buf.Bytes())
	if err != nil {
		return nil, err
	}
	r := &Reader { xdr.Reader { w.Client } }
	var xid uint32
	if err := r.ReadUint32(&xid) ; err != nil {
		return nil, err
	}
	var reply uint32
	if err := r.ReadUint32(&reply); err != nil {
		return nil, err
	}
	var reply_stat uint32
	if err := r.ReadUint32(&reply_stat); err != nil {
		return nil, err
	}
	switch reply_stat {
	case 0:
		
	}
	return nil, nil
}

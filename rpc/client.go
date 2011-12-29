package rpc

import (
	"bufio"
	"net"
	"io"

	"github.com/davecheney/nfs/xdr"
)

type Client struct {
	transport
	read chan []byte
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
	t := &tcpTransport {
		Reader: bufio.NewReader(conn),
               WriteCloser: conn,
        }
	return &Client { t, t.run() }, nil	
}

func (c *Client) Call(call, reply interface{}) error {
	msg := &message {
		Xid: 0xcafebabe,
		Msgtype: 0,
		Body: call,
	}
	buf, err := xdr.Marshal(msg)
	if err != nil {
		return err
	}
	if _, err := c.transport.Write(buf); err != nil {
		return err
	}
	buf, ok := <- c.read
	if !ok {
		return io.EOF
	}
	return xdr.Unmarshal(reply, buf)
}


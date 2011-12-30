package rpc

import (
	"bufio"
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
	t := &tcpTransport {
		Reader: bufio.NewReader(conn),
                WriteCloser: conn,
        }
	return &Client { t }, nil	
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
	if err := c.send(buf); err != nil {
		return err
	}
	buf, err = c.recv()  
	if err != nil {
		return err
	}
	return xdr.Unmarshal(reply, buf)
}


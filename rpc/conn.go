package rpc

import (
	"net"
	"io"
	"bufio"
	"encoding/binary"
	"fmt"
	
	"github.com/davecheney/nfs/xdr"
)

type Conn interface {
	Send(msg *Call) (*Reply, error)
	io.Closer
}

func DialTCP(network, addr string) (Conn, error) {
	a, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP(a.Network(), nil, a)
	if err != nil {
		return nil, err
	}
	return &tcpconn { conn, bufio.NewReader(conn), conn }, nil
}

type tcpconn struct {
	w io.Writer
	r io.Reader
	io.Closer
}

func (t *tcpconn) Send(msg *Call) (*Reply, error) {
	b, err := xdr.Marshal(msg)
	if err != nil {
		return nil, err
	}
	_, err = t.w.Write(b)
	if err != nil {
		return nil, err
	}
	return t.readReply()	
}

func (t *tcpconn) readReply() (*Reply, error) {
	var xid uint
	if err := binary.Read(t.r, binary.BigEndian, &xid); err != nil {
		return nil, err
	}
	var mtype msg_type
	if err := binary.Read(t.r, binary.BigEndian, &mtype); err != nil {
		return nil, err
	}
	if mtype != REPLY {
		return nil, fmt.Errorf("unexpected mtype: %v", mtype)
	}
	var stat int 
	if err := binary.Read(t.r, binary.BigEndian, &stat); err != nil {
		return nil, err
	}
	return nil,nil	
}


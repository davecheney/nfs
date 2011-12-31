package rpc

type Header struct {
	Rpcvers uint32
	Prog    uint32
	Vers    uint32
	Proc    uint32
	Cred    Auth
	Verf    Auth
}

type Auth struct {
	Flavor uint32
	Body   []byte
}

var AUTH_NULL = Auth{0, []byte{}}

type message struct {
	Xid     uint32
	Msgtype uint32
	Body    interface{}
}

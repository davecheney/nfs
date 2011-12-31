package rpc

type msg struct {
	xid   uint
	mtype msg_type
}

type Call struct {
	msg
	call
}

type call struct {
	rpcvers uint
	prog    uint
	vers    uint
	proc    uint
	cred
	verf
}

type cred struct {
}

type verf struct {
}

type Reply struct {
	msg
	reply
}

type reply interface {
}

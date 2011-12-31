package rpc

type msg_type int

const (
	CALL  msg_type = 0
	REPLY msg_type = 1
)

const (
	MSG_ACCEPTED = iota
	MSG_DENIED
)

const (
	SUCCESS = iota
	PROG_UNAVAIL
	PROG_MISMATCH
	PROC_UNAVAIL
	GARBAGE_ARGS
)

const (
	RPC_MISMATCH = iota
	AUTH_ERROR
)

type auth_stat int

const (
	AUTH_BADCRED auth_stat = iota
	AUTH_REJECTEDCRED
	AUTH_BADVERF
	AUTH_REJECTEDVERF
	AUTH_TOOWEAK
)

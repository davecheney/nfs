package rpc

type msg_type int

const (
	CALL msg_type = 0
	REPLY msg_type = 1
)

type reply_stat int

const (
	MSG_ACCEPTED reply_stat = iota
	MSG_DENIED 
)

type accept_stat int

const (
	SUCCESS accept_stat = iota
	PROG_UNAVAIL 
	PROG_MISMATCH 
	PROC_UNAVAIL 
	GARBAGE_ARGS 
)

type reject_stat int

const (
	RPC_MISMATCH reject_stat = iota
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
	

package nfs

// NFS version 3
// RFC 1813

const (
	NFS3_PROG = 100003
	NFS3_VERS = 3
)

// Error represents an unexpected I/O behavior.
type Error struct {
        ErrorString string
}

func (err *Error) Error() string { return err.ErrorString }


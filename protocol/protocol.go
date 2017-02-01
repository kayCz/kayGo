package protocol

import "io"

type Protocol interface {
	Fill(rw io.ReadWriter)
	Send(v interface{}) error
	Receive() (v interface{}, err error)
}

package protocol

import "io"

type ProxifyEncoderDecoder struct {
	ProxifyDecoder
	ProxifyEncoder
}

func NewEncoderDecoder(rw io.ReadWriter) ProxifyEncoderDecoder {
	return ProxifyEncoderDecoder{
		NewDecoder(rw),
		NewEncoder(rw),
	}
}

package multibase

import (
	multi "github.com/multiformats/go-multibase"
)

type Multibase interface {
	Encode(base multi.Encoding, data []byte) (string, error)
	Decode(data string) (multi.Encoding, []byte, error)
}

func Encode(base multi.Encoding, data []byte) (string, error) {

	return multi.Encode(base, data)
}

func Decode(data string) (multi.Encoding, []byte, error) {

	return multi.Decode(data)
}

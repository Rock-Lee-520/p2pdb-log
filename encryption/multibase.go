package encryption

import (
	multi "github.com/multiformats/go-multibase"
)

type Multibase interface {
	BaseEncode(base multi.Encoding, data []byte) (string, error)
	BaseDecode(data string) (multi.Encoding, []byte, error)
}

func BaseEncode(base multi.Encoding, data []byte) (string, error) {

	return multi.Encode(base, data)
}

func BaseDecode(data string) (multi.Encoding, []byte, error) {

	return multi.Decode(data)
}

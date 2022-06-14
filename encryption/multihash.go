package encryption

import (
	multi "github.com/multiformats/go-multihash"
)

type Multihash interface {
	HashEncode(buf []byte, code uint64) ([]byte, error)
	HashDecode(buf []byte) (*multi.DecodedMultihash, error)
}

// Encode a hash digest along with the specified function code.
// Note: the length is derived from the length of the digest itself.
//
// The error return is legacy; it is always nil.
func HashEncode(buf []byte, code uint64) ([]byte, error) {

	return multi.Encode(buf, code)
}

// Decode parses multihash bytes into a DecodedMultihash.
func HashDecode(buf []byte) (*multi.DecodedMultihash, error) {

	return multi.Decode(buf)
}

package encryption

import (
	"testing"

	debug "github.com/favframework/debug"
	cid "github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

func TestGetCid(t *testing.T) {

	//   // And then feed it some data
	var str = "this is a test"
	c, err := GetCid(str)
	if err != nil {
		t.Error(c)
	}
	debug.Dump(c.String())
}

func TestGetCidString(t *testing.T) {

	//   // And then feed it some data
	var str = "this is a test"
	str, err := GetCidString(str)
	if err != nil {
		t.Error(err)
	}
	debug.Dump(str)
}

func TestV0Builder(t *testing.T) {
	data := []byte("this is some test content")

	// Construct c1
	format := cid.V0Builder{}
	c1, err := format.Sum(data)

	if err != nil {
		t.Fatal(err)
	}
	debug.Dump(c1.String())
	// Construct c2
	hash, err := mh.Sum(data, mh.SHA2_256, -1)
	debug.Dump(hash)
	if err != nil {
		t.Fatal(err)
	}
	c2 := cid.NewCidV0(hash)
	debug.Dump(c2.String())
	if !c1.Equals(c2) {
		t.Fatal("cids mismatch")
	}
	if c1.Prefix() != c2.Prefix() {
		t.Fatal("prefixes mismatch")
	}
}

package main

import (
	"fmt"

	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	// Create a cid manually by specifying the 'prefix' parameters
	pref := cid.Prefix{
		Version:  1,
		Codec:    mc.Raw,
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}
	//pref := cid.Prefix()
	// And then feed it some data
	c, err := pref.Sum([]byte("Hello World!"))
	if err != nil {

	}

	fmt.Println("Created CID: ", c)
}

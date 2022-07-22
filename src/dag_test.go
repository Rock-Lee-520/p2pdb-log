package src

import (
	"fmt"
	"os"
	"testing"

	"github.com/ipfs/go-cid"
	// _ "github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	"github.com/ipld/go-ipld-prime/linking"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/storage/memstore"
)

func TestDag(t *testing.T) {
	type Person struct {
		Name    string
		Age     int64 // TODO: optional to match other examples
		Friends []string
	}
	person := &Person{
		Name:    "Michael",
		Friends: []string{"Sarah", "Alex"},
	}
	node := bindnode.Wrap(person, nil)

	nodeRepr := node.Representation()
	dagjson.Encode(nodeRepr, os.Stdout)
}

func TestBuildMap(t *testing.T) {
	n, err := qp.BuildMap(basicnode.Prototype.Any, 4, func(ma datamodel.MapAssembler) {
		qp.MapEntry(ma, "some key", qp.String("some value"))
		qp.MapEntry(ma, "another key", qp.String("another value"))
		qp.MapEntry(ma, "nested map", qp.Map(2, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "deeper entries", qp.String("deeper values"))
			qp.MapEntry(ma, "more deeper entries", qp.String("more deeper values"))
		}))
		qp.MapEntry(ma, "nested list", qp.List(2, func(la datamodel.ListAssembler) {
			qp.ListEntry(la, qp.Int(1))
			qp.ListEntry(la, qp.Int(2))
		}))
	})
	if err != nil {
		panic(err)
	}
	dagjson.Encode(n, os.Stdout)
}

func TestKeyValue(t *testing.T) {
	np := basicnode.Prototype.Any // Pick a prototype: this is how we decide what implementation will store the in-memory data.
	nb := np.NewBuilder()         // Create a builder.
	ma, _ := nb.BeginMap(2)       // Begin assembling a map.
	ma.AssembleKey().AssignString("hey")
	ma.AssembleValue().AssignString("it works!")
	ma.AssembleKey().AssignString("yes")
	ma.AssembleValue().AssignBool(true)
	ma.Finish()     // Call 'Finish' on the map assembly to let it know no more data is coming.
	n := nb.Build() // Call 'Build' to get the resulting Node.  (It's immutable!)

	dagjson.Encode(n, os.Stdout)
}

type Memory struct {
	Bag map[string][]byte
}

func TestLink(t *testing.T) {
	// Creating a Link is done by choosing a concrete link implementation (typically, CID),
	//  getting a LinkSystem that knows how to work with that, and then using the LinkSystem methods.

	// Let's get a LinkSystem.  We're going to be working with CID links,
	//  so let's get the default LinkSystem that's ready to work with those.
	lsys := cidlink.DefaultLinkSystem()

	var store = memstore.Store{}

	// We want to store the serialized data somewhere.
	//  We'll use an in-memory store for this.  (It's a package scoped variable.)
	//  You can use any kind of storage system here;
	//   or if you need even more control, you could also write a function that conforms to the linking.BlockWriteOpener interface.
	lsys.SetWriteStorage(&store)

	// To create any links, first we need a LinkPrototype.
	// This gathers together any parameters that might be needed when making a link.
	// (For CIDs, the version, the codec, and the multihash type are all parameters we'll need.)
	// Often, you can probably make this a constant for your whole application.
	lp := cidlink.LinkPrototype{Prefix: cid.Prefix{
		Version:  1,    // Usually '1'.
		Codec:    0x71, // 0x71 means "dag-cbor" -- See the multicodecs table: https://github.com/multiformats/multicodec/
		MhType:   0x13, // 0x20 means "sha2-512" -- See the multicodecs table: https://github.com/multiformats/multicodec/
		MhLength: 64,   // sha2-512 hash has a 64-byte sum.
	}}

	// And we need some data to link to!  Here's a quick piece of example data:
	n := fluent.MustBuildMap(basicnode.Prototype.Map, 1, func(na fluent.MapAssembler) {
		na.AssembleEntry("hello").AssignString("world")
	})

	// Before we use the LinkService, NOTE:
	//  There's a side-effecting import at the top of the file.  It's for the dag-cbor codec.
	//  The CID LinkSystem defaults use a global registry called the multicodec table;
	//  and the multicodec table is populated in part by the dag-cbor package when it's first imported.
	// You'll need that side-effecting import, too, to copy this example.
	//  It can happen anywhere in your program; once, in any package, is enough.
	//  If you don't have this import, the codec will not be registered in the multicodec registry,
	//  and when you use the LinkSystem we got from the cidlink package, it will return an error of type ErrLinkingSetup.
	// If you initialize a custom LinkSystem, you can control this more directly;
	//  these registry systems are only here as defaults.

	// Now: time to apply the LinkSystem, and do the actual store operation!
	lnk, err := lsys.Store(
		linking.LinkContext{}, // The zero value is fine.  Configure it it you want cancellability or other features.
		lp,                    // The LinkPrototype says what codec and hashing to use.
		n,                     // And here's our data.
	)
	if err != nil {
		panic(err)
	}

	// That's it!  We got a link.
	fmt.Printf("link: %s\n", lnk)
	fmt.Printf("concrete type: `%T`\n", lnk)

	// Remember: the serialized data was also stored to the 'store' variable as a side-effect.
	//  (We set this up back when we customized the LinkSystem.)
	//  We'll pick this data back up again in the example for loading.

	// Output:
	// link: bafyrgqhai26anf3i7pips7q22coa4sz2fr4gk4q4sqdtymvvjyginfzaqewveaeqdh524nsktaq43j65v22xxrybrtertmcfxufdam3da3hbk
	// concrete type: `cidlink.Link`

}

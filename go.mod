module github.com/Rock-liyi/p2pdb-log

go 1.16

require (
	berty.tech/go-ipfs-log v1.8.0
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/caarlos0/env/v6 v6.9.3
	github.com/favframework/debug v0.0.0-20150708094948-5c7e73aafb21
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-datastore v0.5.1
	github.com/ipfs/go-ipfs v0.12.2
	github.com/ipfs/interface-go-ipfs-core v0.5.2
	github.com/ipld/go-ipld-prime v0.14.2
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.4.0
	github.com/libp2p/go-libp2p v0.17.0
	github.com/libp2p/go-libp2p-core v0.13.0
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multicodec v0.3.0
	github.com/multiformats/go-multihash v0.1.0
	github.com/multiformats/go-varint v0.0.6
	github.com/stretchr/testify v1.7.1
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.7

)

replace github.com/Rock-liyi/p2pdb-log => ../p2pdb-log

// Package keystore defines a local key manager for OrbitDB and IPFS Log.
package keystore // import "github.com/Rock-liyi/p2pdb-log/keystore"

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	lru "github.com/hashicorp/golang-lru"
	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p-core/crypto"

	"github.com/Rock-liyi/p2pdb-log/message"
)

type Keystore struct {
	store datastore.Datastore
	cache *lru.Cache
}

// Sign signs a value using a given private key.
func (k *Keystore) Sign(privKey crypto.PrivKey, bytes []byte) ([]byte, error) {
	return privKey.Sign(bytes)
}

// Verify verifies a signature.
func (k *Keystore) Verify(signature []byte, publicKey crypto.PubKey, data []byte) error {
	ok, err := publicKey.Verify(data, signature)
	if err != nil {
		return message.ErrSigNotVerified.Wrap(err)
	}

	if !ok {
		return message.ErrSigNotVerified
	}

	return nil
}

// NewKeystore creates a new keystore.
func NewKeystore(store datastore.Datastore) (*Keystore, error) {
	cache, err := lru.New(128)
	if err != nil {
		return nil, message.ErrKeyStoreInitFailed.Wrap(err)
	}

	return &Keystore{
		store: store,
		cache: cache,
	}, nil
}

// HasKey checks whether a given key ID exist in the keystore.
func (k *Keystore) HasKey(ctx context.Context, id string) (bool, error) {
	storedKey, ok := k.cache.Peek(id)

	if ok == false {
		value, err := k.store.Get(ctx, datastore.NewKey(id))
		if err != nil {
			return false, message.ErrKeyNotInKeystore.Wrap(err)
		}

		if storedKey != nil {
			k.cache.Add(id, base64.StdEncoding.EncodeToString(value))
		}
	}

	return storedKey != nil, nil
}

// CreateKey creates a new key in the key store.
func (k *Keystore) CreateKey(ctx context.Context, id string) (crypto.PrivKey, error) {
	// FIXME: I kept Secp256k1 for compatibility with OrbitDB, should we change this?
	priv, _, err := crypto.GenerateSecp256k1Key(rand.Reader)
	if err != nil {
		return nil, message.ErrKeyGenerationFailed.Wrap(err)
	}

	keyBytes, err := priv.Raw()
	if err != nil {
		return nil, message.ErrInvalidPrivKeyFormat.Wrap(err)
	}

	if err := k.store.Put(ctx, datastore.NewKey(id), keyBytes); err != nil {
		return nil, message.ErrKeyStorePutFailed.Wrap(err)
	}

	k.cache.Add(id, base64.StdEncoding.EncodeToString(keyBytes))

	return priv, nil
}

// GetKey retrieves a key from the keystore.
func (k *Keystore) GetKey(ctx context.Context, id string) (crypto.PrivKey, error) {
	var err error
	var keyBytes []byte

	cachedKey, ok := k.cache.Get(id)
	if !ok || cachedKey == nil {
		keyBytes, err = k.store.Get(ctx, datastore.NewKey(id))

		if err != nil {
			return nil, message.ErrKeyNotInKeystore.Wrap(err)
		}
		k.cache.Add(id, base64.StdEncoding.EncodeToString(keyBytes))
	} else {
		keyBytes, err = base64.StdEncoding.DecodeString(cachedKey.(string))
		if err != nil {
			return nil, message.ErrInvalidPrivKeyFormat.Wrap(err)
		}
	}

	privateKey, err := crypto.UnmarshalSecp256k1PrivateKey(keyBytes)
	if err != nil {
		return nil, message.ErrInvalidPrivKeyFormat.Wrap(err)
	}

	return privateKey, nil
}

var _ Interface = &Keystore{}

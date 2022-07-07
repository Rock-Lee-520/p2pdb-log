package identityprovider // import "github.com/Rock-liyi/p2pdb-log/identityprovider"

import (
	"context"
	"encoding/hex"

	"github.com/libp2p/go-libp2p-core/crypto"

	"github.com/Rock-liyi/p2pdb-log/keystore"
	"github.com/Rock-liyi/p2pdb-log/message"
)

type OrbitDBIdentityProvider struct {
	keystore keystore.Interface
}

// VerifyIdentity checks an OrbitDB identity.
func (p *OrbitDBIdentityProvider) VerifyIdentity(identity *Identity) error {
	return nil
}

// NewOrbitDBIdentityProvider creates a new identity for use with OrbitDB.
func NewOrbitDBIdentityProvider(options *CreateIdentityOptions) Interface {
	return &OrbitDBIdentityProvider{
		keystore: options.Keystore,
	}
}

// GetID returns the identity's ID.
func (p *OrbitDBIdentityProvider) GetID(ctx context.Context, options *CreateIdentityOptions) (string, error) {
	private, err := p.keystore.GetKey(ctx, options.ID)
	if err != nil || private == nil {
		private, err = p.keystore.CreateKey(ctx, options.ID)
		if err != nil {
			return "", message.ErrKeyStoreCreateEntry.Wrap(err)
		}
	}

	pubBytes, err := private.GetPublic().Raw()
	if err != nil {
		return "", message.ErrPubKeySerialization.Wrap(err)
	}

	return hex.EncodeToString(pubBytes), nil
}

// SignIdentity signs an OrbitDB identity.
func (p *OrbitDBIdentityProvider) SignIdentity(ctx context.Context, data []byte, id string) ([]byte, error) {
	key, err := p.keystore.GetKey(ctx, id)
	if err != nil {
		return nil, message.ErrKeyNotInKeystore
	}

	//data, _ = hex.DecodeString(hex.EncodeToString(data))

	// FIXME? Data is a unicode encoded hex as a byte (source lib uses Buffer.from(hexStr) instead of Buffer.from(hexStr, "hex"))
	data = []byte(hex.EncodeToString(data))

	signature, err := key.Sign(data)
	if err != nil {
		return nil, message.ErrSigSign.Wrap(err)
	}

	return signature, nil
}

// Sign signs a value using the current.
func (p *OrbitDBIdentityProvider) Sign(ctx context.Context, identity *Identity, data []byte) ([]byte, error) {
	key, err := p.keystore.GetKey(ctx, identity.ID)
	if err != nil {
		return nil, message.ErrKeyNotInKeystore.Wrap(err)
	}

	sig, err := key.Sign(data)
	if err != nil {
		return nil, message.ErrSigSign.Wrap(err)
	}

	return sig, nil
}

func (p *OrbitDBIdentityProvider) UnmarshalPublicKey(data []byte) (crypto.PubKey, error) {
	pubKey, err := crypto.UnmarshalSecp256k1PublicKey(data)
	if err != nil {
		return nil, message.ErrInvalidPubKeyFormat
	}

	return pubKey, nil
}

// GetType returns the current identity type.
func (*OrbitDBIdentityProvider) GetType() string {
	return "orbitdb"
}

var _ Interface = &OrbitDBIdentityProvider{}

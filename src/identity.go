package src

import (
	ic "github.com/libp2p/go-libp2p-core/crypto"
)

type IdentitySignature struct {
	ID        []byte `json:"id,omitempty"`
	PublicKey []byte `json:"publicKey,omitempty"`
}

type Identity struct {
	ID         string             `json:"id,omitempty"`
	PublicKey  []byte             `json:"publicKey,omitempty"`
	Signatures *IdentitySignature `json:"signatures,omitempty"`
	Type       string             `json:"type,omitempty"`
}

// Filtered gets fields that should be present in the CBOR representation of identity.
func (i *Identity) Filtered() *Identity {
	return &Identity{
		ID:         i.ID,
		PublicKey:  i.PublicKey,
		Signatures: i.Signatures,
		Type:       i.Type,
	}
}

// GetPublicKey returns the public key of an identity.
func (i *Identity) GetPublicKey() (ic.PubKey, error) {
	return ic.UnmarshalPublicKey(i.PublicKey)
}

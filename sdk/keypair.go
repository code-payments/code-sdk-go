package codesdk

import (
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/mr-tron/base58"
)

type KeyPair struct {
	publicKey  *PublicKey
	privateKey *PrivateKey
}

func GenerateKeyPair() (*KeyPair, error) {
	_, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	return NewKeyPairFromRawValue(privateKey)
}

func NewKeyPairFromRawValue(rawValue ed25519.PrivateKey) (*KeyPair, error) {
	privateKey, err := NewPrivateKeyFromBytes(rawValue)
	if err != nil {
		return nil, err
	}

	return &KeyPair{
		publicKey:  privateKey.GetPublicKey(),
		privateKey: privateKey,
	}, nil
}

func NewKeyPairFromSeed(seed []byte) (*KeyPair, error) {
	privateKey := ed25519.NewKeyFromSeed(seed)
	return NewKeyPairFromRawValue(privateKey)
}

func (k *KeyPair) GetPublicKey() *PublicKey {
	return k.publicKey
}

func (k *KeyPair) GetPrivateKey() *PrivateKey {
	return k.privateKey
}

func (k *KeyPair) Sign(message []byte) []byte {
	return k.privateKey.Sign(message)
}

func (k *KeyPair) Verify(message, signature []byte) bool {
	return ed25519.Verify(k.publicKey.ToBytes(), message, signature)
}

type PublicKey struct {
	bytesValue  ed25519.PublicKey
	stringValue string
}

func NewPublicKeyFromBytes(value []byte) (*PublicKey, error) {
	key := &PublicKey{
		bytesValue:  value,
		stringValue: base58.Encode(value),
	}

	if err := key.validate(); err != nil {
		return nil, err
	}

	return key, nil
}

func NewPublicKeyFromString(value string) (*PublicKey, error) {
	decoded, err := base58.Decode(value)
	if err != nil {
		return nil, err
	}

	key := &PublicKey{
		bytesValue:  decoded,
		stringValue: value,
	}

	if err := key.validate(); err != nil {
		return nil, err
	}

	return key, nil
}

func (k *PublicKey) ToBytes() ed25519.PublicKey {
	return k.bytesValue
}

func (k *PublicKey) ToBase58() string {
	return k.stringValue
}

func (k *PublicKey) validate() error {
	if len(k.bytesValue) != ed25519.PublicKeySize {
		return fmt.Errorf("key must be an ed25519 public key of size %d", ed25519.PublicKeySize)
	}

	if base58.Encode(k.bytesValue) != k.stringValue {
		return errors.New("bytes and string representation don't match")
	}

	return nil
}

type PrivateKey struct {
	bytesValue  ed25519.PrivateKey
	stringValue string
}

func NewPrivateKeyFromBytes(value []byte) (*PrivateKey, error) {
	key := &PrivateKey{
		bytesValue:  value,
		stringValue: base58.Encode(value),
	}

	if err := key.validate(); err != nil {
		return nil, err
	}

	return key, nil
}

func NewPrivateKeyFromString(value string) (*PrivateKey, error) {
	decoded, err := base58.Decode(value)
	if err != nil {
		return nil, err
	}

	key := &PrivateKey{
		bytesValue:  decoded,
		stringValue: value,
	}

	if err := key.validate(); err != nil {
		return nil, err
	}

	return key, nil
}

func (k *PrivateKey) ToBytes() ed25519.PrivateKey {
	return k.bytesValue
}

func (k *PrivateKey) ToBase58() string {
	return k.stringValue
}

func (k *PrivateKey) GetPublicKey() *PublicKey {
	publicKey, _ := NewPublicKeyFromBytes(k.bytesValue.Public().(ed25519.PublicKey))
	return publicKey
}

func (k *PrivateKey) Sign(message []byte) []byte {
	return ed25519.Sign(k.bytesValue, message)
}

func (k *PrivateKey) validate() error {
	if len(k.bytesValue) != ed25519.PrivateKeySize {
		return fmt.Errorf("key must be an ed25519 private key of size %d", ed25519.PrivateKeySize)
	}

	if base58.Encode(k.bytesValue) != k.stringValue {
		return errors.New("bytes and string representation don't match")
	}

	return nil
}

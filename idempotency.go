package codesdk

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"github.com/mr-tron/base58"
)

type IdempotencyKey [codePayloadNonceSize]byte

func GenerateIdempotencyKey() IdempotencyKey {
	var res IdempotencyKey
	rand.Read(res[:])
	return res
}

func NewIdempotencyKeyFromSeed(seed string) IdempotencyKey {
	var res IdempotencyKey
	hashed := sha256.Sum256([]byte(seed))
	copy(res[:], hashed[:])
	return res
}

func NewIdempotencyKeyFromClientSecret(data string) (IdempotencyKey, error) {
	var res IdempotencyKey

	decoded, err := base58.Decode(data)
	if err != nil {
		return res, err
	}

	if len(decoded) != codePayloadNonceSize {
		return res, errors.New("invalid nonce size")
	}

	copy(res[:], decoded[:])

	return res, nil
}

func (k IdempotencyKey) String() string {
	return base58.Encode(k[:])
}

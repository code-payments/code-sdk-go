package codesdk

import (
	"crypto/sha256"
)

func GenerateRendezvousKey(p *CodePayload) (*KeyPair, error) {
	seed := sha256.Sum256(p.toBytes())
	return NewKeyPairFromSeed(seed[:])
}

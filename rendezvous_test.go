package codesdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRendezvousKey(t *testing.T) {
	payload, err := NewCodePayload(CodePayloadPaymentRequest, USD, 1.00, testCodePayloadNonce)
	require.NoError(t, err)

	actual, err := GenerateRendezvousKey(payload)
	require.NoError(t, err)

	assert.Equal(t, "2ezSqLgSwi3iQMx4LvWdYviMchrxq4C8iDjdyiYrfQeFLNFRxkiscEnFh8Zjk8MvCFDmeYveE8SuUpRDowjaMdcd", actual.GetPrivateKey().ToBase58())
	assert.Equal(t, "VF5ptmbp7UXqCqvQN7Hv8eSpUY7RybXkC2kaPNUFDwf", actual.GetPublicKey().ToBase58())
}

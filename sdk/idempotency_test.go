package codesdk

import (
	"testing"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateIdempotencyKey(t *testing.T) {
	seen := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		generated := GenerateIdempotencyKey()
		seen[base58.Encode(generated[:])] = struct{}{}
	}
	assert.Len(t, seen, 100)
}

func TestNewIdempotencyKeyFromSeed(t *testing.T) {
	expected := []byte{
		0x4b, 0x64, 0x1e, 0x9a, 0x92,
		0x3d, 0x1e, 0xa5, 0x7e, 0x18,
		0xfe,
	}
	actual := NewIdempotencyKeyFromSeed("test_string")
	assert.EqualValues(t, expected, actual)
}

func TestNewIdempotencyKeyFromClientSecret(t *testing.T) {
	expected := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04,
		0x05, 0x06, 0x07, 0x08, 0x09,
		0x0a,
	}
	clientSecret := base58.Encode(expected)

	actual, err := NewIdempotencyKeyFromClientSecret(clientSecret)
	require.NoError(t, err)

	assert.EqualValues(t, expected, actual)

	_, err = NewIdempotencyKeyFromClientSecret("invalid")
	assert.Error(t, err)
}

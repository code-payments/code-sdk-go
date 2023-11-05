package codesdk

import (
	"testing"

	"github.com/mr-tron/base58"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateKeyPair(t *testing.T) {
	seen := make(map[string]struct{})
	for i := 0; i < 100; i++ {
		generated, err := GenerateKeyPair()
		require.NoError(t, err)
		seen[generated.GetPrivateKey().ToBase58()] = struct{}{}
	}
	assert.Len(t, seen, 100)
}

func TestNewKeyPairFromSeed(t *testing.T) {
	seed := make([]byte, 32)
	for i := 0; i < len(seed); i++ {
		seed[i] = 8
	}
	keypair, err := NewKeyPairFromSeed(seed)
	require.NoError(t, err)
	assert.Equal(t, "2KW2XRd9kwqet15Aha2oK3tYvd3nWbTFH1MBiRAv1BE1", keypair.GetPublicKey().ToBase58())
}

func TestSignAndVerify(t *testing.T) {
	keypair1, err := GenerateKeyPair()
	require.NoError(t, err)
	keypair2, err := GenerateKeyPair()
	require.NoError(t, err)

	message1 := []byte("message1")
	message2 := []byte("message2")

	signature1 := keypair1.Sign(message1)
	signature2 := keypair2.Sign(message1)

	assert.NotEqualValues(t, signature1, signature2)

	assert.True(t, keypair1.Verify(message1, signature1))
	assert.False(t, keypair1.Verify(message1, signature2))
	assert.False(t, keypair1.Verify(message2, signature1))

	assert.True(t, keypair2.Verify(message1, signature2))
	assert.False(t, keypair2.Verify(message2, signature2))
	assert.False(t, keypair2.Verify(message1, signature1))
}

func TestPublicKeyConversions(t *testing.T) {
	expectedStringValue := "CiDwVBFgWV9E5MvXWoLgnEgn2hK7rJikbvfWavzAQz3"
	expectedBytesValue, _ := base58.Decode(expectedStringValue)

	publicKey, err := NewPublicKeyFromBytes(expectedBytesValue)
	require.NoError(t, err)
	assert.Equal(t, expectedStringValue, publicKey.ToBase58())
	assert.EqualValues(t, expectedBytesValue, publicKey.ToBytes())

	publicKey, err = NewPublicKeyFromString(expectedStringValue)
	require.NoError(t, err)
	assert.Equal(t, expectedStringValue, publicKey.ToBase58())
	assert.EqualValues(t, expectedBytesValue, publicKey.ToBytes())
}

func TestPrivateKeyConversions(t *testing.T) {
	expectedPrivateKeyStringValue := "5R8n6JLJa6uLAt7xri9Yk7iZmbDr6CpgxBiWojePsAwijD1Nbzjp15XebjcSuqe7pFdfyfb77W4dmSArMfvgPigE"
	expectedPrivateKeyBytesValue, _ := base58.Decode(expectedPrivateKeyStringValue)

	expectedPublicKeyStringValue := "pWiwCcUwumuYbFNJNu2FJ8u6btrfNLzS1iVaMfA6HFc"

	privateKey, err := NewPrivateKeyFromBytes(expectedPrivateKeyBytesValue)
	require.NoError(t, err)
	assert.Equal(t, expectedPrivateKeyStringValue, privateKey.ToBase58())
	assert.EqualValues(t, expectedPrivateKeyBytesValue, privateKey.ToBytes())
	assert.Equal(t, expectedPublicKeyStringValue, privateKey.GetPublicKey().ToBase58())

	privateKey, err = NewPrivateKeyFromString(expectedPrivateKeyStringValue)
	require.NoError(t, err)
	assert.Equal(t, expectedPrivateKeyStringValue, privateKey.ToBase58())
	assert.EqualValues(t, expectedPrivateKeyBytesValue, privateKey.ToBytes())
	assert.Equal(t, expectedPublicKeyStringValue, privateKey.GetPublicKey().ToBase58())
}

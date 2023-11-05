package codesdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCodePayloadNonce = IdempotencyKey{
	0x01, 0x02, 0x03, 0x04, 0x05,
	0x06, 0x07, 0x08, 0x09, 0x10,
	0x11,
}

func TestCodePayload_PaymentRequest_Fiat_Encoding(t *testing.T) {
	expected := []byte{
		0x02, 0x8c, 0xFF, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x01,
		0x02, 0x03, 0x04, 0x05, 0x06,
		0x07, 0x08, 0x09, 0x10, 0x11,
	}

	payload, err := NewCodePayload(CodePayloadPaymentRequest, USD, 2814749767109.11, testCodePayloadNonce)
	require.NoError(t, err)

	assert.EqualValues(t, expected, payload.toBytes())
}

func TestCodePayload_PaymentRequest_Kin_Encoding(t *testing.T) {
	expected := []byte{
		0x02, 0x00, 0x88, 0x13, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01,
		0x02, 0x03, 0x04, 0x05, 0x06,
		0x07, 0x08, 0x09, 0x10, 0x11,
	}

	payload, err := NewCodePayload(CodePayloadPaymentRequest, KIN, 50.0, testCodePayloadNonce)
	require.NoError(t, err)

	assert.EqualValues(t, expected, payload.toBytes())
}

func TestCodePayload_PaymentRequest_InvalidCurrency(t *testing.T) {
	_, err := NewCodePayload(CodePayloadPaymentRequest, "invalid", 1.00, testCodePayloadNonce)
	assert.Equal(t, ErrInvalidCurrency, err)
}

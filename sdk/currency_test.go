package codesdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsValidCurrency(t *testing.T) {
	assert.True(t, IsValidCurrency("kin"))
	assert.True(t, IsValidCurrency("usd"))
	assert.True(t, IsValidCurrency("eur"))

	assert.False(t, IsValidCurrency("invalid"))
}

func TestCurrencyCodeToIndex(t *testing.T) {
	actual, err := KIN.toIndex()
	require.NoError(t, err)
	assert.Equal(t, 0, actual)

	actual, err = USD.toIndex()
	require.NoError(t, err)
	assert.Equal(t, 140, actual)

	actual, err = EUR.toIndex()
	require.NoError(t, err)
	assert.Equal(t, 43, actual)

	_, err = CurrencyCode("invalid").toIndex()
	assert.Equal(t, ErrInvalidCurrency, err)
}

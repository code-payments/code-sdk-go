package codesdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentRequestIntent_Rounding(t *testing.T) {
	for _, tc := range []struct {
		input    float64
		expected float64
	}{
		{0.05, 0.05},
		{0.054999999, 0.05},
		{0.055, 0.06},
		{0.055000001, 0.06},
		{0.06, 0.06},
	} {
		intent, err := NewPaymentRequestIntent(USD, tc.input, "E8otxw1CVX9bfyddKu3ZB3BVLa4VVF9J7CTPdnUwT9jR")
		require.NoError(t, err)
		assert.Equal(t, tc.expected, intent.convertedAmount)
	}
}

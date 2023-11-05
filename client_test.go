package codesdk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentRequestIntent(t *testing.T) {
	destination := "E8otxw1CVX9bfyddKu3ZB3BVLa4VVF9J7CTPdnUwT9jR"
	for _, tc := range []struct {
		currency CurrencyCode
		amount   float64
	}{
		{USD, 0.05},
		{KIN, 10_000.00},
	} {
		ctx := context.Background()

		intent, err := NewPaymentRequestIntent(tc.currency, tc.amount, destination)
		require.NoError(t, err)

		client := NewWebClient()
		resp, err := client.CreatePaymentRequest(ctx, intent)
		require.NoError(t, err)

		idempotencyKey, err := NewIdempotencyKeyFromClientSecret(resp.ClientSecret)
		require.NoError(t, err)

		intent, err = NewPaymentRequestIntent(tc.currency, tc.amount, destination, WithIdempotencyKey(idempotencyKey))
		require.NoError(t, err)
		assert.Equal(t, intent.GetIntentId(), resp.IntentId)
	}
}

func TestGetIntentStatus(t *testing.T) {
	ctx := context.Background()
	client := NewWebClient()
	for _, tc := range []struct {
		intentId string
		expected IntentState
	}{
		{"9Nkao1TNKdjjRcLayHeZwWnHCJrbYxq7iL3Zu8TxT5fX", IntentStateUnknown},
		{"1RfVoYZQ1jNq5HdVq66xgERURaJgmBrkUgYd6vHXkhk", IntentStatePending},
		{"395Facg6FY1wZyG7vSdUg2R3gZJtkQqxz89G3oowzGiN", IntentStateConfirmed},
	} {
		actual, err := client.GetIntentStatus(ctx, tc.intentId)
		require.NoError(t, err)
		assert.Equal(t, tc.expected, actual)
	}
}

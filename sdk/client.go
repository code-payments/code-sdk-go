package codesdk

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mr-tron/base58"
	"google.golang.org/protobuf/proto"
)

const (
	v1ApiBaseUrl    = "https://api.getcode.com/v1/"
	createIntentUrl = v1ApiBaseUrl + "createIntent"
	getStatusUrl    = v1ApiBaseUrl + "getStatus"
)

// The state of an intent
type IntentState string

const (
	// The intent doesn't exist
	IntentStateUnknown IntentState = "unknown"
	// The intent exists, but the user hasn't submitted a payment
	IntentStatePending IntentState = "pending"
	// The user has submitted a payment. Fulfillment on the blockchain is either
	// in progress, or completed, by the Code sequencer.
	IntentStateConfirmed IntentState = "confirmed"
)

type Client struct {
	httpClient *http.Client
}

func NewWebClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

type CreatePaymentRequestResponse struct {
	IntentId     string `json:"id"`
	ClientSecret string `json:"clientSecret"`
}

// CreatePaymentRequest creates a payment request intent. The response object
// can be used directly as the return value for the Code SDK on the browser.
func (c *Client) CreatePaymentRequest(
	ctx context.Context,
	intent *PaymentRequestIntent,
	opts ...PaymentRequestOption,
) (*CreatePaymentRequestResponse, error) {
	optionalParameters := applyPaymentRequestOptions(opts...)

	protoMessage, err := proto.Marshal(intent.toProtoMessage())
	if err != nil {
		return nil, err
	}

	signature, err := intent.sign()
	if err != nil {
		return nil, err
	}

	params := new(bytes.Buffer)
	err = json.NewEncoder(params).Encode(map[string]any{
		"intent":    intent.GetIntentId(),
		"message":   base64.RawURLEncoding.EncodeToString(protoMessage),
		"signature": base58.Encode(signature),
		"webhook":   optionalParameters.webhookUrl,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, createIntentUrl, params)
	if err != nil {
		return nil, err
	}

	_, err = c.do(ctx, req)
	if err != nil {
		return nil, err
	}

	return &CreatePaymentRequestResponse{
		IntentId:     intent.GetIntentId(),
		ClientSecret: intent.GetClientSecret(),
	}, nil
}

type PaymentRequestOption func(*optionalPaymentRequestParameters)

func WithWebhookUrl(webhookUrl string) PaymentRequestOption {
	return func(opts *optionalPaymentRequestParameters) {
		opts.webhookUrl = &webhookUrl
	}
}

type optionalPaymentRequestParameters struct {
	webhookUrl *string
}

func applyPaymentRequestOptions(opts ...PaymentRequestOption) *optionalPaymentRequestParameters {
	res := &optionalPaymentRequestParameters{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

type GetIntentStateResponse struct {
	Status IntentState `json:"status"`
}

// GetIntentStatus returns the state of the intent
func (c *Client) GetIntentStatus(ctx context.Context, intentId string) (*GetIntentStateResponse, error) {
	url := fmt.Sprintf("%s?intent=%s", getStatusUrl, intentId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil
	}

	body, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}

	jsonRespBody := struct {
		Status string `json:"status"`
	}{}
	err = json.Unmarshal(body, &jsonRespBody)
	if err != nil {
		return nil, err
	}

	var state IntentState
	switch strings.ToLower(jsonRespBody.Status) {
	case "submitted":
		state = IntentStateConfirmed
	case "pending":
		state = IntentStatePending
	default:
		state = IntentStateUnknown
	}
	return &GetIntentStateResponse{
		Status: state,
	}, nil
}

func (c *Client) do(ctx context.Context, req *http.Request) ([]byte, error) {
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var isCodeApiResponse bool
	jsonRespBody := struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}{}
	err = json.Unmarshal(body, &jsonRespBody)
	if err == nil {
		isCodeApiResponse = true
	}

	if resp.StatusCode != http.StatusOK {
		errorDescription := string(body)
		if isCodeApiResponse {
			errorDescription = jsonRespBody.Error
		}
		return nil, fmt.Errorf("https status %d: %s", resp.StatusCode, errorDescription)
	}

	if !jsonRespBody.Success {
		return nil, errors.New(jsonRespBody.Error)
	}

	return body, nil
}

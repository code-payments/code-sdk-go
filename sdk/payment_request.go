package codesdk

import (
	"math"

	"google.golang.org/protobuf/proto"

	codepb "github.com/code-wallet/code-sdk-go/genproto"
)

// PaymentRequestIntent is an intent to request a payment be made to a destination
// for a specifc amount. For fiat values, exchange rates are computed dynamically
// at time of payment.
type PaymentRequestIntent struct {
	currency        CurrencyCode
	convertedAmount float64
	destination     *PublicKey
	nonce           IdempotencyKey
	rendezvousKey   *KeyPair
}

func NewPaymentRequestIntent(
	currency CurrencyCode,
	amount float64,
	destination string,
	opts ...IntentOption,
) (*PaymentRequestIntent, error) {
	optionalParameters := applyIntentOptions(opts...)

	convertedAmount := math.Round(100.0*amount) / 100.0

	destinationPublicKey, err := NewPublicKeyFromString(destination)
	if err != nil {
		return nil, err
	}

	payload, err := NewCodePayload(
		CodePayloadPaymentRequest,
		currency,
		convertedAmount,
		optionalParameters.idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	rendezvousKey, err := GenerateRendezvousKey(payload)
	if err != nil {
		return nil, err
	}

	return &PaymentRequestIntent{
		currency:        currency,
		convertedAmount: convertedAmount,
		destination:     destinationPublicKey,
		nonce:           optionalParameters.idempotencyKey,
		rendezvousKey:   rendezvousKey,
	}, nil
}

// GetIntentId is the unique ID for the intent. It is the public key of the
// rendezvous key pair.
func (p *PaymentRequestIntent) GetIntentId() string {
	return p.rendezvousKey.GetPublicKey().ToBase58()
}

// GetRendezvousKey returns a unique key pair for the scan code payload for
// the intent, which is used during the scanning process to establish a secure
// communication channel anonymously to coordinate a flow.
func (p *PaymentRequestIntent) GetRendezvousKey() *KeyPair {
	return p.rendezvousKey
}

// GetClientSecret returns a secret value required by the Code SDK at the
// browser to reconstruct the intent. Your server should never share this
// value until the intent is successfully created against Code server.
func (p *PaymentRequestIntent) GetClientSecret() string {
	return p.nonce.String()
}

func (p *PaymentRequestIntent) toProtoMessage() *codepb.RequestToReceiveBill {
	msg := &codepb.RequestToReceiveBill{
		RequestorAccount: &codepb.SolanaAccountId{
			Value: p.destination.ToBytes(),
		},
	}

	if p.currency == KIN {
		msg.ExchangeData = &codepb.RequestToReceiveBill_Exact{
			Exact: &codepb.ExchangeData{
				Currency:     string(p.currency),
				ExchangeRate: 1.0,
				NativeAmount: p.convertedAmount,
				Quarks:       uint64(p.convertedAmount * QuarksPerKin),
			},
		}
	} else {
		msg.ExchangeData = &codepb.RequestToReceiveBill_Partial{
			Partial: &codepb.ExchangeDataWithoutRate{
				Currency:     string(p.currency),
				NativeAmount: p.convertedAmount,
			},
		}
	}

	return msg
}

func (p *PaymentRequestIntent) sign() ([]byte, error) {
	envelope := &codepb.Envelope{
		Kind: &codepb.Envelope_RequestToReceiveBill{
			RequestToReceiveBill: p.toProtoMessage(),
		},
	}

	marshalled, err := proto.Marshal(envelope)
	if err != nil {
		return nil, err
	}

	return p.rendezvousKey.Sign(marshalled), nil
}

type IntentOption func(*optionalIntentParameters)

func WithIdempotencyKey(idempotencyKey IdempotencyKey) IntentOption {
	return func(opts *optionalIntentParameters) {
		opts.idempotencyKey = idempotencyKey
	}
}

type optionalIntentParameters struct {
	idempotencyKey IdempotencyKey
}

func applyIntentOptions(opts ...IntentOption) *optionalIntentParameters {
	res := &optionalIntentParameters{
		idempotencyKey: GenerateIdempotencyKey(),
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

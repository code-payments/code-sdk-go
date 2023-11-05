package codesdk

import (
	"google.golang.org/protobuf/proto"

	codepb "github.com/code-wallet/code-sdk-go/genproto"
)

type PaymentRequestIntent struct {
	currency      CurrencyCode
	amount        float64
	destination   *PublicKey
	nonce         IdempotencyKey
	rendezvousKey *KeyPair
}

func NewPaymentRequestIntent(
	currency CurrencyCode,
	amount float64,
	destination *PublicKey,
	opts ...OptionalIntentParameters,
) (*PaymentRequestIntent, error) {
	optionalIntentParamters := applyOptionalIntentParameters(opts...)

	amount = float64(uint64(100*amount)) / 100.0

	payload, err := NewCodePayload(
		CodePayloadPaymentRequest,
		currency,
		amount,
		optionalIntentParamters.idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	rendezvousKey, err := GenerateRendezvousKey(payload)
	if err != nil {
		return nil, err
	}

	return &PaymentRequestIntent{
		currency:      currency,
		amount:        amount,
		destination:   destination,
		nonce:         optionalIntentParamters.idempotencyKey,
		rendezvousKey: rendezvousKey,
	}, nil
}

func (p *PaymentRequestIntent) GetIntentId() string {
	return p.rendezvousKey.GetPublicKey().ToBase58()
}

func (p *PaymentRequestIntent) GetClientSecret() string {
	return p.nonce.String()
}

func (p *PaymentRequestIntent) GetRendezvousKey() *KeyPair {
	return p.rendezvousKey
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
				NativeAmount: p.amount,
				Quarks:       uint64(p.amount * QuarksPerKin),
			},
		}
	} else {
		msg.ExchangeData = &codepb.RequestToReceiveBill_Partial{
			Partial: &codepb.ExchangeDataWithoutRate{
				Currency:     string(p.currency),
				NativeAmount: p.amount,
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

type optionalIntentParamters struct {
	idempotencyKey IdempotencyKey
}

type OptionalIntentParameters func(*optionalIntentParamters)

func WithIdempotencyKey(idempotencyKey IdempotencyKey) OptionalIntentParameters {
	return func(opts *optionalIntentParamters) {
		opts.idempotencyKey = idempotencyKey
	}
}

func applyOptionalIntentParameters(opts ...OptionalIntentParameters) *optionalIntentParamters {
	res := &optionalIntentParamters{
		idempotencyKey: GenerateIdempotencyKey(),
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

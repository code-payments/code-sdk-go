package codesdk

import (
	"errors"
)

const (
	codePayloadTypeSize   = 1
	codePayloadAmountSize = 8
	codePayloadNonceSize  = 11
	codePayloadSize       = codePayloadTypeSize + codePayloadAmountSize + codePayloadNonceSize
)

type CodePayloadType uint8

const (
	CodePayloadCash CodePayloadType = iota
	CodePayloadGiftCard
	CodePayloadPaymentRequest
)

// CodePayload is the payload format for scan codes.
type CodePayload struct {
	kind   CodePayloadType
	amount amountBuffer
	nonce  IdempotencyKey
}

func NewCodePayload(kind CodePayloadType, currency CurrencyCode, amount float64, nonce IdempotencyKey) (*CodePayload, error) {
	if kind != CodePayloadPaymentRequest {
		return nil, errors.New("only payment request codes are supported")
	}

	amountBuffer, err := newCurrencyAmountBuffer(currency, amount)
	if err != nil {
		return nil, err
	}

	return &CodePayload{
		kind:   kind,
		amount: amountBuffer,
		nonce:  nonce,
	}, nil
}

func (p *CodePayload) toBytes() []byte {
	var buffer [codePayloadSize]byte
	buffer[0] = byte(p.kind)

	amountBuffer := p.amount.toBytes()
	for i := 0; i < codePayloadAmountSize; i++ {
		buffer[i+codePayloadTypeSize] = amountBuffer[i]
	}

	for i := 0; i < codePayloadNonceSize; i++ {
		buffer[i+codePayloadTypeSize+codePayloadAmountSize] = p.nonce[i]
	}

	return buffer[:]
}

type amountBuffer interface {
	toBytes() [codePayloadAmountSize]byte
}

type currencyAmountBuffer struct {
	currency CurrencyCode
	amount   float64
}

func newCurrencyAmountBuffer(currency CurrencyCode, amount float64) (amountBuffer, error) {
	_, err := currency.toIndex()
	if err != nil {
		return nil, err
	}

	return &currencyAmountBuffer{
		currency: currency,
		amount:   amount,
	}, nil
}

func (b *currencyAmountBuffer) toBytes() [codePayloadAmountSize]byte {
	var buffer [codePayloadAmountSize]byte

	currencyIndex, _ := b.currency.toIndex()
	buffer[0] = byte(currencyIndex)

	amountToSerialize := uint64(100.0 * b.amount)
	for i := 1; i < codePayloadAmountSize; i++ {
		buffer[i] = byte(amountToSerialize >> uint64(8*(i-1)) & uint64(0xFF))
	}

	return buffer
}

package swt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
)

type Token struct {
	Protocol byte
	payload  IPayload
	raw      []byte
}
type Option func(*Token)

func New(payload IPayload, ops ...Option) string {
	token := &Token{Protocol: 1, payload: payload, raw: make([]byte, 0)}
	for _, op := range ops {
		op(token)
	}
	signingBytes := token.SigningBytes()
	bs := append(signingBytes, token.Sign(signingBytes, payload.SignKey())...)
	return EncodeSegment(bs)

}
func (t *Token) SigningBytes() []byte {
	bs := bytes.NewBuffer([]byte{t.Protocol})
	bs.Write(t.payload.Encode())
	return bs.Bytes()
}
func (t *Token) Sign(signingBytes, keyBytes []byte) []byte {
	hasher := hmac.New(sha256.New, keyBytes)
	hasher.Write(signingBytes)
	return hasher.Sum(nil)
}

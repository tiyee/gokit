package jwt

import (
	"encoding/json"
	"github.com/tiyee/gokit/encodes"
	"strings"
)

type Token[C IClaims] struct {
	Raw           string        // The raw token.  Populated when you Parse a token
	SigningMethod SigningMethod // The signing method used or to be used
	Header        Header        // The first segment of the token
	Claims        C             // The second segment of the token
	Signature     string        // The third segment of the token.  Populated when you Parse a token
	Valid         bool          // Is the token valid?  Populated when you Parse/Verify a token
}

type Option[C IClaims] func(token *Token[C])

func New[C IClaims](claims C, opts ...Option[C]) *Token[C] {
	token := &Token[C]{
		Claims:        claims,
		SigningMethod: nil,
		Header: Header{
			Typ: "JWT",
			Alg: HS256,
		},
	}
	for _, opt := range opts {
		opt(token)
	}
	token.SigningMethod = GetSigningMethod(token.Header.Alg)
	return token
}

// SignedString Get the complete, signed token
func (t *Token[C]) SignedString() (string, error) {
	var sig, sstr string
	var err error
	if sstr, err = t.SigningString(); err != nil {
		return "", err
	}
	if sig, err = t.SigningMethod.Sign(sstr, t.Claims.GetSigningKey()); err != nil {
		return "", err
	}
	return strings.Join([]string{sstr, sig}, "."), nil
}

// SigningString Generate the signing string.  This is the
// most expensive part of the whole deal.  Unless you
// need this for something special, just go straight for
// the SignedString.
func (t *Token[C]) SigningString() (string, error) {
	var err error
	parts := make([]string, 2)
	for i, _ := range parts {
		var jsonValue []byte
		if i == 0 {
			if jsonValue, err = json.Marshal(t.Header); err != nil {
				return "", err
			}
		} else {
			if jsonValue, err = json.Marshal(t.Claims); err != nil {
				return "", err
			}
		}

		parts[i] = EncodeSegment(jsonValue)
	}
	return strings.Join(parts, "."), nil
}

// EncodeSegment Encode JWT specific base64url encoding with padding stripped
func EncodeSegment(seg []byte) string {
	return encodes.Base64URLEncoding(seg)
	//return strings.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
}

// DecodeSegment Decode JWT specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
	//if l := len(seg) % 4; l > 0 {
	//	seg += strings.Repeat("=", 4-l)
	//}
	return encodes.Base64URLDecoding(seg)
	//return base64.URLEncoding.DecodeString(seg)
}

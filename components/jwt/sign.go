package jwt

import (
	"crypto"
	"crypto/hmac"
)

type Signer string

const HS256 Signer = "HS256"

// SigningMethod Implement SigningMethod to add new methods for signing or verifying tokens.
type SigningMethod interface {
	Verify(signingString, signature string, key []byte) error // Returns nil if signature is valid
	Sign(signingString string, key []byte) (string, error)    // Returns encoded signature or error
	Alg() Signer                                              // returns the alg identifier for this method (example: 'HS256')
}

func GetSigningMethod(alg Signer) SigningMethod {
	switch alg {
	case HS256:
		return NewSigningMethodHMAC()
	default:
		return nil
	}
}
func NewSigningMethodHMAC() SigningMethod {
	return &SigningMethodHMAC{HS256, crypto.SHA256}
}

type SigningMethodHMAC struct {
	Name Signer
	Hash crypto.Hash
}

func (m *SigningMethodHMAC) Alg() Signer {
	return m.Name
}

// Verify the signature of HSXXX tokens.  Returns nil if the signature is valid.
func (m *SigningMethodHMAC) Verify(signingString, signature string, keyBytes []byte) error {

	// Decode signature, for comparison
	sig, err := DecodeSegment(signature)
	if err != nil {
		return err
	}

	// Can we use the specified hashing method?
	if !m.Hash.Available() {
		return ErrHashUnavailable
	}

	// This signing method is symmetric, so we validate the signature
	// by reproducing the signature from the signing string and key, then
	// comparing that against the provided signature.
	hasher := hmac.New(m.Hash.New, keyBytes)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return ErrSignatureInvalid
	}

	// No validation errors.  Signature is good.
	return nil
}

// Sign Implements the Sign method from SigningMethod for this signing method.
// Key must be []byte
func (m *SigningMethodHMAC) Sign(signingString string, keyBytes []byte) (string, error) {

	if !m.Hash.Available() {
		return "", ErrHashUnavailable
	}

	hasher := hmac.New(m.Hash.New, keyBytes)
	hasher.Write([]byte(signingString))

	return EncodeSegment(hasher.Sum(nil)), nil

}

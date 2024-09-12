package jwt

import (
	"bytes"
	"encoding/json"
	"strings"
)

func Parse[T IClaims](tokenString string) (*Token[T], error) {
	token, parts, err := parseUnverified[T](tokenString)
	if err != nil {
		return token, err
	}
	vErr := &ValidationError{}

	// Perform validation
	token.Signature = parts[2]
	if err = token.SigningMethod.Verify(strings.Join(parts[0:2], "."), token.Signature, token.Claims.GetSigningKey()); err != nil {
		vErr.Inner = err
		vErr.Errors |= ValidationErrorSignatureInvalid
	}
	if !token.Claims.SkipClaimsValidation() {
		if err = token.Claims.Valid(); err != nil {
			vErr.Inner = err
			vErr.Errors |= ValidationErrorClaimsInvalid
		}
	}

	if vErr.valid() {
		token.Valid = true
		return token, nil
	}

	return token, vErr
}
func parseUnverified[T IClaims](tokenString string) (token *Token[T], parts []string, err error) {
	parts = strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, parts, NewValidationError("token contains an invalid number of segments", ValidationErrorMalformed)
	}

	token = &Token[T]{Raw: tokenString}

	// parse Header
	var headerBytes []byte
	if headerBytes, err = DecodeSegment(parts[0]); err != nil {
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			return token, parts, NewValidationError("tokenstring should not contain 'bearer '", ValidationErrorMalformed)
		}
		return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
	}
	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
		return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
	}

	// parse Claims
	var claimBytes []byte

	if claimBytes, err = DecodeSegment(parts[1]); err != nil {
		return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
	}
	dec := json.NewDecoder(bytes.NewBuffer(claimBytes))

	err = dec.Decode(&token.Claims)

	// Handle decode error
	if err != nil {
		return token, parts, &ValidationError{Inner: err, Errors: ValidationErrorMalformed}
	}

	// Lookup signature method
	if token.SigningMethod = GetSigningMethod(token.Header.Alg); token.SigningMethod == nil {
		return token, parts, NewValidationError("signing method (alg) is unavailable.", ValidationErrorUnverifiable)
	}

	return token, parts, nil
}

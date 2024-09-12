package jwt_test

import "github.com/tiyee/gokit/components/jwt"

type UserClaim struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (u UserClaim) SkipClaimsValidation() bool {
	return true
}

func (u UserClaim) GetSigningKey() []byte {
	return []byte("bejson")
}

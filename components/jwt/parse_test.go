package jwt_test

import (
	"github.com/tiyee/gokit/components/jwt"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

var hmacTestData2 = []struct {
	name        string
	tokenString string
	alg         string
	claims      UserClaim
	valid       bool
}{
	{
		name:        "web sample",
		tokenString: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3MjYxMzE1MDMsImlhdCI6MTcyNjA0NTEwMywic3ViIjoiZGVtbyIsInVzZXJfaWQiOjF9.IZ6c-a9qSnKt6spbIu2Ci_Nyd3hoG7Kk3fFLJ2gJHLA",
		alg:         "HS256",
		claims: UserClaim{

			StandardClaims: jwt.StandardClaims{
				Subject:   "demo",
				IssuedAt:  1726045103,
				ExpiresAt: 1726131503,
			},
			UserId: 1,
		},
		valid: true,
	},
	{
		name:        "web sample2",
		tokenString: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3MjYxMzE1MDMsImlhdCI6MTcyNjA0NTEwMywibmJmIjoxNzI2MDUzNDAwLCJzdWIiOiJkZW1vIiwidXNlcl9pZCI6MX0.n-TO3ZTAPOtTY9hIObZkAxU8qfIRn9CX1tV4MZTb4uE",
		alg:         "HS256",
		claims: UserClaim{
			StandardClaims: jwt.StandardClaims{
				Subject:   "demo",
				IssuedAt:  1726045103,
				ExpiresAt: 1726131503,
				NotBefore: 1726052400 + 1000,
			},
			UserId: 1,
		},
		valid: true,
	},
}

func TestParse(t *testing.T) {
	as := assert.NewAssert(t, "TestParse")

	for _, data := range hmacTestData2 {
		claims := data.claims
		tokenString, err := jwt.New(claims).SignedString()
		assert.IsNil(as, err)
		tk, err := jwt.Parse[UserClaim](tokenString)
		if err != nil {
			t.Errorf("[%v] Error signing token: %v", "testParse", err)
		}

		assert.Equal(as, claims.IssuedAt, tk.Claims.IssuedAt)
		assert.Equal(as, claims.Subject, tk.Claims.Subject)
		assert.Equal(as, claims.ExpiresAt, tk.Claims.ExpiresAt)
		assert.Equal(as, claims.UserId, tk.Claims.UserId)
	}

}

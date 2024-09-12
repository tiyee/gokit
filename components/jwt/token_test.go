package jwt_test

import (
	"encoding/json"
	"fmt"
	"github.com/tiyee/gokit/components/jwt"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

var hmacTestData = []struct {
	name        string
	tokenString string
	alg         string
	claims      UserClaim
	valid       bool
}{
	{
		name:        "web sample",
		tokenString: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3MjYxMzE1MDMsImp0aSI6IndlYiIsImlhdCI6MTcyNjA0NTEwMywic3ViIjoiZGVtbyIsInVzZXJfaWQiOjF9.43cM6-RbQCJZo99iCNRLJ6wcO-Pg77aIT6QVXzlf8Vo",
		alg:         "HS256",
		claims: UserClaim{

			StandardClaims: jwt.StandardClaims{
				Subject:   "demo",
				IssuedAt:  1726045103,
				ExpiresAt: 1726131503,
				Id:        "web",
			},
			UserId: 1,
		},
		valid: true,
	},
	{
		name:        "web sample2",
		tokenString: "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3MjYxMzE1MDMsImlhdCI6MTcyNjA0NTEwMywibmJmIjoxNzI2MDUyNDAwLCJzdWIiOiJkZW1vIiwidXNlcl9pZCI6MX0.ufXkrk9zIE2CydVbMYnd2MS0o4Fi8BbvAQu2Vs50wtM",
		alg:         "HS256",
		claims: UserClaim{
			StandardClaims: jwt.StandardClaims{
				Subject:   "demo",
				IssuedAt:  1726045103,
				ExpiresAt: 1726131503,
				NotBefore: 1726052400,
			},
			UserId: 1,
		},
		valid: true,
	},
}

func TestHMACSign(t *testing.T) {
	as := assert.NewAssert(t, "TestNew")
	for _, data := range hmacTestData {

		bs, err := json.Marshal(data.claims)
		assert.IsNil(as, err)
		fmt.Println(string(bs))
		fmt.Println()
		tk := jwt.New(data.claims)
		tokenString, err := tk.SignedString()
		assert.IsNil(as, err)
		assert.Equal(as, data.tokenString, tokenString)
	}
}

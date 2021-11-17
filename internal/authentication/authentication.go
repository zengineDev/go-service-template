package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtCustomClaims struct {
	Audience  []string `json:"aud,omitempty"`
	Scope     string   `json:"scope,omitempty"`
	ExpiresAt int64    `json:"exp,omitempty"`
	Id        string   `json:"jti,omitempty"`
	IssuedAt  int64    `json:"iat,omitempty"`
	Issuer    string   `json:"iss,omitempty"`
	NotBefore int64    `json:"nbf,omitempty"`
	Subject   string   `json:"sub,omitempty"`
}

func verifyExp(exp *time.Time, now time.Time, required bool) bool {
	if exp == nil {
		return !required
	}
	return now.Before(*exp) || now.Equal(*exp)
}

func (c *JwtCustomClaims) VerifyExpiresAt(cmp int64, req bool) bool {
	if c.ExpiresAt == 0 {
		return verifyExp(nil, time.Unix(cmp, 0), req)
	}

	t := time.Unix(c.ExpiresAt, 0)
	return verifyExp(&t, time.Unix(cmp, 0), req)
}

func (c JwtCustomClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := jwt.TimeFunc().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if !c.VerifyExpiresAt(now, false) {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

func ForContext(ctx context.Context) *JwtCustomClaims {
	raw, _ := ctx.Value("gql.user").(*JwtCustomClaims)
	return raw
}

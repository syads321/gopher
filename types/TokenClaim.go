package types

import (
	"github.com/dgrijalva/jwt-go"
)

// TokenClaim dfadf
type TokenClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

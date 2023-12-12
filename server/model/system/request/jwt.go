package request

import (
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

// BaseClaims base data jwt
type BaseClaims struct {
	ID          uint
	Username    string
	Role       string
}

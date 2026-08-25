package jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewBasicOptions(userid uint) jwt.MapClaims {
	return jwt.MapClaims{
		"sub": userid,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
}

func NewJWTMethod(method string) (jwt.SigningMethod, error) {
	switch strings.ToUpper(strings.TrimSpace(method)) {
	// HMAC
	case "HS256", "HMAC256":
		return jwt.SigningMethodHS256, nil
	case "HS384", "HMAC384":
		return jwt.SigningMethodHS384, nil
	case "HS512", "HMAC512":
		return jwt.SigningMethodHS512, nil

	// RSA PKCS#1 v1.5
	case "RS256":
		return jwt.SigningMethodRS256, nil
	case "RS384":
		return jwt.SigningMethodRS384, nil
	case "RS512":
		return jwt.SigningMethodRS512, nil

	// RSA-PSS
	case "PS256":
		return jwt.SigningMethodPS256, nil
	case "PS384":
		return jwt.SigningMethodPS384, nil
	case "PS512":
		return jwt.SigningMethodPS512, nil

	// ECDSA
	case "ES256":
		return jwt.SigningMethodES256, nil
	case "ES384":
		return jwt.SigningMethodES384, nil
	case "ES512":
		return jwt.SigningMethodES512, nil

	// EdDSA / Ed25519
	case "EDDSA", "ED25519":
		return jwt.SigningMethodEdDSA, nil

	default:
		return nil, fmt.Errorf("unsupported JWT signing method: %q", method)
	}
}

func GenerateToken(method jwt.SigningMethod) (string, error) {
	token := jwt.New(method)
	tokenString, err := token.SignedString(os.Getenv("SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenWithOptions(method jwt.SigningMethod, options ...map[string]any) (string, error) {
	claims := jwt.MapClaims{}

	for _, option := range options {
		for key, value := range option {
			claims[key] = value
		}
	}

	token := jwt.NewWithClaims(method, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

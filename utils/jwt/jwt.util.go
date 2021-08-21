package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type TokenPayload struct {
	ID uint
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION"))
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})
	token, err := t.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		panic(err)
	}
	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)
	if err != nil {
		return nil, err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("Something went wrong")
	}
	return &TokenPayload{ID: uint(id)}, nil
}

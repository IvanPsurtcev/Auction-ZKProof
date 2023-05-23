package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type (
	hmacProvider struct {
		hmacSecret []byte
		issuer     string
		duration   time.Duration
	}

	HmacProvider interface {
		CreateStandard(subject string) (string, error)
		Verify(tokenString string) (*jwt.RegisteredClaims, error)
	}
)

func NewHmacProvider(hmacSecret string, issuer string, duration time.Duration) HmacProvider {
	return &hmacProvider{
		hmacSecret: []byte(hmacSecret),
		issuer:     issuer,
		duration:   duration,
	}
}

func (j *hmacProvider) CreateStandard(subject string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    j.issuer,
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(j.duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.hmacSecret)
}

func (j *hmacProvider) Verify(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.hmacSecret, nil
	})
	if err != nil {
		return nil, ErrAuthError
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrAuthError
}

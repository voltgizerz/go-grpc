package services

import (
	"errors"
	"os"
	"time"

	"github.com/auth-service/models"
	"github.com/golang-jwt/jwt"
)

// JwtWrapper - .
type JwtWrapper struct {
	SecretKey string
	Issuer    string
}

// NewJwtWrapper - create new jwt wrapper.
func NewJwtWrapper() JwtWrapper {
	return JwtWrapper{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
		Issuer:    "felix",
	}
}

// JwtClaims - .
type JwtClaims struct {
	jwt.StandardClaims
	ID        int64
	Username  string
	ExpiresAt int64
}

// GenerateToken - generate new jwt token.
func (w *JwtWrapper) GenerateToken(user models.User) (signedToken string, err error) {
	expiredAt := time.Now().Local().Add(time.Hour * 1).Unix()
	claims := &JwtClaims{
		ID:        user.ID,
		Username:  user.Username,
		ExpiresAt: expiredAt,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken - validate jwt token.
func (w *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.StandardClaims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("jwt is expired")
	}

	return claims, nil
}

package jwtext

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

func NewClaims(id string, email, username string, duration time.Duration) *Claims {
	return &Claims{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			ExpiresAt: time.Now().Add(duration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

func LoadClaims(key, token string) (*Claims, error) {
	claims := new(Claims)
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "can not parse jwt token \"%v\"", token)
	} else if !jwtToken.Valid {
		return nil, errors.Wrapf(err, "invalid jwt token \"%v\"", token)
	}
	return claims, nil
}

type Claims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (c Claims) Token(key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(key))
}

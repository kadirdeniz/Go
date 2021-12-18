package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Create(id uint) (string, error) {
	atClaim := jwt.MapClaims{
		"ID":  id,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	return at.SignedString([]byte("secret key"))
}

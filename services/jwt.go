package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const key = "namdzvcl"

func CreateToken(typeid, atype string) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["type"] = atype
	atClaims["type_id"] = typeid
	atClaims["exp"] = time.Now().Add(time.Minute * 180).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

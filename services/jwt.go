package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const key = "namdzvcl"

func CreateToken(userId uint64) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if token.Valid {
		var payload map[string]interface{}
		bytes, e := json.Marshal(token.Claims)
		if e != nil {
			return nil, errors.New("Token false")
		}
		e = json.Unmarshal(bytes, &payload)
		if e != nil {
			return nil, errors.New("Token false")
		}
		return payload, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("Not a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, errors.New("Token expired")
		} else {
			return nil, errors.New("Token is not valid")
		}
	} else {
		return nil, errors.New("Token false")
	}
}

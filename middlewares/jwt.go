package middlewares

import (
	"encoding/json"
	"errors"
	"go-orderfood/responses"

	"net/http"

	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
)

const key = "namdzvcl"

var Jwt = func(ctx *context.Context) {
	auth := ctx.Request.Header["Auth"]
	if len(auth) < 1 {
		ctx.Output.JSON(responses.UnAuthResponse, true, true)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}
	payload, err := ParseToken(auth[0])

	if err != nil {
		ctx.Output.JSON(responses.UnAuthResponse, true, true)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return
	}
	// log.Println(payload["type_id"].(string))
	// log.Println(payload["type"].(string))
	ctx.Request.Header.Set("id", payload["type_id"].(string))
	ctx.Request.Header.Set("type", payload["type"].(string))
	// log.Println(ctx.Request.Header)
}

func ParseToken(tokenString string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if token == nil {
		return nil, errors.New("Token nil")
	}
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

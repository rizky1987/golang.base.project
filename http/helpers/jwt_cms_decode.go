package helpers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"hotel/commonHelpers"
)

type JWTSession struct {
	ID       mssql.UniqueIdentifier `json:"ID"`
	Username string                 `json:"Username"`
}

func (u *HTTPHelper) ValidateCMSJWTData(c echo.Context) (JWTSession, error) {

	dataJwt := JWTSession{}
	jwtString := c.Request().Header.Get("Authorization")
	authorizationSplits := commonHelpers.StringSplitToArrayString(jwtString, "Bearer")

	if len(authorizationSplits) < 2 {

		return dataJwt, errors.New("Please input your token")
	}

	jwtToken := commonHelpers.TrimWhiteSpace(authorizationSplits[1])
	if jwtToken == "" {

		return dataJwt, errors.New("Please input your token")
	}

	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmaKey), nil
	})

	if token == nil {

		return dataJwt, errors.New("your token is null")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims != nil {

		dataSession := claims["rl"].(string)
		key := []byte(keyConfig)

		decoded, err := hex.DecodeString(dataSession)
		if err != nil {
			return dataJwt, errors.New("there is something when decode data")
		}

		plaintext, err := Decrypt(decoded, key)
		if err != nil {
			return dataJwt, errors.New("there is something when Decrypt data")
		}

		err = json.Unmarshal(plaintext, &dataJwt)
		if err != nil {
			return dataJwt, errors.New("there is something when Unmarshal data")
		}
	}

	if dataJwt.Username == "" {
		return dataJwt, errors.New("username data from jwt is empty")
	}

	return dataJwt, nil
}

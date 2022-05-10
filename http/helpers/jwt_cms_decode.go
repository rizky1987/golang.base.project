package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"example/commonHelpers"
	config "example/config/env"
)

type JWTSession struct {
	ID         string  `json:"rl-id"`
	Name       string  `json:"rl-fullname"`
	RlRole     string  `json:"rl-role"`
	BranchID   string  `json:"rl-branch-id"`
	Email      string  `json:"rl-email"`
	CustomerID string  `json:"rl-customer-id"`
	Role       string  `json:"role"`
	Nbf        float64 `json:"nbf"`
	Exp        float64 `json:"exp"`
	Iat        float64 `json:"iat"`
}

func (u *HTTPHelper) ValidateCMSJWTData(c echo.Context) (JWTSession, error) {

	config := config.NewViperConfig()

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

		return []byte(config.GetString("security.jwt_secret")), nil
	})

	if token == nil {
		return dataJwt, errors.New("your token is null")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid && claims != nil {
		jsonbody, err := json.Marshal(claims)
		if err != nil {
			return dataJwt, err
		}

		if err := json.Unmarshal(jsonbody, &dataJwt); err != nil {
			return dataJwt, errors.New("you need to update JWTSession Struct")
		}
	} else {
		return dataJwt, errors.New("your token is invalid")
	}

	today := time.Now().UTC()
	jtwExpireDateTime := commonHelpers.Float64ToDateTimeUTC(dataJwt.Exp)
	if today.After(jtwExpireDateTime) {
		return dataJwt, errors.New("your token has been expired")
	}

	if commonHelpers.TrimWhiteSpace(dataJwt.Name) == "" || commonHelpers.TrimWhiteSpace(dataJwt.ID) == "" {
		return dataJwt, errors.New("Invalid JWT session data. Name and ID should not be null")
	}

	return dataJwt, nil
}

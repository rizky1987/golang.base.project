package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	SQLEntity "hotel/databases/entities/sql"

	"github.com/dgrijalva/jwt-go"
)

var (
	hmaKey      string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjAsImlhdCI6MTU3NzE3MTUxNiwianRpIjoiTHJEUFlQRmwzOERQRTI3MUtUOUVmUT09IiwibmJmIjoxNTc3MTcxNTE2LCJweGwiOiIxOTE2ODY3YjY2OGUyNTIwN2NjZDQ2ZDNiY2Y5OGUwOGZmNmE3MDExMzljNGMxMjZkYmJhNTE3NDA0MzMyNDM5Nzg3ZWRlYmI2M2M1Njk2YmU0MmEyZTY4ZmQzNDc5MDE2MmQ2ZTMxYjQwNWI4NDQ5Y2IxM2NmYmVhMDczNjhiN2JmNzRjMTVjZjc2YTljNWU3Zjc0NDJkOWEyNTMxMGI5YmEzNGQzODA2YzIwZjllZWQzOWM2MzllZDk1OTE1MzVhMTZkMDZkMGZlZWQ2NDIwZWZmMDg4ZDZlZjM5MjUwYmU0NzZiYjdjMmFhNzJhNDFmYTgxNmYzNDg5ODE1ZTZjNWQzZTA0NWJkZGE4NTMwMDNkOWZhZDU2ZTU3YjAyYzkzNzY0MmIxODUwYTY4MzJkIiwic3ViIjoiSDJXTVhOTC02UjRYV0FNLUJRUUo3UVctRTdMQ1c3WSJ9.aZ-tBaIbmEeXEGs2YtXssZ6PvVwnzkM0pwc3_Am3m9k" //JWT Secret
	keyConfig   string = "kdnskrydkaqpld9352md93js53kd01m2"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              /// internal encript data
	UserSession *SQLEntity.User
)

func GenetarateToken(user *SQLEntity.User) (string, error) {

	hexid := fmt.Sprintf("%x", user.ID.String())

	userSession := SQLEntity.User{
		ID:       user.ID,
		Username: user.Username,
	}

	jsonString, err := json.Marshal(userSession) // Set Data
	if err != nil {
		fmt.Println("Error encoding JSON")
		return "", errors.New("Error encoding JSON")
	}

	key := []byte(keyConfig)

	encrypt, err := Encrypt(jsonString, key)
	if err != nil {
		return "", err
	}

	rl := hex.EncodeToString(encrypt)

	currentTimestamp := time.Now().UTC().Unix()
	var ttl int64 = (3600 * 10) // expired time in second
	// md5 of sub & iat
	h := md5.New()
	io.WriteString(h, hexid)
	io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": currentTimestamp,
		"exp": currentTimestamp + ttl, // 0 <<<< "0" for unlimited expired
		"nbf": currentTimestamp,
		"jti": h.Sum(nil),
		"rl":  rl,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmaKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

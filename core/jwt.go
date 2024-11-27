package core

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type StandardClaims struct {
	jwt.MapClaims
	Id string
}
type JwtLib struct {
	expTimeAccessTokenInMin  uint16
	expTimeRefreshTokenInMin uint16
	secret                   []byte
	accessSecret             []byte
}

func NewJwtLib() JwtLib {
	jwt := JwtLib{}
	i, err := strconv.ParseInt(os.Getenv("JWT_ACCESS_EXP_IN_MIN"), 10, 16)
	if err != nil {
		fmt.Println("invalid JWT_ACCESS_EXP_IN_MIN")
	}
	jwt.expTimeAccessTokenInMin = uint16(i)
	i, err = strconv.ParseInt(os.Getenv("JWT_REFRESH_EXP_IN_MIN"), 10, 16)
	if err != nil {
		fmt.Println("invalid JWT_REFRESH_EXP_IN_MIN")
	}
	jwt.expTimeRefreshTokenInMin = uint16(i)
	token, tokenErr := os.ReadFile("keys/private.pem")
	if tokenErr != nil {
		fmt.Println("invalid private file")
	}
	jwt.secret = token
	token, tokenErr = os.ReadFile("keys/public.pem")
	if tokenErr != nil {
		fmt.Println("invalid private file")
	}
	jwt.accessSecret = token
	return jwt
}

func (jwtLib JwtLib) SignJwt(claims StandardClaims) (string, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id": claims.Id,
		// "nbf": time.Now().Add(time.Duration(jwtLib.expTimeAccessTokenInMin)),
	})
	accessTokenString, err := token.SignedString(jwtLib.accessSecret)
	if err != nil {
		fmt.Println(err.Error())
	}
	reToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  claims.Id,
		"nbf": time.Now().Add(time.Duration(jwtLib.expTimeRefreshTokenInMin)),
	})
	refreshTokenString, refreshErr := reToken.SignedString(jwtLib.secret)
	if refreshErr != nil {
		fmt.Println(err.Error())
	}
	return accessTokenString, refreshTokenString
}

func (jwtLib JwtLib) ValidateJwt(tokenString string) (StandardClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("unexpected signing method")
			return nil, errors.New("Invalid Token")
		}
		return []byte(jwtLib.accessSecret), nil
	})
	if err != nil {
		fmt.Println(err.Error(), 'a')
		return StandardClaims{}, errors.New("Invalid Token")
	}
	fmt.Println(token.Claims)
	claims, ok := token.Claims.(StandardClaims)
	fmt.Println(ok, "c")
	if ok {
		return claims, nil
	}
	fmt.Println(err, "b")
	return StandardClaims{}, errors.New("Invalid Token")
}

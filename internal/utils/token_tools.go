package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenUtil interface {
	generateToken(id string, email string, nickname string) (string, error)

	parseToken(token string)
}

type UserInfo struct {
	ID       string
	Email    string
	Nickname string
}

// 构造palyload的数据格式
type MyClaim struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

// secret签名
var mySignatureSecret []byte = []byte("!@#qwe")

func GenerateToken(id string, email string, nickname string) (string, error) {
	myclaim := MyClaim{
		ID:       id,
		Email:    email,
		Nickname: nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	//生成token
	signature, err := jwt.NewWithClaims(jwt.SigningMethodHS256, myclaim).SignedString(mySignatureSecret)
	if err != nil {
		return "", fmt.Errorf("token生成错误，%v", err)
	}

	return signature, nil
}

func ParseToken(token string) (*UserInfo, error) {
	if token == "" {
		return nil, fmt.Errorf("authorization token required")
	}

	if len(token) > 7 && strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}
	result, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return mySignatureSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return &UserInfo{
		ID:       result.Claims.(*MyClaim).ID,
		Email:    result.Claims.(*MyClaim).Email,
		Nickname: result.Claims.(*MyClaim).Nickname,
	}, nil
}

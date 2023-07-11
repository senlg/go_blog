package jwtauth

import (
	"errors"
	"go_blog/global"
	"go_blog/models"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID     uint
	Username   string
	Role       models.Role
	GrantScope string
	jwt.RegisteredClaims
}

type Jwt struct {
}

// 随机字符串
var Letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 加载盐ID
func (m *MyCustomClaims) RandStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(rand_bytes)
}

func (j *Jwt) CreateToken(claim MyCustomClaims) (string, error) {
	new_claim := MyCustomClaims{
		UserID:     claim.UserID,
		Username:   claim.Username,
		GrantScope: "read_user_info", // 作用范围
		Role:       claim.Role,       // 用户身份
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                 // 签发者
			Subject:   claim.Username,                                // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},    // 签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), // 过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                // 最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                // 签发时间
			ID:        claim.RandStr(10),                             // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, new_claim).SignedString([]byte(global.Config.System.Secret))
	return token, err
}

func (j *Jwt) ParseToken(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.System.Secret), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim 无效")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

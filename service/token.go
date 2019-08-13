package service

import (
	"code.qlteam.com/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/willerhe/webbase/configer"
	"time"
)

type token int

var Token token

// todo TOKEN2 生成token
func (token) General(user model.User) string {
	sign := []byte(configer.Config.Get("token.sign").(string))
	claims := new(model.Token)
	// 过期时间 1小时
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.UID = user.ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(sign)
	return ss
}

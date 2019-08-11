package service

import (
	"code.qlteam.com/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/willerhe/webbase/configer"
)

type token int

var Token token

func (token) General(user model.User) string {
	sign := []byte(configer.Config.Get("token.sign").(string))
	claims := &jwt.StandardClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(sign)
	return ss
}

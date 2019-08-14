package service

import (
	"code.qlteam.com/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/willerhe/webbase/configer"
	"log"
	"time"
)

type token int

var Token token

// todo TOKEN2 生成token
func (token) General(user *model.User) string {
	sign := []byte(configer.Config.Get("token.sign").(string))
	claims := new(model.Claims)
	// 过期时间 1小时
	claims.ExpiresAt = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.User = *user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(sign)
	return ss
}

func (token) Parse(token string) *model.Claims {
	claims := &model.Claims{}
	t, err := jwt.ParseWithClaims(token, claims, func(i2 *jwt.Token) (i interface{}, e error) {
		return []byte(configer.Config.Get("token.sign").(string)), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	cl, ok := t.Claims.(*model.Claims)
	if ok && t.Valid {
		fmt.Printf("%v %v", cl.Id, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	return cl

}

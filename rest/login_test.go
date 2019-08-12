package rest

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/willerhe/webbase/configer"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	sign := []byte(configer.Config.Get("token.sign").(string))
	claims := &jwt.StandardClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, _ := token.SignedString(sign)
	fmt.Printf("%v\n", ss)

	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(ss, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configer.Config.Get("token.sign").(string)), nil
	})

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		fmt.Printf("%v", claims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

}

// Override time value for tests.  Restore default value after.
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

func GenToken(data map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//添加令牌关键信息
	Tokenexp, _ := strconv.Atoi(beego.AppConfig.String("token::exp"))
	//添加令牌期限
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(Tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["data"] = data
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(beego.AppConfig.String("token::secret")))
	return tokenString
}

func CheckToken(tokenString string) interface{} {
	var data interface{}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(beego.AppConfig.String("token::secret")), nil
	})

	if token != nil && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		data = claims["data"]
	}
	return data
}

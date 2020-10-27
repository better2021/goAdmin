package common

import (
	"github.com/dgrijalva/jwt-go"
	"goAdmin/model"
	"time"
)

var jwtKey = []byte("a_secret_co")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user model.User)(string,error){
	expirationTime := time.Now().Add(7*24*time.Hour) // 设置token有效时间为7天
	claims := &Claims{
		UserId: user.ID,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer:   "ginGO",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err := token.SignedString(jwtKey)
	if err != nil{
		return "", err
	}
	return tokenString,nil
}

func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims := &Claims{}
	
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey,nil
	})

	return token,claims,err
}
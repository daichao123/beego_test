package tools

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

//go:embed token_jwt_key.pem
var JwtPublicKey string

const (
	DefaultExpireSeconds  int64 = 24 * 60 * 60
	DefaultRefreshSeconds int64 = 1 * 60 * 60 //设置刷新时间为过期前 1 小时

)

func CreateToken(userName string, password string, uid int, app string) (tokenString string, expSeconds int64, expTime int64, refreshTime int64, err error) {
	// token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	expTime = time.Now().Add(time.Second * time.Duration(DefaultExpireSeconds)).Unix()
	refreshTime = expTime - DefaultRefreshSeconds
	//添加令牌期限
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()
	claims["username"] = userName
	claims["uid"] = uid
	claims["password"] = password
	claims["app"] = app
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(JwtPublicKey))
	if err != nil {
		logs.Error("generate json web token failed !! error :", err)
		return tokenString, 0, 0, 0, err
	}
	return tokenString, DefaultExpireSeconds, expTime, refreshTime, err
}

//CheckToken 检查token
func CheckToken(tokenString string) (userName string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(JwtPublicKey), nil
	})
	if err != nil {
		logs.Error(err)
		return "", err
	}
	// token.Valid里已经包含了过期判断
	if token != nil && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		userName = claims["userName"].(string)
	}

	return userName, err
}

// RefreshToken update expireAt and returnDoc a new token
//	只能在过期之前刷新才可以
func RefreshToken(tokenString string) (string, int64, int64, int64, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtPublicKey), nil
		})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", 0, 0, 0, err
	}

	expTime := time.Now().Add(time.Second * time.Duration(DefaultExpireSeconds)).Unix()
	refreshTime := expTime - DefaultRefreshSeconds
	//添加令牌期限
	newClaims := jwt.MapClaims{
		"exp":      expTime,
		"iat":      time.Now().Unix(),
		"username": claims["username"].(string),
		"password": claims["password"].(string),
		"uid":      claims["uid"],
		"app":      claims["app"],
	}

	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString([]byte(JwtPublicKey))
	if err != nil {
		logs.Error("generate new fresh json web token failed !! error :", err)
		return "", 0, 0, 0, err
	}
	return tokenStr, DefaultExpireSeconds, expTime, refreshTime, err
}

// ParseToken 解析 JWT token 从http header 中
func ParseToken(authString string) (t *jwt.Token, err error) {
	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		logs.Error("AuthString invalid:", authString)
		return nil, err
	}
	tokenString := kv[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtPublicKey), nil
	})
	if err != nil {
		logs.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, err
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	if !token.Valid {
		return token, errors.New(fmt.Sprintf("Token invalid:%s", tokenString))
	}
	return token, nil

}

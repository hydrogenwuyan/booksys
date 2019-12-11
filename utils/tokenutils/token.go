package tokenutils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/siddontang/go/log"
	"project/booksys/common"
	"strconv"
	"time"
)

var (
	SECRET_KEY = []byte("book management system")
)

const (
	AccessTokenExpiredSecs = 3600 // 过期时间
)

const (
	TokenOk        = iota //正常
	TokenExpired          //过期
	TokenInvalid          //无效
	TokenParseFail        //token解析失败
)

type Claims struct {
	UserId int64 `json:"UserId"`
	jwt.StandardClaims
}

// 获取token
func GenerateToken(id int64) (token string, err error) {
	claims := Claims{
		id,
		jwt.StandardClaims{
			Id:        fmt.Sprint(id),
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + AccessTokenExpiredSecs),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(SECRET_KEY)
	if err != nil {
		log.Error("generate access_token failed : %v", err)
		return
	}
	return
}

// 设置token
func SetToken(id int64, token string) (err error) {
	err = common.RedisClient.Set(TokenKey(fmt.Sprint(id)), token, time.Second*time.Duration(AccessTokenExpiredSecs)).Err()
	if err != nil {
		log.Error("redis set token fail, error: ", err)
		return
	}

	log.Debug("token[%s] : %s", TokenKey(fmt.Sprint(id)), token)

	return
}

// 检查token
func CheckToken(token string) (result uint8, id int64) {
	t, err := ParseToken(token)
	if err != nil {
		result = TokenParseFail
		return
	}

	var ok bool
	var claims *jwt.StandardClaims
	if claims, ok = t.Claims.(*jwt.StandardClaims); ok && t.Valid {
		if claims.Id == "" {
			log.Error("req token.id == 0, there must be something wrong with token.")
			result = TokenParseFail
			return
		}
		//check if token equals to cached token
		result = CheckTokenSignature(claims.Id)
		if result != TokenOk {
			log.Error("CheckTokenSignature fail uid:%v, result:%v", claims.Id, result)
		}

		idInt, err := strconv.Atoi(claims.Id)
		if err != nil {
			log.Error("CheckToken Id Atoi Fail, error: ", err)
			return
		}
		id = int64(idInt)
		return
	} else {
		result = TokenParseFail
		log.Error("parse token failed : %v", err)
		return
	}

}

func CheckTokenSignature(id string) uint8 {
	//get cached token.
	tokenKey := TokenKey(id)
	err := common.RedisClient.Get(tokenKey).Err()
	if err != nil {
		//log.Error("redis get tokenkey fail, error: %v", err)
		return TokenExpired
	}

	return TokenOk
}

func ParseToken(tokenStr string) (t *jwt.Token, err error) {
	claims := &jwt.StandardClaims{}
	t, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		log.Error("parse token [%s] failed ： %v", tokenStr, err)
		return
	}

	return
}

func TokenKey(id string) string {
	return fmt.Sprintf("booksys.%s", id)
}

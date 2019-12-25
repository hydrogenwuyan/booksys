package tokenutils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"project/booksys/common"
	"time"
)

var (
	SECRET_KEY = []byte("book management system")
)

const (
	AccessTokenExpiredSecs = 3000 // 过期时间
)

const (
	TokenOk        = iota //正常
	TokenExpired          //过期
	TokenInvalid          //无效
	TokenParseFail        //token解析失败
)

type Claims struct {
	UserId string `json:"UserId"`
	jwt.StandardClaims
}

// 获取token
func GenerateToken(id string) (token string, err error) {
	claims := Claims{
		id,
		jwt.StandardClaims{
			Id:        id,
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + AccessTokenExpiredSecs),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString(SECRET_KEY)
	if err != nil {
		common.LogFuncError("generate access_token failed : %v", err)
		return
	}
	return
}

// 设置token
func SetToken(id string, token string) (err error) {
	err = common.RedisClient.Set(tokenKey(id), token, time.Second*time.Duration(AccessTokenExpiredSecs)).Err()
	if err != nil {
		common.LogFuncError("redis set token fail, error: %v", err)
		return
	}

	common.LogFuncDebug("token[%s] : %s", tokenKey(id), token)
	return
}

// 检查解析token
func CheckAndParseToken(token string) (result uint8, id string) {
	return checkAndParseToken(token)
}

// 检查解析token
func checkAndParseToken(token string) (result uint8, id string) {
	t, err := parseToken(token)
	if err != nil {
		result = TokenParseFail
		return
	}

	var ok bool
	var claims *jwt.StandardClaims
	if claims, ok = t.Claims.(*jwt.StandardClaims); ok && t.Valid {
		if claims.Id == "" {
			common.LogFuncError("req token.id == 0, there must be something wrong with token.")
			result = TokenParseFail
			return
		}

		result = checkTokenExpired(claims.Id)
		if result != TokenOk {
			common.LogFuncError("checkToken Expired, uid:%v, result:%v", claims.Id, result)
			return
		}

		id = claims.Id
		return
	} else {
		result = TokenParseFail
		common.LogFuncError("parse token failed : %v", err)
		return
	}

}

// 清理token
func ClearToken(token string) bool {
	result, id := CheckAndParseToken(token)
	if result != TokenOk {
		return false
	}

	key := tokenKey(id)
	err := common.RedisClient.Del(key).Err()
	if err != nil {
		return false
	}

	return true
}

// 检查token是否过去
func checkTokenExpired(id string) uint8 {
	//get cached token.
	tokenKey := tokenKey(id)
	err := common.RedisClient.Get(tokenKey).Err()
	if err != nil {
		//log.Error("redis get tokenkey fail, error: %v", err)
		return TokenExpired
	}

	return TokenOk
}

func parseToken(tokenStr string) (t *jwt.Token, err error) {
	claims := &jwt.StandardClaims{}
	t, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	if err != nil {
		common.LogFuncError("parse token [%s] failed ： %v", tokenStr, err)
		return
	}

	return
}

func tokenKey(id string) string {
	return fmt.Sprintf("booksys.%s", id)
}

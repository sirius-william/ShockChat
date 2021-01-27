// @Title  Token.go
// @Description	Token生成器
// @Author  SiriusWilliam  2021-01-25 12:44
// @Update  2021-01-25 12:44
package utils

import (
	"ShockChatServer/protos"
	"ShockChatServer/utils/redis"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

type tokenClaim struct {
	jwt.StandardClaims
	User *protos.UserLogin
}

type TokenConfig struct {
	SecretKey string `json:"token.key"`
	Minute    int64  `json:"token.minute"`
}

var TokenConf = TokenConfig{"secretKey", 30}

/*
 * 生成Token
 * secretKey：密钥，自定义在配置文件中，服务端运行时不可修改，也不会重新读取
 * user：根据用户登录信息来生成
 * 说明：客户端发送心跳包，来获取token
 */
func CreateToken(user *protos.UserLogin) (string, error) {
	var res string
	// 将用户登录信息（未RSA解密）写入claim内来生成Token
	userTokenClaim := &tokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(TokenConf.Minute) * time.Minute).Unix(),
			Issuer:    "SiriusWilliam",
		},
		User: user,
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userTokenClaim)
	// 加入密钥
	key := TokenConf.SecretKey
	res, err := token.SignedString([]byte(key))
	if err != nil {
		fmt.Println("token error:", err)
		return "", err
	}
	// 写入redis
	_, err = redis.Exec("hset", "token", strconv.Itoa(int(user.Id)), res)
	if err != nil {
		return "", err
	}
	return res, nil
}

func TokenCheck(userid string, token string) bool {
	redisFind, err := redis.Exec("hget", "token", userid)
	if len(redisFind) == 1 && err != nil {
		if redisFind[0] == token {
			return true
		}
	}
	return false
}

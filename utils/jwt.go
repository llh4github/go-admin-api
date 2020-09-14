package utils

import (
	"math"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/llh4github/go-admin-api/global"
	"github.com/sirupsen/logrus"
)

const (
	// Iss jwt签发者
	Iss string = "iss"
	// Sub  jwt所面向的用户
	Sub = "sub"
	// Exp jwt的过期时间，这个过期时间必须要大于签发时间
	Exp = "exp"
	// Iat  jwt的签发时间
	Iat = "iat"
	// Roles 用户所属角色名列表
	Roles = "roles"
)

var jwtConf global.JwtConfig

// InitJwtConf 初始化jwt工具配置信息
// for wire tool
func InitJwtConf() global.JwtConfig {
	jwtConf = global.Conf.JwtConfig
	return jwtConf
}

// CreateTokenWithRoles 创建token
// 	uid 用户的id
// 	roles 此用户所属角色名列表
func CreateTokenWithRoles(uid string, roles []string) (string, error) {

	claims := jwt.MapClaims{}
	claims[Sub] = uid
	claims[Iat] = time.Now().Unix()
	claims[Roles] = roles
	claims[Iss] = jwtConf.Iss
	// jwt里的时间请使用时间戳，可以减少jwt长度和时间转换步骤
	claims[Exp] = time.Now().Add(
		time.Minute * time.Duration(jwtConf.Exp)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(jwtConf.Secret))
	if err != nil {
		logrus.Error("创建token出错！", err)
		return "", err
	}
	return token, nil
}

// CreateToken 创建token
// uid参数通常是用户的id
func CreateToken(uid string) (string, error) {
	return CreateTokenWithRoles(uid, make([]string, 0))
}

// ParseToken 解析token
//
// 会进行过期验证
//
// 	claims : token内的信息。实际上是一个map。
// 	notExp : 如果token未过期且有内容则为true。
func ParseToken(token string) (claims jwt.MapClaims, notExp bool) {
	// 此处会进行过期验证
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.Secret), nil
	})
	if err != nil {
		logrus.Error("解析token出错！", err)
		claims = make(jwt.MapClaims)
		notExp = false
	} else {
		claims = tk.Claims.(jwt.MapClaims)
		l := len(claims)
		if l > 0 {
			notExp = true
		} else {
			notExp = false
		}
	}
	return
}

// GetRoleNames 获取jwt中的角色名列表
func GetRoleNames(token string) (roles []string, length int) {
	m, has := ParseToken(token)
	if has {
		temp := m[Roles].([]interface{}) // 这里不能直接转换为[]string类型
		var l []string
		for _, r := range temp {
			l = append(l, r.(string))
		}
		roles = l
		logrus.Debug("GetRoleNames : ", roles)
		length = len(roles)
	}
	return
}

// GetSub 获取jwt所面向的用户
// 	如果没有获取到 Sub 字段的信息则返回 false
func GetSub(token string) (string, bool) {
	m, has := ParseToken(token)
	if has {
		sub := m[Sub].(string)
		if sub == "" {
			return "", false
		}
		return sub, true
	}
	return "", false
}

// GetExp 获取jwt的过期时间
//
// 返回时间戳（秒）
// 	token 解析不正确，会返回0
func GetExp(token string) int64 {
	m, has := ParseToken(token)
	if !has {
		// 如果token解析不正确，直接返回0
		return 0
	}
	// 解析回来就成float64  -_-|||
	tmp := m[Exp].(float64)
	return int64(tmp * math.Pow10(0))
}

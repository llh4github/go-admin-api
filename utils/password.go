package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 密码加密
// 返回加密后的字符串及可能出现的错误
func HashPassword(raw string) (string, error) {
	originPwd := []byte(raw)
	hashPwd, err := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost)
	return string(hashPwd), err
}

// MatchPassword 测试原始密码与加密后的密码是否匹配
func MatchPassword(raw, hashed string) bool {
	rawPwd := []byte(raw)
	hashedPwd := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(hashedPwd, rawPwd)
	if err != nil {
		return false
	}
	return true
}

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// 小写
func Sha256Encode(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

// 大写
func SHA256Encode(data string) string {
	return strings.ToUpper(Sha256Encode(data))
}

// 加密
func MakePassword(plainpwd, salt string) string {
	return Sha256Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt string, password string) bool {
	sha256Pwd := Sha256Encode(plainpwd + salt)
	fmt.Println(sha256Pwd + "				" + password)
	return sha256Pwd == password
}

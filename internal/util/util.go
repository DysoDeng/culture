package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// 生成密码
func GeneratePassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// 验证密码
func ComparePassword(hashedPassword string, plainPassword string) bool {
	byteHashByte := []byte(hashedPassword)
	plainPasswordByte := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(byteHashByte, plainPasswordByte)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// 生成指定长度数字字符串
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// 生成随机字符串
func RandStringBytesMask(length int) string {

	str := make([]byte, length)

	for i, cache, reMain := length-1, rand.Int63(), letterIdxMax; i >= 0; {
		if reMain == 0 {
			cache, reMain = rand.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			str[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		reMain--
	}

	return string(str)
}

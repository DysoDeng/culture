package util

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	letterBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// CstHour 东八区
const CstHour int64 = 8 * 3600

const TimeLayout = "2006-01-02 15:04:05"
const DateLayout = "2006-01-02"

// GeneratePassword 生成密码
// @param string password 明文密码
// @return string
func GeneratePassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// ComparePassword 验证密码
// @param string hashedPassword hash密码
// @param string plainPassword 明文密码
// @return bool
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

// GenValidateCode 生成指定长度数字字符串
// @param int length 生成字符串长度
// @return string
func GenValidateCode(length int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// RandStringBytesMask 生成随机字符串
// @param int length 生成字符串长度
// @return string
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

// GetLocalIp 获取本机IP地址
// @return string
func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// CreateOrderNo 生成唯一订单号
// @return string
func CreateOrderNo() string {
	sTime := time.Now().Format("20060102150405")

	t := time.Now().UnixNano()
	s := strconv.FormatInt(t, 10)
	b := string([]byte(s)[len(s)-9:])
	c := string([]byte(b)[:7])

	rand.Seed(t)

	sTime += c + strconv.FormatInt(rand.Int63n(999999-100000)+100000, 10)
	return sTime
}

// ResolveTime 将整数转换为时分秒
// @param int seconds 秒数
// @return int hour 小时数
// @return int minute 分钟数
// @return int second 秒数
func ResolveTime(seconds int) (hour, minute, second int) {
	hour = seconds / 3600
	minute = (seconds - hour*3600) / 60
	second = seconds - hour*3600 - minute*60
	return
}

// FileSizeFormat 格式化文件大小
// @param uint64 fileSize 文件大小(字节)
// @return string
func FileSizeFormat(fileSize uint64) string {
	var byteNum float64 = 1024 // byte
	size := float64(fileSize)

	if size < byteNum { // B
		return fmt.Sprintf("%f", size) + "B"
	} else if size < math.Pow(byteNum, 2) { // KB
		return fmt.Sprintf("%.2f", size/byteNum) + "KB"
	} else if size < math.Pow(byteNum, 3) { // MB
		return fmt.Sprintf("%.2f", size/math.Pow(byteNum, 2)) + "MB"
	} else if size < math.Pow(byteNum, 4) { // GB
		return fmt.Sprintf("%.2f", size/math.Pow(byteNum, 3)) + "GB"
	}
	return fmt.Sprintf("%.2f", size/math.Pow(byteNum, 4)) + "TB"
}

// FindIds 查找整形ID
func FindIds(id uint64, ids []uint64) bool {
	for _, l := range ids {
		if id == l {
			return true
		}
	}
	return false
}

// FindStringIds 查找分布式ID
func FindStringIds(id string, ids []string) bool {
	for _, l := range ids {
		if id == l {
			return true
		}
	}
	return false
}

package utils

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

// IsEmpty 判断字符串是否为空
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// ContainsString 是否包含子串
func ContainsString(a []string, x string) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

// ContainsInt64 是否包含 int64 数字
func ContainsInt64(a []int64, x int64) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

// ContainsInt8 是否包含 int8 数字
func ContainsInt8(a []int8, x int8) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

// RandomStr 随机字符串
func RandomStr(n int) string {
	if n <= 0 {
		return ""
	}

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	return string(result)
}

// RandomInt 随机整数，最大的 int 9223372036854775807，n 最大为 19，math.MaxInt
func RandomInt(n int) int {
	if n <= 0 {
		return 0
	}

	str := "123456789"
	bytes := []byte(str)
	var result []byte
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	num, _ := strconv.Atoi(string(result))

	return num
}

// GenXID 生成唯一 ID, ref: https://github.com/rs/xid
func GenXID() string {
	return xid.New().String()
}

// NewUuid returns an uuid string.
func NewUuid() string {
	return uuid.New().String()
}

func GetString(ctx context.Context, key string) string {
	if val, ok := ctx.Value(key).(string); ok && val != "" {
		return val
	}
	return ""
}

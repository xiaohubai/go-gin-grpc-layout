package strings

import "strings"

// 判断字符串是否空
func IsEmpty(str string) bool {
	return len(str) == 0 || str == "" || strings.TrimSpace(str) == ""
}

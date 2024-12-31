package json

import "encoding/json"

// Marshal 把数据转换成json字符串
func Marshal(data interface{}) string {
	buf, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(buf)
}

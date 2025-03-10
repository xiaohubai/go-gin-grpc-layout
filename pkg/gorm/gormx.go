package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/spf13/cast"
)

// JSONField 通用的JSON字段类型
type JSONField[T any] struct {
	Data T
}

// Scan 实现 sql.Scanner 接口，将 json 字符串转换为自身
func (j *JSONField[T]) Scan(value any) error {
	if value == nil {
		*j = JSONField[T]{}
		return nil
	}

	var b []byte
	switch t := value.(type) {
	case string:
		b = []byte(t)
	case []byte:
		b = t
	default:
		return errors.New("type assertion failed, not a string or []byte")
	}

	if cast.ToString(b) == "null" {
		*j = JSONField[T]{}
		return nil
	}

	return json.Unmarshal(b, &j.Data)
}

// Value 实现 driver.Valuer 接口，将自身转换为 json 字符串
func (j JSONField[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Data)
}

// MarshalJSON 实现 json.Marshaler 接口，将自身转换为 json 字符串
func (j JSONField[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Data)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口，将 json 字符串转换为自身
func (j *JSONField[T]) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json: UnmarshalJSON on nil pointer")
	}

	return json.Unmarshal(data, &j.Data)
}

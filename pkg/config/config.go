package config

import (
	"errors"

	"github.com/spf13/viper"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils/strings"
)

func Read[T any](conf string) (*T, error) {
	if strings.IsEmpty(conf) {
		return nil, errors.New("config file path is empty")
	}

	var t T

	err := newFileConf(conf, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func newFileConf(filePath string, c any) error {
	v := viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(c); err != nil {
		return err
	}
	return nil
}

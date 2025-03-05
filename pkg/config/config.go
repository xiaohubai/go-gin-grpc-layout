package config

import (
	"errors"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils"
)

func Read[T any](filePath string) (*T, error) {
	if utils.IsEmpty(filePath) {
		return nil, errors.New("config file path is empty")
	}

	var t T
	err := newFileConf(filePath, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func newFileConf(filePath string, c any) error {
	ext := filepath.Ext(filePath)
	if !utils.ContainsString([]string{"yaml", "yml", "toml"}, ext) {
		return errors.New("unsupported config file extension: " + ext)
	}

	v := viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType(ext)

	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(c); err != nil {
		return err
	}

	return nil
}

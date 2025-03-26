package config

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils"
)

var appConf = flag.String("app_conf", "./configs/config.yaml", "app config file")

func Init() error {
	var c *Conf
	if err := newFileConf(*appConf, c); err != nil {
		return err
	}
	if c.IsRelease() {
		// TODO 读取远程配置
	}
	conf = c
	return nil
}

func newFileConf(filePath string, c any) error {
	ext := filepath.Ext(filePath)[1:]

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

package config

import (
	"context"
	"errors"
	"flag"
	"time"

	"github.com/spf13/viper"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/etcd"
)

type Configs struct {
	App    `json:"app" yaml:"app"`
	Remote `json:"remote" yaml:"remote"`
}

type App struct {
	Name    string `json:"name" yaml:"name"`
	Env     string `json:"env" yaml:"env"`
	Version string `json:"version" yaml:"version"`
	ID      string `json:"id" yaml:"id"`
}

type Remote struct {
	Type      string        `json:"type" yaml:"type"`
	Endpoints []string      `json:"endpoints" yaml:"endpoints"`
	Timeout   time.Duration `json:"timeout" yaml:"timeout"`
}

var configs string
var c Configs

func GetAppConfigs() *Configs {
	return &c
}

func InitConfig[T any](ctx context.Context, conf ...string) (*T, error) {
	flag.StringVar(&configs, "configs", "./configs/app.toml", "app config file")
	flag.Parse()
	if err := newFileConf(ctx, configs, &c); err != nil {
		return nil, err
	}

	var t T
	if c.App.Env == "test" {
		if err := newFileConf(ctx, conf[0], &t); err != nil {
			return nil, err
		}
	}
	if c.App.Env == "prod" {
		if err := newRemoteConf(ctx, &c.Remote, &t); err != nil {
			return nil, err
		}
	}
	return &t, nil
}

func newFileConf(ctx context.Context, filePath string, c any) error {
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

func newRemoteConf(ctx context.Context, remote *Remote, c any) error {
	switch remote.Type {
	case "etcd":
		conf := etcd.Conf{
			Endpoints: remote.Endpoints,
			Timeout:   remote.Timeout,
		}
		e, err := etcd.NewConf(conf)
		if err != nil {
			return err
		}
		e.kv = c
		ss, err := e.GetKV(ctx, "configs")
		return
	}
	return errors.New("empty remote type source")
}

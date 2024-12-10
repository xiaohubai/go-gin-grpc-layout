package etcd

import (
	"bytes"
	"context"
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Conf struct {
	Endpoints []string      `json:"endpoints"`
	Timeout   time.Duration `json:"timeout"`
	KV        any           `json:"kv"`
}

type Etcd[T any] struct {
	c  *clientv3.Client
	kv T
}

func NewConf(conf Conf) (*Etcd[any], error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints, // etcd 服务地址
		DialTimeout: conf.Timeout,   // 连接超时时间
	})
	if err != nil {
		return nil, err
	}
	s := conf.KV
	return &Etcd[s]{c: cli}, nil
}

func (e *Etcd[T]) GetKV(ctx context.Context, name string) (*T, error) {
	kv := clientv3.NewKV(e.c)
	resp, err := kv.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, errors.New("未找到键值")
	}

	vp := viper.New()
	confType := strings.TrimPrefix(filepath.Ext(name), ".")
	vp.SetConfigType(confType)
	vp.SetConfigFile(name)
	err = vp.ReadConfig(bytes.NewBuffer(resp.Kvs[0].Value))
	if err != nil {
		return nil, errors.New("Viper解析配置失败")
	}
	if err := vp.Unmarshal(e.kv); err != nil {
		return nil, err
	}
	return &e.kv, nil
}

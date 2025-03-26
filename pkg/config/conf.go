package config

import "time"

type Conf struct {
	App           App           `json:"app" yaml:"app"`
	Remote        Remote        `json:"remote" yaml:"remote"`
	Server        Server        `json:"server" yaml:"server"`
	Data          Data          `json:"data" yaml:"data"`
	Log           Log           `json:"logs" yaml:"logs"`
	OpenTelemetry OpenTelemetry `json:"opentelemetry" yaml:"opentelemetry"`
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

type Server struct {
	HTTP HTTP `json:"http" yaml:"http"`
	GRPC GRPC `json:"grpc" yaml:"grpc"`
}

type HTTP struct {
	Name    string `json:"name" yaml:"name"`
	Network string `json:"network" yaml:"network"`
	Addr    string `json:"addr" yaml:"addr"`
	Timeout string `json:"timeout" yaml:"timeout"`
}

type GRPC struct {
	Name    string `json:"name" yaml:"name"`
	Network string `json:"network" yaml:"network"`
	Addr    string `json:"addr" yaml:"addr"`
	Timeout string `json:"timeout" yaml:"timeout"`
}

type OpenTelemetry struct {
	Jaeger     Jaeger     `json:"jaeger" yaml:"jaeger"`
	Prometheus Prometheus `json:"prometheus" yaml:"prometheus"`
}

type Jaeger struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

type Prometheus struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

// LogsConfig 表示日志配置的结构体
type Log struct {
	Level      string `json:"level" yaml:"level"`
	Format     string `json:"format" yaml:"format"`
	FileName   string `json:"fileName" yaml:"fileName"`
	MaxSize    int    `json:"maxSize" yaml:"maxSize"`
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"`
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

type Data struct {
	MySQL MySQL `json:"mysql" yaml:"mysql"`
	Redis Redis `json:"redis" yaml:"redis"`
}

// MySQLConfig 表示 MySQL 数据库配置的结构体
type MySQL struct {
	Driver string `json:"driver" yaml:"driver"`
	Source string `json:"source" yaml:"source"`
}

// RedisConfig 表示 Redis 数据库配置的结构体
type Redis struct {
	DB           int    `json:"db" yaml:"db"`
	Addr         string `json:"addr" yaml:"addr"`
	Password     string `json:"password" yaml:"password"`
	DialTimeout  string `json:"dialTimeout" yaml:"dialTimeout"`
	ReadTimeout  string `json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout string `json:"writeTimeout" yaml:"writeTimeout"`
}

var conf *Conf

func GetConfig() *Conf {
	return conf
}

func (c *Conf) IsTest() bool {
	return c.App.Env == "test"
}

func (c *Conf) IsRelease() bool {
	return c.App.Env == "release"
}

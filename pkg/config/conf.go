package config

import "time"

type Conf struct {
	App    `json:"app" yaml:"app"`
	Remote `json:"remote" yaml:"remote"`
	Server Server `json:"server" yaml:"server"`
	Data   Data   `json:"data" yaml:"data"`
	Log    Log    `json:"logs" yaml:"logs"`
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

// ServerConfig 表示服务器配置的结构体
type Server struct {
	HTTP HTTP `json:"http" yaml:"http"`
	GRPC GRPC `json:"grpc" yaml:"grpc"`
}

// HTTPConfig 表示 HTTP 服务器配置的结构体
type HTTP struct {
	Network string `json:"network" yaml:"network"`
	Addr    string `json:"addr" yaml:"addr"`
	Timeout string `json:"timeout" yaml:"timeout"`
}

// GRPCConfig 表示 gRPC 服务器配置的结构体
type GRPC struct {
	Network string `json:"network" yaml:"network"`
	Addr    string `json:"addr" yaml:"addr"`
	Timeout string `json:"timeout" yaml:"timeout"`
}

// DataConfig 表示数据源配置的结构体
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

var conf Conf

func GetConfig() *Conf {
	return &conf
}

// 获取env
func GetEnv() string {
	return conf.App.Env
}

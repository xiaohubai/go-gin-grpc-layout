package log

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/xiaohubai/go-gin-grpc-layout/internal/pkg/conf"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(c *conf.Conf) error {
	// 配置日志文件写入器
	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", c.Logs.Filename, time.Now().Format("2006010213")),
		MaxSize:    int(c.Logs.MaxSize),
		MaxBackups: int(c.Logs.MaxBackups),
		MaxAge:     int(c.Logs.MaxAge),
		Compress:   c.Logs.Compress,
	}

	// 设置日志级别
	var level slog.Level
	switch c.Logs.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		return fmt.Errorf("unsupported log level: %s", c.Logs.Level)
	}

	slog.New(slog.NewJSONHandler(fileWriter, nil)).With(slog.LevelKey, level)

	return nil
}

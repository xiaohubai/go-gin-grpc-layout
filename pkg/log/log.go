package log

import (
	"fmt"
	"time"

	"fmt"
	"log/slog"
	"time"

	"github.com/natefinch/lumberjack"
)

func New(c *conf.Logs, g *conf.Global) (slog.Logger, error) {
	// 配置日志文件写入器
	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", c.Filename, time.Now().Format("2006-01-02")),
		MaxSize:    int(c.MaxSize),
		MaxBackups: int(c.MaxBackups),
		MaxAge:     int(c.MaxAge),
		Compress:   c.Compress,
	}

	// 创建slog.Logger实例
	logger := slog.New(slog.NewJSONHandler(fileWriter, nil))

	// 设置日志级别
	var level slog.Level
	switch c.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		return nil, fmt.Errorf("unsupported log level: %s", c.Level)
	}
	logger.SetLevel(level)

	// 添加全局上下文信息
	logger = logger.With("env", g.Env, "service_id", g.Id, "service_name", g.AppName, "service_version", g.Version)

	// 返回一个包装了slog.Logger的接口，如果需要的话
	return logger, nil
}

package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"github.com/xiaohubai/go-gin-grpc-layout/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(cf *config.Log) error {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "requestId",
		LineEnding:     zapcore.DefaultLineEnding,                          //默认换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                      //小写
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), //输出时间
		EncodeCaller:   zapcore.ShortCallerEncoder,                         //记录调用路径
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	if err := os.MkdirAll(cf.FileName, 0755); err != nil {
		return fmt.Errorf("create log directory failed: %v", err)
	}

	// 配置日志文件写入器
	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", cf.FileName, time.Now().Format("2006010215")),
		MaxSize:    cf.MaxSize * 1024,
		MaxBackups: cf.MaxBackups,
		MaxAge:     cf.MaxAge,
		Compress:   cf.Compress,
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(cf.Level)); err != nil {
		return err
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(fileWriter),
		zapLevel,
	)

	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)

	return nil
}

func AddField(key string, value any) zapcore.Field {
	return zap.Any(key, value)
}
func Info(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("requestId", utils.GetString(ctx, "requestId")))
	zap.L().Info(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = append(fields, zap.String("requestId", utils.GetString(ctx, "requestId")))
	zap.L().Error(msg, fields...)
}

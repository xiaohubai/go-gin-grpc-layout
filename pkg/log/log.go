package log

import (
	"fmt"
	"os"
	"time"

	"github.com/xiaohubai/go-gin-grpc-layout/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(cf *config.Log) error {
	if err := os.MkdirAll(cf.FileName, 0755); err != nil {
		return fmt.Errorf("create log directory failed: %v", err)
	}

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "ts",
		LineEnding:     zapcore.DefaultLineEnding,                              //默认换行
		EncodeLevel:    zapcore.LowercaseLevelEncoder,                          //小写
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"), //输出时间
		EncodeCaller:   zapcore.ShortCallerEncoder,                             //记录调用路径
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// 配置日志文件写入器
	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", cf.FileName, time.Now().Format("2006010213")),
		MaxSize:    int(cf.MaxSize),
		MaxBackups: int(cf.MaxBackups),
		MaxAge:     int(cf.MaxAge),
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

func Info(msg string, str ...string) {
	fields := []zapcore.Field{}
	for i := 0; i < len(str); i += 2 {
		fields = append(fields, zap.Any(str[i], fields[i+1]))
	}
	zap.L().Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	zap.S().Error(msg, fields)
}

package alog

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/AgoraLab/goframework/alog/alogtypes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger 基于 zap 封装实现了 Logger 接口
type zapLogger struct {
	io.Writer
	*zap.Logger
}

func newZapLogger(
	loggerName string, // 日志实例名，用于检索日志时进行过滤；必填
	level zapcore.Level,
	formatoJson bool,
	writer io.Writer,
) (*zapLogger, error) {

	if loggerName = strings.TrimSpace(loggerName); loggerName == "" {
		return nil, errors.New("loggerName is required")
	}
	if writer == nil {
		return nil, errors.New("writer cannot be nil")
	}

	zapConfig := zap.NewProductionConfig()
	zapConfig.Level.SetLevel(level)
	if !formatoJson {
		zapConfig.Encoding = "console"
	} else {
		zapConfig.Encoding = "json"
	}
	zapLoggerInst, err := zapConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddCaller(),
		zap.ErrorOutput(zapcore.AddSync(writer)),
		zap.Fields(zap.String(alogtypes.FieldKey_LoggerName, loggerName)),
		zap.AddStacktrace(zapcore.FatalLevel),
	)
	if err != nil {
		return nil, fmt.Errorf("new zap logger failed, %v", err)
	}

	return &zapLogger{
		Writer: writer,
		Logger: zapLoggerInst,
	}, nil
}

func (l *zapLogger) Close() error {
	if l.Logger != nil {
		return l.Logger.Sync()
	}
	return nil
}

func (l *zapLogger) Enabled(level zapcore.Level) bool {
	return l.Core().Enabled(level)
}

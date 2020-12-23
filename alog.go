package alog

import (
	"errors"
	"io"

	"go.uber.org/zap/zapcore"
)

const (
	FieldKey_TraceID = "traceid"
)

var (
	ErrUnsupported = errors.New("unsupported")
)

// Logger 抽象出了日志组件的基本接口
type Logger interface {
	LogCloser
	LogLevelEnabler
	LogLevelHotReloader

	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

// LogCloser 关闭日志组件
type LogCloser interface {
	io.Closer
}

// LogLevelEnabler 判断是否启用了某种级别的日志输出
type LogLevelEnabler interface {
	zapcore.LevelEnabler
}

// LogLevelHotReloader 热更日志级别
// 如果不支持热日志级别，则返回错误 ErrUnsupported
type LogLevelHotReloader interface {
	HotReloadLogLevel(level zapcore.Level) error
}

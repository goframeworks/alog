package alog

import (
	"io"

	"go.uber.org/zap/zapcore"
)

// Factory 抽象出了 alog 的工厂接口
type Factory interface {
	// ZapLogger 基于zap封装的日志组件，输出为结构化数据
	// loggerName: 当前日志实例的名字，一般可以写当前服务名字，用于做日志筛选
	// formatoJson: 是否输出为json格式的日志
	ZapLogger(loggerName string, level zapcore.Level, formatoJson bool) (Logger, error)
}

// Logger 抽象出了日志组件的基本接口
type Logger interface {
	LogWriter
	LogCloser
	zapcore.LevelEnabler

	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
}

type LogWriter interface {
	io.Writer
}

type LogCloser interface {
	io.Closer
}

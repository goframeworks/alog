package alog

import (
	"os"

	"go.uber.org/zap/zapcore"
)

// facotory 实现了 Factory 接口
type factory struct{}

func NewFactory() Factory {
	return &factory{}
}

func (f *factory) ZapLogger(loggerName string, level zapcore.Level, formatoJson bool) (Logger, error) {
	return newZapLogger(loggerName, level, formatoJson, os.Stderr)
}

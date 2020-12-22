// +build unittest

package alog

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestFactory_ZapLogger(t *testing.T) {
	logger, err := NewFactory().ZapLogger(t.Name(), zapcore.DebugLevel, false)
	if err != nil {
		t.Fatalf("new logger failed, %v", err)
	}
	logger.Debug("this is debug msg", zap.String("hello", "debug"))
	logger.Info("this is info msg", zap.String("hello", "info"))
	logger.Warn("this is warn msg", zap.String("hello", "warn"))
	logger.Error("this is error msg", zap.String("hello", "error"))

	////////////////////////////////////////////////////////////////////////////

	logger, err = NewFactory().ZapLogger(t.Name(), zapcore.WarnLevel, true)
	if err != nil {
		t.Fatalf("new logger failed, %v", err)
	}
	logger.Debug("this is debug msg", zap.String("hello", "debug"))
	logger.Info("this is info msg", zap.String("hello", "info"))
	logger.Warn("this is warn msg", zap.String("hello", "warn"))
	logger.Error("this is error msg", zap.String("hello", "error"))
}

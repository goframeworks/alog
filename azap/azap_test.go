// +build unittest

package azap

import (
	"os"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestFactory_ZapLogger(t *testing.T) {
	logger, err := NewLogger(t.Name(), WithLogLevel(zapcore.DebugLevel), WithWriter(os.Stderr), WithStructuredFormat(false))
	if err != nil {
		t.Fatalf("new logger failed, %v", err)
	}
	logger.Debug("this is debug msg", zap.String("hello", "debug"))
	logger.Info("this is info msg", zap.String("hello", "info"))
	logger.Warn("this is warn msg", zap.String("hello", "warn"))
	logger.Error("this is error msg", zap.String("hello", "error"))

	////////////////////////////////////////////////////////////////////////////

	logger, err = NewLogger(t.Name(), WithLogLevel(zapcore.WarnLevel), WithWriter(os.Stderr), WithStructuredFormat(true))
	if err != nil {
		t.Fatalf("new logger failed, %v", err)
	}
	logger.Debug("this is debug msg", zap.String("hello", "debug"))
	logger.Info("this is info msg", zap.String("hello", "info"))
	logger.Warn("this is warn msg", zap.String("hello", "warn"))
	logger.Error("this is error msg", zap.String("hello", "error"))
}

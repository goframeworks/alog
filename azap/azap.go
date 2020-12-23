package azap

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/goframeworks/alog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger 基于 zap 封装的日志接口
type ZapLogger interface {
	alog.Logger
	ZapLoggerInitializer
}

type ZapLoggerInitializer interface {
	InitLogLevel(v zapcore.Level)
	InitWriter(w io.Writer)
	InitStructuredFormat(v bool)
}

// zapLogger 基于 zap 封装实现了 Logger 接口
type zapLogger struct {
	*zap.Logger
	zapConfig zap.Config

	logLevel zapcore.Level
	writer   io.Writer
	encoding string
}

// NewLogger 新建基于zap的日志实例
// 默认输出到 stderr
// 默认日志级别 INFO
// 默认编码方式 JSON
func NewLogger(loggerName string, opts ...ZapLoggerOption) (ZapLogger, error) {
	if loggerName = strings.TrimSpace(loggerName); loggerName == "" {
		return nil, errors.New("loggerName is required, we suggest to use service name")
	}

	logger := &zapLogger{
		logLevel: zapcore.InfoLevel,
		writer:   os.Stderr,
		encoding: "json",
	}
	for _, opt := range opts {
		opt(logger)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(logger.logLevel)
	zapConfig.Encoding = logger.encoding
	zapConfig.EncoderConfig = encoderConfig

	zapLoggerInst, err := zapConfig.Build(
		zap.AddCallerSkip(1),
		zap.AddCaller(),
		zap.ErrorOutput(zapcore.AddSync(logger.writer)),
		zap.AddStacktrace(zapcore.FatalLevel),
	)
	if err != nil {
		return nil, fmt.Errorf("new zap logger failed, %v", err)
	}

	logger.zapConfig = zapConfig
	logger.Logger = zapLoggerInst.Named(loggerName)
	return logger, nil
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

func (l *zapLogger) HotReloadLogLevel(level zapcore.Level) error {
	l.zapConfig.Level.SetLevel(level)
	return nil
}

func (l *zapLogger) InitLogLevel(v zapcore.Level) {
	l.logLevel = v
}

func (l *zapLogger) InitWriter(w io.Writer) {
	if w == nil {
		l.writer = os.Stderr
	} else {
		l.writer = w
	}
}

func (l *zapLogger) InitStructuredFormat(v bool) {
	if v {
		l.encoding = "json"
	} else {
		l.encoding = "console"
	}
}

// ZapLoggerOption zap 日志组件的可选能力
type ZapLoggerOption func(logger ZapLoggerInitializer)

// WithLogLevel 初始化日志级别
func WithLogLevel(v zapcore.Level) ZapLoggerOption {
	return func(logger ZapLoggerInitializer) {
		logger.InitLogLevel(v)
	}
}

// WithWriter 设置日志写入
func WithWriter(w io.Writer) ZapLoggerOption {
	return func(logger ZapLoggerInitializer) {
		logger.InitWriter(w)
	}
}

// WithStructuredFormat 设置是否启用结构化输出格式
func WithStructuredFormat(v bool) ZapLoggerOption {
	return func(logger ZapLoggerInitializer) {
		logger.InitStructuredFormat(v)
	}
}

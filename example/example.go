package main

import (
	"os"
	"time"

	"github.com/goframeworks/alog"
	"github.com/goframeworks/alog/azap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func AZapExample1() {
	traceId := "8b168be0-a7d5-47f4-8e20-a2afe4269662"
	logger, err := azap.NewLogger("azap-example1",
		azap.WithLogLevel(zapcore.DebugLevel),
		azap.WithWriter(os.Stderr),
		azap.WithStructuredFormat(false),
	)
	if err != nil {
		panic(err)
	}

	logger.Debug("this is azap example1", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Info("this is azap example1", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Warn("this is azap example1", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Error("this is azap example1", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Fatal("this is azap example1", zap.String(alog.FieldKey_TraceID, traceId))
}

func AZapExample2() {
	traceId := "5b61e94d-3a59-468e-9198-8d662119c01e"
	logger, err := azap.NewLogger("azap-example2",
		azap.WithLogLevel(zapcore.WarnLevel),
		azap.WithWriter(os.Stderr),
		azap.WithStructuredFormat(false),
	)
	if err != nil {
		panic(err)
	}

	logger.Debug("this is azap example2", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Info("this is azap example2", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Warn("this is azap example2", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Error("this is azap example2", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Fatal("this is azap example2", zap.String(alog.FieldKey_TraceID, traceId))
}

func AZapExample3() {
	traceId := "be0c950e-bd99-44dd-9d68-d7d494b62276"
	logger, err := azap.NewLogger("azap-example3",
		azap.WithLogLevel(zapcore.DebugLevel),
		azap.WithWriter(os.Stderr),
		azap.WithStructuredFormat(true),
	)
	if err != nil {
		panic(err)
	}

	logger.Debug("this is azap example3", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Info("this is azap example3", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Warn("this is azap example3", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Error("this is azap example3", zap.String(alog.FieldKey_TraceID, traceId))
	logger.Fatal("this is azap example3", zap.String(alog.FieldKey_TraceID, traceId))
}

func AzapLevelReload() {
	traceId := "be0c950e-bd99-44dd-9d68-d7d494b62276"
	logger, err := azap.NewLogger("azap-level-reload",
		azap.WithLogLevel(zapcore.DebugLevel),
		azap.WithWriter(os.Stderr),
		azap.WithStructuredFormat(false),
	)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			logger.Debug("debug message", zap.String(alog.FieldKey_TraceID, traceId))
			logger.Info("info message", zap.String(alog.FieldKey_TraceID, traceId))
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
	if err := logger.HotReloadLogLevel(zapcore.InfoLevel); err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)
}

func main() {
	//AZapExample1()
	//AZapExample2()
	//AZapExample3()
	AzapLevelReload()
}

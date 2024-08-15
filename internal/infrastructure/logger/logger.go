package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewLogger() *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zapcore.DebugLevel)
	return zap.New(core, zap.AddCaller())
}

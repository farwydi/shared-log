package shared_log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func NewDebugLogger() *zap.Logger {
    return zap.New(zapcore.NewCore(
        zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
        zapcore.AddSync(os.Stdout),
        zap.NewAtomicLevelAt(zap.DebugLevel),
    ))
}
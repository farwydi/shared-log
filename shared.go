package shared_log

import "go.uber.org/zap"

func Info(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Errof(err error, fields ...zap.Field) {
	zap.L().Error(err.Error(), fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zap.L().Panic(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	zap.L().Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zap.L().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

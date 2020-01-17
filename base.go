package shared_log

import "go.uber.org/zap"

type IBaseLogger interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
}

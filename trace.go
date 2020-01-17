package shared_log

import "go.uber.org/zap"

type traceLog struct {
	defaultFields []zap.Field
}

func (t traceLog) Errof(err error, fields ...zap.Field) {
	Errof(err, append(t.defaultFields, fields...)...)
}

func (t traceLog) Warn(msg string, fields ...zap.Field) {
	Warn(msg, append(t.defaultFields, fields...)...)
}

func (t traceLog) Error(msg string, fields ...zap.Field) {
	Error(msg, append(t.defaultFields, fields...)...)
}

func (t traceLog) Debug(msg string, fields ...zap.Field) {
	Debug(msg, append(t.defaultFields, fields...)...)
}

func (t traceLog) Panic(msg string, fields ...zap.Field) {
	Panic(msg, append(t.defaultFields, fields...)...)
}

func (t traceLog) Info(msg string, fields ...zap.Field) {
	Info(msg, append(t.defaultFields, fields...)...)
}

func NewTrace(defaultFields ...zap.Field) IBaseLogger {
	return &traceLog{
		defaultFields: defaultFields,
	}
}

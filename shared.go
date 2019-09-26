package shared_log

import "go.uber.org/zap"

var _logger = NewDebugLogger()

func ReplaceLogger(logger *zap.Logger) {
    _logger = logger
}

func Extract() *zap.Logger {
    return _logger
}

func Info(msg string, fields ...zap.Field) {
    if _logger != nil {
        _logger.Info(msg, fields...)
    }
}

func Errof(err error) {
    if _logger != nil {
        _logger.Error(err.Error())
    }
}

func Panic(msg string, fields ...zap.Field) {
    if _logger != nil {
        _logger.Panic(msg, fields...)
    }
}

func Debug(msg string, fields ...zap.Field) {
    if _logger != nil {
        _logger.Debug(msg, fields...)
    }
}

func Error(msg string, fields ...zap.Field) {
    if _logger != nil {
        _logger.Info(msg, fields...)
    }
}

func Warn(msg string, fields ...zap.Field) {
    if _logger != nil {
        _logger.Warn(msg, fields...)
    }
}

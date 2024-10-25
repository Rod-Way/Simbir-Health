package logger

import (
	"context"

	"go.uber.org/zap"
)

const (
	LoggerKey   = "logger"
	RequestID   = "requestID"
	ServiceName = "account"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

type logger struct {
	serviceName string
	logger      *zap.Logger
}

// rewriting info log
func (l logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String(ServiceName, l.serviceName), zap.String(RequestID, ctx.Value(RequestID).(string)))
	l.logger.Info(msg, fields...)
}

// rewriting error log
func (l logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, zap.String(ServiceName, l.serviceName), zap.String(RequestID, ctx.Value(RequestID).(string)))
	l.logger.Error(msg, fields...)
}

// Creating new logger
func New(serviceName string) Logger {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	return &logger{
		serviceName: serviceName,
		logger:      zapLogger,
	}
}

// getting logger from context
func GetLoggerFromCtx(ctx context.Context) Logger {
	return ctx.Value(LoggerKey).(Logger)
}

package zaplogger

import (
	"fmt"

	"go.uber.org/zap"
)

// Print logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Print(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// Debug logs a message at level Debug on the ZapLogger.
func (l *ZapLogger) Debug(args ...interface{}) {
	l.Log.Debug(fmt.Sprint(args...))
}

// Info logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Info(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warn(args ...interface{}) {
	l.Log.Warn(fmt.Sprint(args...))
}

// Error logs a message at level Error on the ZapLogger.
func (l *ZapLogger) Error(args ...interface{}) {
	l.Log.Error(fmt.Sprint(args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func (l *ZapLogger) Fatal(args ...interface{}) {
	l.Log.Fatal(fmt.Sprint(args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) Panic(args ...interface{}) {
	l.Log.Panic(fmt.Sprint(args...))
}

// With return a logger with an extra field.
func (l *ZapLogger) With(key string, value interface{}) *ZapLogger {
	return &ZapLogger{l.Log.With(zap.Any(key, value))}
}

// WithField return a logger with an extra field.
func (l *ZapLogger) WithField(key string, value interface{}) *ZapLogger {
	return &ZapLogger{l.Log.With(zap.Any(key, value))}
}

// WithFields return a logger with extra fields.
func (l *ZapLogger) WithFields(fields map[string]interface{}) *ZapLogger {
	i := 0
	var clog *ZapLogger
	for k, v := range fields {
		if i == 0 {
			clog = l.WithField(k, v)
		} else {
			clog = clog.WithField(k, v)
		}
		i++
	}
	return clog
}

package zaplogger

import (
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger is a logger which compatible to logrus/std zlog/prometheus.
// it implements Print() Println() Printf() Dbug() Debugln() Debugf() Info() Infoln() Infof() Warn() Warnln() Warnf()
// Error() Errorln() Errorf() Fatal() Fataln() Fatalf() Panic() Panicln() Panicf() With() WithField() WithFields()
type ZapLogger struct {
	zlog *zap.Logger
}

// NewZapLogger return ZapLogger with caller field
func NewZapLogger(debugLevel bool) *ZapLogger {
	return &ZapLogger{NewLogger(debugLevel).WithOptions(zap.AddCallerSkip(1))}
}

// InitZaplogger
func InitZapLogger(log *zap.Logger) *ZapLogger {
	return &ZapLogger{
		log,
	}
}

// Print logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Print(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Println logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Println(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Printf logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Printf(format string, args ...interface{}) {
	l.zlog.Info(fmt.Sprintf(format, args...))
}

// Debug logs a message at level Debug on the ZapLogger.
func (l *ZapLogger) Debug(args ...interface{}) {
	l.zlog.Debug(fmt.Sprint(args...))
}

// Debugln logs a message at level Debug on the ZapLogger.
func (l *ZapLogger) Debugln(args ...interface{}) {
	l.zlog.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level Debug on the ZapLogger.
func (l *ZapLogger) Debugf(format string, args ...interface{}) {
	l.zlog.Debug(fmt.Sprintf(format, args...))
}

// Info logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Info(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Infoln logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Infoln(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Infof(format string, args ...interface{}) {
	l.zlog.Info(fmt.Sprintf(format, args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warn(args ...interface{}) {
	l.zlog.Warn(fmt.Sprint(args...))
}

// Warnln logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warnln(args ...interface{}) {
	l.zlog.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warnf(format string, args ...interface{}) {
	l.zlog.Warn(fmt.Sprintf(format, args...))
}

// Error logs a message at level Error on the ZapLogger.
func (l *ZapLogger) Error(args ...interface{}) {
	l.zlog.Error(fmt.Sprint(args...))
}

// Errorln logs a message at level Error on the ZapLogger.
func (l *ZapLogger) Errorln(args ...interface{}) {
	l.zlog.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Error on the ZapLogger.
func (l *ZapLogger) Errorf(format string, args ...interface{}) {
	l.zlog.Error(fmt.Sprintf(format, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func (l *ZapLogger) Fatal(args ...interface{}) {
	l.zlog.Fatal(fmt.Sprint(args...))
}

// Fatalln logs a message at level Fatal on the ZapLogger.
func (l *ZapLogger) Fatalln(args ...interface{}) {
	l.zlog.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Fatal on the ZapLogger.
func (l *ZapLogger) Fatalf(format string, args ...interface{}) {
	l.zlog.Fatal(fmt.Sprintf(format, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) Panic(args ...interface{}) {
	l.zlog.Panic(fmt.Sprint(args...))
}

// Panicln logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) Panicln(args ...interface{}) {
	l.zlog.Panic(fmt.Sprint(args...))
}

// Panicf logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) Panicf(format string, args ...interface{}) {
	l.zlog.Panic(fmt.Sprintf(format, args...))
}

// With return a logger with an extra field.
func (l *ZapLogger) With(key string, value interface{}) *ZapLogger {
	return &ZapLogger{l.zlog.With(zap.Any(key, value))}
}

// WithField return a logger with an extra field.
func (l *ZapLogger) WithField(key string, value interface{}) *ZapLogger {
	return &ZapLogger{l.zlog.With(zap.Any(key, value))}
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

package zaplogger

import (
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger is a logger which compatible to logrus/std zlog/prometheus.
// it implements Print() Println() Printf() Dbug() Debugln() Debugf() Info() Infoln() Infof() Warn() Warnln() Warnf()
// Error() Errorln() Errorf() Fatal() Fataln() Fatalf() Panic() Panicln() Panicf() With() WithField() WithFields()

type ZapLogger struct {
	Log *zap.Logger
}

// NewProduction new log for production
func NewProduction() *ZapLogger {
	log, _ := zap.NewProduction()
	return &ZapLogger{
		Log: log,
	}
}

// NewDevelopment new log for development
func NewDevelopment() *ZapLogger {
	log, _ := zap.NewDevelopment()
	return &ZapLogger{
		Log: log,
	}
}

// InitZaoLogger initial
func InitZaoLogger(log *zap.Logger) *ZapLogger {
	return &ZapLogger{
		log,
	}
}

// NewZapLogger return ZapLogger with caller field
func NewZapLogger() *ZapLogger {
	return &ZapLogger{NewLogger().WithOptions(zap.AddCallerSkip(1))}
}

// Debug logs a message at level Debug on the ZapLogger.
func (l *ZapLogger) Debug(args ...interface{}) {
	l.Log.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level Debug on the ZapLogger.
func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.Log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Info(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Infof(template string, args ...interface{}) {

	l.Log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warn(args ...interface{}) {
	l.Log.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Warnf(template string, args ...interface{}) {

	l.Log.Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at level Error on the ZapLogger.
func (l *ZapLogger) Error(args ...interface{}) {
	l.Log.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Errorf(template string, args ...interface{}) {

	l.Log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func (l *ZapLogger) Fatal(args ...interface{}) {
	l.Log.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Fatalf(template string, args ...interface{}) {

	l.Log.Fatal(fmt.Sprintf(template, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func (l *ZapLogger) Panic(args ...interface{}) {
	l.Log.Panic(fmt.Sprint(args...))
}

// Panicf  logs a message at level Warn on the ZapLogger.
func (l *ZapLogger) Panicf(template string, args ...interface{}) {

	l.Log.Panic(fmt.Sprintf(template, args...))
}

// With return a logger with an extra field.
func (l *ZapLogger) With(key string, value interface{}) *ZapLogger {
	return &ZapLogger{l.Log.With(zap.Any(key, value))}
}

// Printf logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Printf(format string, args ...interface{}) {
	l.Log.Info(fmt.Sprintf(format, args...))
}

// Print logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Print(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
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

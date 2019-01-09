package zaplogger

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"runtime"
	"strings"
)

// CallerEncoder will add caller to log. format is "filename:lineNum:funcName", e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

//
func NewLoggerConfig(debugLevel bool) (loggerConfig zap.Config) {
	loggerConfig = zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.EncodeCaller = CallerEncoder
	if debugLevel {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return
}

// NewCustomLoggers is a shortcut to get normal logger, noCallerLogger.
func NewCustomLoggers(debugLevel bool) (logger, noCallerLogger *zap.Logger) {
	loggerConfig := NewLoggerConfig(debugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	loggerConfig.DisableCaller = true
	noCallerLogger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewLogger return a normal logger
func NewLogger(debugLevel bool) (logger *zap.Logger) {
	loggerConfig := NewLoggerConfig(debugLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewNoCallerLogger return a no caller key value, will be faster
func NewNoCallerLogger(debugLevel bool) (noCallerLogger *zap.Logger) {
	loggerConfig := NewLoggerConfig(debugLevel)
	loggerConfig.DisableCaller = true
	noCallerLogger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// ZapLogger is a logger which compatible to logrus/std log/prometheus.
// it implements Print() Println() Printf() Dbug() Debugln() Debugf() Info() Infoln() Infof() Warn() Warnln() Warnf()
// Error() Errorln() Errorf() Fatal() Fataln() Fatalf() Panic() Panicln() Panicf() With() WithField() WithFields()
type ZapLogger struct {
	zlog *zap.Logger
}

// NewCompatibleLogger return ZapLogger with caller field
func NewCompatibleLogger(debugLevel bool) *ZapLogger {
	return &ZapLogger{NewLogger(debugLevel).WithOptions(zap.AddCallerSkip(1))}
}

// Print logs a message at level Info on the compatibleLogger.
func (l ZapLogger) Print(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Println logs a message at level Info on the compatibleLogger.
func (l ZapLogger) Println(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Printf logs a message at level Info on the compatibleLogger.
func (l ZapLogger) Printf(format string, args ...interface{}) {
	l.zlog.Info(fmt.Sprintf(format, args...))
}

// Debug logs a message at level Debug on the compatibleLogger.
func (l ZapLogger) Debug(args ...interface{}) {
	l.zlog.Debug(fmt.Sprint(args...))
}

// Debugln logs a message at level Debug on the compatibleLogger.
func (l ZapLogger) Debugln(args ...interface{}) {
	l.zlog.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level Debug on the compatibleLogger.
func (l ZapLogger) Debugf(format string, args ...interface{}) {
	l.zlog.Debug(fmt.Sprintf(format, args...))
}

// Info logs a message at level Info on the compatibleLogger.
func (l ZapLogger) Info(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Infoln logs a message at level Info on the compatibleLogger.
func (l ZapLogger) Infoln(args ...interface{}) {
	l.zlog.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the compatibleLogger.
func (l ZapLogger) Infof(format string, args ...interface{}) {
	l.zlog.Info(fmt.Sprintf(format, args...))
}

// Warn logs a message at level Warn on the compatibleLogger.
func (l ZapLogger) Warn(args ...interface{}) {
	l.zlog.Warn(fmt.Sprint(args...))
}

// Warnln logs a message at level Warn on the compatibleLogger.
func (l ZapLogger) Warnln(args ...interface{}) {
	l.zlog.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the compatibleLogger.
func (l ZapLogger) Warnf(format string, args ...interface{}) {
	l.zlog.Warn(fmt.Sprintf(format, args...))
}

// Error logs a message at level Error on the compatibleLogger.
func (l ZapLogger) Error(args ...interface{}) {
	l.zlog.Error(fmt.Sprint(args...))
}

// Errorln logs a message at level Error on the compatibleLogger.
func (l ZapLogger) Errorln(args ...interface{}) {
	l.zlog.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Error on the compatibleLogger.
func (l ZapLogger) Errorf(format string, args ...interface{}) {
	l.zlog.Error(fmt.Sprintf(format, args...))
}

// Fatal logs a message at level Fatal on the compatibleLogger.
func (l ZapLogger) Fatal(args ...interface{}) {
	l.zlog.Fatal(fmt.Sprint(args...))
}

// Fatalln logs a message at level Fatal on the compatibleLogger.
func (l ZapLogger) Fatalln(args ...interface{}) {
	l.zlog.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Fatal on the compatibleLogger.
func (l ZapLogger) Fatalf(format string, args ...interface{}) {
	l.zlog.Fatal(fmt.Sprintf(format, args...))
}

// Panic logs a message at level Painc on the compatibleLogger.
func (l ZapLogger) Panic(args ...interface{}) {
	l.zlog.Panic(fmt.Sprint(args...))
}

// Panicln logs a message at level Painc on the compatibleLogger.
func (l ZapLogger) Panicln(args ...interface{}) {
	l.zlog.Panic(fmt.Sprint(args...))
}

// Panicf logs a message at level Painc on the compatibleLogger.
func (l ZapLogger) Panicf(format string, args ...interface{}) {
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

// FormatStdLog set the output of stand package log to zaplog
func FormatStdLog() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(&logWriter{NewNoCallerLogger(false)})
}

type logWriter struct {
	logger *zap.Logger
}

// Write implement io.Writer, as std log's output
func (w logWriter) Write(p []byte) (n int, err error) {
	i := bytes.Index(p, []byte(":")) + 1
	j := bytes.Index(p[i:], []byte(":")) + 1 + i
	caller := bytes.TrimRight(p[:j], ":")
	// find last index of /
	i = bytes.LastIndex(caller, []byte("/"))
	// find penultimate index of /
	i = bytes.LastIndex(caller[:i], []byte("/"))
	w.logger.Info("stdLog", zap.ByteString("caller", caller[i+1:]), zap.ByteString("log", bytes.TrimSpace(p[j:])))
	return len(p), nil
}

package zaplogger
/**
import (
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// callerEncoder will add caller to zlog. format is "filename:lineNum:funcName", e.g:"zaplog/zaplog_test.go:15:zaplog.TestNewLogger"
func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))
}

// timeEncoder specifics the time format
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// milliSecondsDurationEncoder serializes a time.Duration to a floating-point number of micro seconds elapsed.
func milliSecondsDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

// newLoggerConfig
func newLoggerConfig(debugLevel bool, te zapcore.TimeEncoder, de zapcore.DurationEncoder) (loggerConfig zap.Config) {
	loggerConfig = zap.NewProductionConfig()
	if te == nil {
		loggerConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		loggerConfig.EncoderConfig.EncodeTime = te
	}
	if de == nil {
		loggerConfig.EncoderConfig.EncodeDuration = milliSecondsDurationEncoder
	} else {
		loggerConfig.EncoderConfig.EncodeDuration = de
	}
	loggerConfig.EncoderConfig.EncodeCaller = callerEncoder
	if debugLevel {
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	return
}

// NewLogger return a normal logger
func NewLogger(debugLevel bool) (logger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, nil, nil)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewCustomLogger return a normal logger with given timeEncoder
func NewCustomLogger(debugLevel bool, te zapcore.TimeEncoder, de zapcore.DurationEncoder) (logger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, te, de)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewNoCallerLogger return a no caller key value, will be faster
func NewNoCallerLogger(debugLevel bool) (noCallerLogger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, nil, nil)
	loggerConfig.DisableCaller = true
	noCallerLogger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	return
}

// NewNormalLoggers is a shortcut to get normal logger, noCallerLogger.
func NewNormalLoggers(debugLevel bool) (logger, noCallerLogger *zap.Logger) {
	loggerConfig := newLoggerConfig(debugLevel, nil, nil)
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
*/

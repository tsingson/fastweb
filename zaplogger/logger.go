package zaplogger

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/diode"
	"github.com/spf13/afero"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/tsingson/fastx/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// 	ProxyAddr string
	LumberLogger *lumberjack.Logger
)

// NewZapLog  initial a zap logger
func NewZapLog(path, logFileNamePrefix string, stdoutFlag bool) *zap.Logger {

	dataTimeFmtInFileName := time.Now().Format("2006-01-02-15")

	if len(path) == 0 {
		path, _ = utils.GetCurrentExecDir()
	}

	logpath := path + "/" // + dataTimeFmtInFileName
	errLogPath := path + "/err/"
	//
	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logpath)
	if !check {

		err := afs.MkdirAll(logpath, 0755)
		if err != nil {

		}
	}

	check, _ = afero.DirExists(afs, errLogPath)
	if !check {
		err := afs.MkdirAll(errLogPath, 0755)
		if err != nil {

		}
	}

	var logfilename string
	if len(logFileNamePrefix) == 0 {
		// 	logfilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"
		logfilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"

	} else {
		// 	logfilename = logpath + "/" + logFileNamePrefix + "-pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"
		logfilename = logpath + "/" + logFileNamePrefix + "-" + dataTimeFmtInFileName + ".zlog"

	}

	LumberLogger = &lumberjack.Logger{
		Filename:   logfilename,
		MaxSize:    10, // megabytes
		MaxBackups: 31,
		MaxAge:     31,    // days
		Compress:   false, // 开发时不压缩
	}

	wdiode := diode.NewWriter(LumberLogger, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})

	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	var w zapcore.WriteSyncer
	var level zapcore.Level
	if stdoutFlag {
		w = zapcore.NewMultiWriteSyncer(zapcore.AddSync(wdiode), zapcore.AddSync(os.Stdout))
		level = zapcore.InfoLevel
	} else {
		w = zapcore.NewMultiWriteSyncer(zapcore.AddSync(wdiode))
		level = zapcore.ErrorLevel
	}

	log := newZapLogger(true, false, level, w)
	log.Info("zap logger init succcess")

	return log
}

// newZapLogger
func newZapLogger(encodeAsJSON, callerFlag bool, level zapcore.Level, output zapcore.WriteSyncer) *zap.Logger {
	opts := []zap.Option{}
	if callerFlag {
		opts = append(opts, zap.AddCaller())
		opts = append(opts, zap.AddStacktrace(zap.WarnLevel))
	}
	return zap.New(newZapCore(encodeAsJSON, level, output), opts...)

}

// newZapLogger
func newZapCore(jsonFlag bool, level zapcore.Level, output zapcore.WriteSyncer) zapcore.Core {

	cfg := zapcore.EncoderConfig{
		TimeKey:        "logtime",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder //
	if jsonFlag {
		encoder = zapcore.NewJSONEncoder(cfg)
	}
	encoder = zapcore.NewConsoleEncoder(cfg)

	return zapcore.NewCore(encoder, output, zap.NewAtomicLevelAt(level))
}

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

// NewZapLog  init a log
func NewZapLog(path, prefix string, stdoutFlag bool) *zap.Logger {

	opts := []zap.Option{}

	if stdoutFlag {
		// opts = append(opts, zap.AddCaller())
		// opts = append(opts, zap.AddStacktrace(zap.WarnLevel))

		std := NewStdoutCore(zapcore.DebugLevel)
		debug := NewZapCore(path, prefix, zapcore.InfoLevel)

		return zap.New(zapcore.NewTee(std, debug), opts...)
	} else {
		errlog := NewZapCore(path, prefix, zapcore.ErrorLevel)
		return zap.New(errlog)
	}

}

// NewZapLog  initial a zap logger
func NewZapCore(path, prefix string, level zapcore.Level) zapcore.Core {

	dataTimeFmtInFileName := time.Now().Format("2006-01-02-15")
	var err error
	var logPath string

	logPath, err = buildLogPath(path)
	if err != nil {
		// TODO: handle error
	}

	var logFilename string
	if len(prefix) == 0 {
		// 	logFilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"
		logFilename = logPath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"

	} else {
		// 	logFilename = logpath + "/" + prefix + "-pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeFmtInFileName + ".zlog"
		logFilename = logPath + "/" + prefix + "-" + dataTimeFmtInFileName + ".zlog"

	}
	var LumberLogger *lumberjack.Logger
	LumberLogger = &lumberjack.Logger{
		Filename:   logFilename,
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
	w = zapcore.AddSync(wdiode)

	return newZapCore(true, level, w)

}

func NewStdoutCore(level zapcore.Level) zapcore.Core {
	var w zapcore.WriteSyncer

	w = zapcore.AddSync(os.Stdout)

	return newZapCore(true, level, w)
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
	} else {
		encoder = zapcore.NewConsoleEncoder(cfg)
	}

	return zapcore.NewCore(encoder, output, zap.NewAtomicLevelAt(level))
}

// buildLogPath
func buildLogPath(path string) (logPath string, err error) {
	if len(path) == 0 {
		path, _ = utils.GetCurrentExecDir()
	}
	logPath = path + "/"

	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logPath)
	if !check {
		err := afs.MkdirAll(logPath, 0755)
		if err != nil {
			return "", err
		}
	}

	tf := logPath + "test.log"
	err = afero.WriteFile(afs, tf, []byte("file b"), 0644)
	if err != nil {
		return "", err
	} else {
		afs.Remove(tf)
	}

	return logPath, nil
}

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
	var (
		logfilename string
		err         error
	)
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
		err = afs.MkdirAll(logpath, 0755)
		if err != nil {

		}
	}

	check, _ = afero.DirExists(afs, errLogPath)
	if !check {
		err = afs.MkdirAll(errLogPath, 0755)
		if err != nil {

		}
	}

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
	if stdoutFlag {
		w = zapcore.NewMultiWriteSyncer(zapcore.AddSync(wdiode), zapcore.AddSync(os.Stdout))
	} else {
		w = zapcore.NewMultiWriteSyncer(zapcore.AddSync(wdiode))
	}

	log := newZapLogger(true, false, zapcore.ErrorLevel, w)
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
	return zap.New(newZapCore(encodeAsJSON, callerFlag, level, output ), opts...)
}

// newZapLogger
func newZapCore(encodeAsJSON, callerFlag bool, level zapcore.Level, output zapcore.WriteSyncer) zapcore.Core {
	var encoder zapcore.Encoder

	encCfg := zapcore.EncoderConfig{
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

	opts := []zap.Option{}
	if callerFlag {
		opts = append(opts, zap.AddCaller())
		opts = append(opts, zap.AddStacktrace(zap.WarnLevel))
	}

	if encodeAsJSON {
		encoder = zapcore.NewJSONEncoder(encCfg)
	}
	encoder = zapcore.NewConsoleEncoder(encCfg)

	return zapcore.NewCore(encoder, output, zap.NewAtomicLevelAt(level)
}

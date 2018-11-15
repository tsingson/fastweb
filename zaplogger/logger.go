package zaplogger

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/diode"
	"github.com/spf13/afero"
	"github.com/tsingson/fastx/utils"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// 	ProxyAddr string

	LumberLogger *lumberjack.Logger
)

// NewZapLog  initial a zap logger
func NewZapLog(path, logFileNamePrefix string, stdoutFlag bool) *zap.Logger {
	var logfilename string
	dataTimeStr := time.Now().Format("2006-01-02-15")
	if len(path) == 0 {
		path, _ = utils.GetCurrentExecDir()
	}

	logpath := path + "/" // + dataTimeStr
	//
	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logpath)
	if !check {
		afs.MkdirAll(logpath, 0755)
	}

	if len(logFileNamePrefix) == 0 {
		// 	logfilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeStr + ".zlog"
		logfilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeStr + ".zlog"

	} else {
		// 	logfilename = logpath + "/" + logFileNamePrefix + "-pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeStr + ".zlog"
		logfilename = logpath + "/" + logFileNamePrefix + dataTimeStr + ".zlog"
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

	// -------------------------------------
	/**
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			LumberLogger.Rotate()
		}
	}()
	*/
	log := newZapLogger(true, w)
	log.Info("zap logger init succcess")

	return log
}

// newZapLogger
func newZapLogger(encodeAsJSON bool, output zapcore.WriteSyncer) *zap.Logger {
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

	opts := []zap.Option{zap.AddCaller()}
	opts = append(opts, zap.AddStacktrace(zap.WarnLevel))
	encoder := zapcore.NewConsoleEncoder(encCfg)
	if encodeAsJSON {
		encoder = zapcore.NewJSONEncoder(encCfg)
	}

	return zap.New(zapcore.NewCore(encoder, output, zap.NewAtomicLevelAt(zap.DebugLevel)), opts...)
}

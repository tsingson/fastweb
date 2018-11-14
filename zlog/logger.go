package zlog

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/spf13/afero"
	"github.com/tsingson/fastweb/utils"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// 日志文件的时间格式
	LogFileNameTimeFormat = "2006-01-02-15"

	// 日志文件大小限制
	LogMaxSize = 10

	// 日志文件的备份天数
	LogMaxBackups = 31

	// 日志文件的最大天数
	LogMaxAge = 31
)

var (
	// 日志
	Log zerolog.Logger

	// 缓存日志
	LumberLogger *lumberjack.Logger
)

// 初始化 zero log 日志  
func NewZeroLog(path, logFileNamePrefix string, stdoutFlag bool) zerolog.Logger {
	var logfilename string
	dataTimeStr := time.Now().Format(LogFileNameTimeFormat)
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
		// 	logfilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeStr + ".log"
		logfilename = logpath + "/pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeStr + ".log"

	} else {
		// 	logfilename = logpath + "/" + logFileNamePrefix + "-pid-" + strconv.Itoa(os.Getpid()) + "-" + dataTimeStr + ".log"
		logfilename = logpath + "/" + logFileNamePrefix + dataTimeStr + ".log"
	}

	LumberLogger = &lumberjack.Logger{
		Filename:   logfilename,
		MaxSize:    LogMaxSize, // megabytes
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge, // days
		Compress:   false,     // 开发时不压缩
	}

	wdiode := diode.NewWriter(LumberLogger, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})

	var writers []io.Writer

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if stdoutFlag {
		writers = []io.Writer{
			wdiode,
			os.Stdout,
		}
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		writers = []io.Writer{
			wdiode,
		}
	}

	multi := io.MultiWriter(writers...)
	// 	zerolog.TimeFieldFormat = time.RFC3339Nano //  "2006-01-02/15:04:05.999999999" //15:04:05.999999999
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999999999" // 15:04:05.999999999
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"
	zerolog.CallerFieldName = "c"
	zerolog.DurationFieldUnit = time.Nanosecond

	// 	vklog := zerolog.New(multi).With().Timestamp().Caller().Logger()
	vklog := zerolog.New(multi).With().Timestamp().Logger()
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
	// vklog.Info().Msg("----------------zerolog init success------------------")
	return vklog
}

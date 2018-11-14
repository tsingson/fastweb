// 本示例主要用于展示如何为一个logger实例配置一个或多个writer。

package multi

import (
	"fmt"
	"time"

	"github.com/tsingson/fastweb/zaputils"
)

// defaultLogger 使用vega/log包默认logger输出日志。
func defaultLogger() {
	// vega/log包中默认logger实例仅包含consoleWriter，
	// 开发者可以使用wzap.Info等方法使用该logger实例输出日志。
	zaputils.Debug("default log debug")
	zaputils.Info("default log info")
	zaputils.Warn("default log warn")
	zaputils.Error("default log error")

	// 构造使用fileWriter的logger实例，并覆盖vega/log包的默认logger后，
	// 开发者可以使用wzap.Info等方法，来使用覆盖后的logger输出日志。
	logger := zaputils.New(
		zaputils.WithLevel(zaputils.Info),
		zaputils.WithPath("./defaultLogger.log"),
		zaputils.WithFields(zaputils.Int("hahaah", 10), zaputils.String("dadsada", "fafasfa")),
	)
	// 使用SetDefaultLogger方法将指定logger实例注入到vega/log包中。
	// 使用wzap.Debug等方法会调用注入的logger实例输出日志。
	zaputils.SetDefaultLogger(logger)
	zaputils.Debug("debug")
	zaputils.Info("info", "name", 123)
	zaputils.Warn("warn")
	zaputils.Error("error")
}

// fileWriterLogger 仅使用fileWriter写日志。
func fileWriterLogger() {
	logger := zaputils.New(
		zaputils.WithLevelCombo("Warn | Error | Panic | Fatal"), // 只有级别为Warn或Error、Panic、Fatal日志才会被写入。
		zaputils.WithPath("./fileWriterLogger.log"),
	)
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}

// multiWriterLogger 使用多个writer(两个fileWriter、一个consoleWriter)同时写日志。
func multiWriterLogger() {
	zaputils.SetDefaultFields(zaputils.Int("appid", 100010), zaputils.String("appname", "test-go"))
	logger := zaputils.New(
		zaputils.WithOutput(
			zaputils.WithLevel(zaputils.Info),
			zaputils.WithPath("./multiWriterLogger1.log"),
		),
		zaputils.WithOutput(
			zaputils.WithLevelCombo("Warn | Error | Panic | Fatal"),
			zaputils.WithPath("./multiWriterLogger2.log"),
		),
		zaputils.WithOutput(
			zaputils.WithLevelMask(zaputils.DebugLevel),
			zaputils.WithColorful(true),
		),
		zaputils.WithLevelMask(zaputils.InfoLevel|zaputils.WarnLevel|zaputils.FatalLevel|zaputils.ErrorLevel),
		zaputils.WithPath("./multiWriterLogger.log"),
	)
	logger.Debug("debug", "aaa", 123, "bbb", 1234)
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}

func main() {
	fmt.Println(zaputils.DebugLevel, zaputils.InfoLevel, zaputils.WarnLevel, zaputils.ErrorLevel, zaputils.FatalLevel)
	multiWriterLogger()
	fileWriterLogger()
	defaultLogger()

	time.Sleep(time.Second)
}

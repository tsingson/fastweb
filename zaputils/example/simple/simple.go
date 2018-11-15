// 本示例主要用于展示如何为一个logger实例配置一个或多个writer。

package simple

import (
	"time"

	"github.com/tsingson/fastx/zaputils"
)

func main() {
	// 使用SetDefaultFields方法新增全局默认filed，所有logger实例记录日志时都将包含该field。
	zaputils.SetDefaultFields(
		zaputils.String("iid", "test12313131231451"),
	)
	zaputils.SetDefaultDir("./log/")
	logger := zaputils.New(
		zaputils.WithLevel(zaputils.Info),
		zaputils.WithPath("simple.log"),
		// 使用WithFields方法对指定logger实例新增默认filed，该logger实例记录日志时都将包含该field。
		zaputils.WithFields(zaputils.Int("key1", 10), zaputils.String("dadsada", "fafasfa")),
		zaputils.WithOutput(
			zaputils.WithLevelMask(zaputils.DebugLevel),
			zaputils.WithColorful(true),
			zaputils.WithFields(zaputils.Int("key2", 10), zaputils.String("dadsada", "fafasfa")),
		),
	)
	zaputils.SetDefaultLogger(logger)
	zaputils.Debug("debug")
	zaputils.Info("info")
	zaputils.Warn("warn")
	zaputils.Error("error")

	time.Sleep(time.Second)
}

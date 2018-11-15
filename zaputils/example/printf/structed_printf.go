// 本示例主要用于展示如何为使用structured、printf两种风格来写日志。

package printf

import (
	"github.com/tsingson/fastx/zaputils"
)

// AsyncWriteFile 展示了如何使用structured风格写入日志（默认异步写入）。
func AsyncWriteFile() {
	logger := zaputils.New(
		zaputils.WithPath("./async.log"),   // 日志写入文件为/tmp/async.log
		zaputils.WithLevel(zaputils.Error), // 日志写入级别为Error，此时只有级别大于等于Error级别日志才会被写入
	)
	// 写入Info级别日志（由于本级别小小于Error级别，此日志不会被写入）
	logger.Info("some information about LiLei", // 日志主题
		"name", "LiLei", // 字段名:name; 字段值:LiLei
		"age", 17, // 字段名:age; 字段值:17
		"sex", "male", // 字段名:sex; 字段值:"male"
	)
	// 写入Error级别日志（此级别大于等于Error级别，会被写入）
	logger.Error("some information about Hanmeimei",
		"name", "Hanmeimei", // 字段名:name; 字段值:Hanmeimei
		"age", 17, // 字段名:age; 字段值:17
		"sex", "female", // 字段名:sex; 字段值:"female"
	)
	// 写入Panic级别日志（此级别大于Error级别，会被写入，同时panic）
	logger.Panic("some information about LiLei", // 日志主题
		"name", "LiLei", // 字段名:name; 字段值:LiLei
		"age", 17, // 字段名:age; 字段值:17
		"sex", "male", // 字段名:sex; 字段值:"male"
	)
	// 写入Fatal级别日志（此级别大于Error级别，会被写入，且应用会执行os.Exit()退出）
	logger.Fatal("some information about Hanmeimei",
		"name", "Hanmeimei", // 字段名:name; 字段值:Hanmeimei
		"age", 17, // 字段名:age; 字段值:17
		"sex", "female", // 字段名:sex; 字段值:"female"
	)
}

// SyncWriteFile 展示了如何使用Structured风格同步写日志文件。
func SyncWriteFile() {
	logger := zaputils.New(
		zaputils.WithPath("./sync.log"),
		zaputils.WithLevel(zaputils.Error),
	)
	defer logger.Sync()

	logger.Error("some information about Hanmeimei",
		"name", "Hanmeimei", // 字段名:name; 字段值:Hanmeimei
		"age", 17, // 字段名:age; 字段值:17
		"sex", "female", // 字段名:sex; 字段值:"female"
	)
}

// PrintfWriteFile 展示了如何使用printf风格写入日志。
// 注意：printf风格写入日志性能低于structed风格，且不利于Kibana日志查询，请谨慎使用。
func PrintfWriteFile() {
	logger := zaputils.New(
		zaputils.WithPath("./sync.log"),
		zaputils.WithLevel(zaputils.Error),
	)
	defer logger.Sync()

	logger.Errorf("sync write %s", "How are you? I'm fine, thank you.")
}

func main() {
	PrintfWriteFile()
	SyncWriteFile()
	AsyncWriteFile()
}

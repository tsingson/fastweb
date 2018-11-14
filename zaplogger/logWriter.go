package zaplogger

import (
	"bytes"
	"log"

	"go.uber.org/zap"
)

// FormatStdLog set the output of stand package zlog to zaplog
func FormatStdLog() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(&logWriter{NewNoCallerLogger(false)})
}

type logWriter struct {
	logw *zap.Logger
}

// Write implement io.Writer, as std zlog's output
func (w logWriter) Write(p []byte) (n int, err error) {
	i := bytes.Index(p, []byte(":")) + 1
	j := bytes.Index(p[i:], []byte(":")) + 1 + i
	caller := bytes.TrimRight(p[:j], ":")
	// find last index of /
	i = bytes.LastIndex(caller, []byte("/"))
	// find penultimate index of /
	i = bytes.LastIndex(caller[:i], []byte("/"))
	w.logw.Info("stdLog", zap.ByteString("caller", caller[i+1:]), zap.ByteString("zlog", bytes.TrimSpace(p[j:])))
	n = len(p)
	err = nil
	return
}


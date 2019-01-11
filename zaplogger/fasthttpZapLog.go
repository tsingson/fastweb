package zaplogger

import (
	"fmt"
	"time"

	"github.com/tsingson/phi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/valyala/fasthttp"
)

// ZapLogger is a logger which compatible to logrus/std zlog/prometheus.
// it implements Print() Println() Printf() Dbug() Debugln() Debugf() Info() Infoln() Infof() Warn() Warnln() Warnf()
// Error() Errorln() Errorf() Fatal() Fataln() Fatalf() Panic() Panicln() Panicf() With() WithField() WithFields()
type ZapLogger struct {
	zlog *zap.Logger
}

// InitZaplogger
func InitZapLogger(log *zap.Logger) *ZapLogger {
	return &ZapLogger{
		log,
	}
}

// NewZapLogger return ZapLogger with caller field
func NewZapLogger(debugLevel bool) *ZapLogger {
	return &ZapLogger{NewLogger(debugLevel).WithOptions(zap.AddCallerSkip(1))}
}

// Printf logs a message at level Info on the ZapLogger.
func (l *ZapLogger) Printf(format string, args ...interface{}) {
	l.zlog.Info(fmt.Sprintf(format, args...))
}

// FastHttpZapLogHandler
// middle-ware for fasthttp
func (l *ZapLogger) FastHttpZapLogHandler(next phi.HandlerFunc) phi.HandlerFunc {
	return func(ctx *fasthttp.RequestCtx) {
		startTime := time.Now()
		next(ctx)

		var addrField zapcore.Field
		xRealIp := ctx.Request.Header.Peek("X-Real-IP")
		if len(xRealIp) > 0 {
			addrField = zap.ByteString("addr", ctx.Request.Header.Peek("X-Real-IP"))
		} else {
			addrField = zap.String("addr", ctx.RemoteAddr().String())
		}

		if ctx.Response.StatusCode() < 400 {
			l.zlog.Info("access",
				zap.Int("code", ctx.Response.StatusCode()),
				zap.Duration("time", time.Since(startTime)),
				zap.ByteString("method", ctx.Method()),
				zap.ByteString("path", ctx.Path()),
				zap.ByteString("agent", ctx.UserAgent()),
				zap.ByteString("req", ctx.RequestURI()),
				addrField)
		} else {
			l.zlog.Warn("access",
				zap.Int("code", ctx.Response.StatusCode()),
				zap.Duration("time", time.Since(startTime)),
				zap.ByteString("method", ctx.Method()),
				zap.ByteString("path", ctx.Path()),
				zap.ByteString("agent", ctx.UserAgent()),
				zap.ByteString("req", ctx.RequestURI()),
				addrField)
		}
	}
}

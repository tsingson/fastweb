package zaplogger

import (
	"time"

	"github.com/tsingson/phi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/valyala/fasthttp"
)

// FastHttpZapLogHandler
// middle-ware for fasthttp
func (l *ZapLogger) FastHttpZapLogHandler(next phi.RequestHandlerFunc) phi.RequestHandlerFunc {
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
			l.Log.Info("access",
				zap.Int("code", ctx.Response.StatusCode()),
				zap.Duration("time", time.Since(startTime)),
				zap.ByteString("method", ctx.Method()),
				zap.ByteString("path", ctx.Path()),
				zap.ByteString("agent", ctx.UserAgent()),
				zap.ByteString("req", ctx.RequestURI()),
				addrField)
		} else {
			l.Log.Warn("access",
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

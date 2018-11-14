package zaplogger

import (
	"time"

	"github.com/tsingson/phi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/valyala/fasthttp"
)



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
				addrField)
		} else {
			l.zlog.Warn("access",
				zap.Int("code", ctx.Response.StatusCode()),
				zap.Duration("time", time.Since(startTime)),
				zap.ByteString("method", ctx.Method()),
				zap.ByteString("path", ctx.Path()), addrField)
		}
		/**
		output.Printf("[%v] %v | %s | %s %s - %v - %v | %s",
			end.Format("2006/01/02 - 15:04:05"),
			ctx.RemoteAddr(),
			getHttp(ctx),
			ctx.Method(),
			ctx.RequestURI(),
			ctx.Response.Header.StatusCode(),
			end.Sub(begin),
			ctx.UserAgent(),
		)
		*/
	}
}


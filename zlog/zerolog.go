package zlog

import (
	"time"

	"github.com/valyala/fasthttp"

	"github.com/rs/zerolog"
	"github.com/tsingson/phi"
)

type (
	// FastHTTPLoggerAdapter  Adapter for passing apex/log used as gramework Logger into fasthttp
	FastHTTPLoggerAdapter struct {
		zlog zerolog.Logger
		fasthttp.Logger
	}
)

func FastHttpZeroLogHandler(next phi.RequestHandlerFunc) phi.RequestHandlerFunc {
	return func(ctx *fasthttp.RequestCtx) {
		begin := time.Now()
		next(ctx)
		end := time.Now()
		Log.Info().
			// 	Uint64("Connect Num", ctx.ConnRequestNum()).
			// 	Str("remoteAddr", ctx.RemoteAddr().String()).
			Str("remoteIp", ctx.RemoteIP().String()).
			Str("agent", string(ctx.UserAgent())).
			Str("method", string(ctx.Method())).
			Str("url", string(ctx.RequestURI())).
			Int("status", ctx.Response.Header.StatusCode()).
			Time("start", begin).
			Dur("duration", end.Sub(begin)).
			// 	Time("end", end).
			Msg("fasthttp")
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

// NewFastHTTPLoggerAdapter create new *FastHTTPLoggerAdapter
func NewFastHTTPLoggerAdapter(log zerolog.Logger) *FastHTTPLoggerAdapter {

	return &FastHTTPLoggerAdapter{
		zlog: log,
	}

}

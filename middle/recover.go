//  recover handler for phi mux
package middle

import (
	"github.com/tsingson/fastx/utils"
	"github.com/tsingson/phi"
	"github.com/valyala/fasthttp"
	"runtime/debug"
)

//  Recoverer(next phi.HandlerFunc) phi.HandlerFunc
func Recoverer(next phi.RequestHandlerFunc) phi.RequestHandlerFunc {
	fn := func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if rvr := recover(); rvr != nil {
				/**
				logEntry := GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
					debug.PrintStack()
				}
				*/
				Log.Error().Str("stack", utils.BytesToStringUnsafe(debug.Stack()))

				ctx.Error(utils.BytesToStringUnsafe(debug.Stack()), 500)
			}
		}()
		next.Handler(ctx)
	}
	return phi.RequestHandlerFunc(fn)
}

// design and code by tsingson

package main

import (
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"strings"
	"time"
)

// basicAuth returns the username and password provided in the request's
// Authorization header, if the request uses HTTP Basic Authentication.
// See RFC 2617, Section 2.
func basicAuth(ctx *fasthttp.RequestCtx) (username, password string, ok bool) {
	auth := ctx.Request.Header.Peek("Authorization")
	if auth == nil {
		return
	}
	return parseBasicAuth(string(auth))
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ("Aladdin", "open sesame", true).
func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}

// AuthRequestHandler is the basic auth-fasthttp handler
func AuthRequestHandler(next fasthttp.RequestHandler, requiredUser, requiredPassword string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := basicAuth(ctx)

		if hasAuth && user == requiredUser && password == requiredPassword {
			// Delegate request to the given handle
			next(ctx)
			return
		}
		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
		ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
	})
}

func LogRequestHandler(handler fasthttp.RequestHandler, log *zap.SugaredLogger) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		start := time.Now()
		// Delegate request to the given handle
		handler(ctx)
		end := time.Now()
		latency := end.Sub(start)
		/**
		log.Infow(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("time", end.Format(timeFormat)),
			zap.Duration("latency", latency),
		)
		*/
		log.Infow("Failed   URL: %s", url, "failed to fetch URL",
			// Structured context as loosely typed key-value pairs.
			"url", url,
			"attempt", 3,
			"backoff", time.Second,
			"latency", latency,
		)
		log.Sync()

		// end
	})
}

// Protected is the Protected handler
func Protected(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Protected!\n")
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set("Last-Modified", "Thu, 18 Jun 2015 10:24:27 GMT")

	ctx.Response.Header.Set("Accept-Ranges", "bytes")

	ctx.Response.Header.Set("E-Tag", "55829c5b-17")

	ctx.Response.Header.Set("Server", "golang-http-server")

	fmt.Fprintf(ctx, "<h1>\nHello world!\n</h1>\n")

}

// design and code by tsingson

// MultiParams is the multi params handler
func MultiParams(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ctx.UserValue("name"), ctx.UserValue("word"))
}

// QueryArgs is used for uri query args test #11:
// if the req uri is /ping?name=foo, output: Pong! foo
// if the req uri is /piNg?name=foo, redirect to /ping, output: Pong!
func QueryArgs(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	fmt.Fprintf(ctx, "Pong! %s\n", string(name))
}

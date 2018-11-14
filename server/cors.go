package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func Cors(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	})
}

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "some-api")
}

func example() {
	var handler fasthttp.RequestHandler
	handler = Index
	handlerMiddleWare := Cors(handler)

	//fasthttp.RequestHeader{}

	fastrouter := fasthttprouter.New()
	fastrouter.GET("/", handlerMiddleWare)
}

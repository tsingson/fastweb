package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/davecgh/go-spew/spew"
	"github.com/json-iterator/go"
	"github.com/kavu/go_reuseport"
	"github.com/koding/multiconfig"
	"github.com/valyala/fasthttp"
	//	"github.com/valyala/fasthttp/expvarhandler"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"runtime"
	"time"
)

type (
	ZapLog      = zap.SugaredLogger
	FastHandler fasthttp.RequestHandler
)

var (
	zaplog          *zap.SugaredLogger
	listener        net.Listener
	slistener       net.Listener
	err             error
	Configuration   *Config
	url             string = "test localhost"
	log_encoder_cfg        = zapcore.EncoderConfig{
		/**
		TimeKey:       "T",

		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",
		*/
		TimeKey:       "eventTime",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",

		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     time_encoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

func time_encoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
}

func ZapProductionConfig(Configuration *Config) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		//Encoding: "console",
		//EncoderConfig:   zap.NewProductionEncoderConfig(),
		EncoderConfig:    log_encoder_cfg,
		OutputPaths:      []string{"stderr", Configuration.Server.AccessFile},
		ErrorOutputPaths: []string{"stderr", Configuration.Server.ErrorFile},
	}
}
func init() {
	m := multiconfig.NewWithPath("config-fasthttp.toml") // supports TOML and JSON

	// Get an empty struct for your configuration
	Configuration = new(Config)

	// Populated the serverConf struct
	err := m.Load(Configuration) // Check for error
	m.MustLoad(Configuration)    // Panic's if there is any error
	if err != nil {
		zaplog.Errorw("config file tome error ")
	} else {
		spew.Dump(Configuration)
	}

	logger, _ := ZapProductionConfig(Configuration).Build()

	defer logger.Sync() // flushes buffer, if any
	zaplog = logger.Sugar()
	zaplog.Sync()

}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fastrouter := fasthttprouter.New()
	fastrouter.NotFound = fsHandler(Configuration.Server.Dir, 0)
	/**
	fastrouter.GET("/protected/", AuthRequestHandler(Protected, user, pass))
	fastrouter.GET("/exovar/", expvarhandler.ExpvarHandler)
	fastrouter.GET("/promethus/", prometheusHandler)

	fastrouter.POST("/post/:name", handlerHeader)

	fastrouter.GET("/set", setHandler)
	fastrouter.GET("/get", getHandler)
	fastrouter.GET("/clear", clearHandler)
	fastrouter.GET("/destroy", destroyHandler)
	*/
	if len(Configuration.Server.Addr) > 0 {
		listener, err = reuseport.Listen("tcp", Configuration.Server.Addr)
		if err != nil {
			panic(err)
		}
		go func() {
			if err := fasthttp.Serve(listener, fastrouter.Handler); err != nil {
				zaplog.Fatalf("error in ListenAndServe: %s", err)
			}
		}()
	}
	if len(Configuration.Server.AddrTLS) > 0 {
		zaplog.Infof("Starting HTTPS server on %q", Configuration.Server.AddrTLS)
		slistener, err = reuseport.Listen("tcp", Configuration.Server.AddrTLS)
		if err != nil {
			panic(err)
		}
		go func() {
			if err := fasthttp.ServeTLS(slistener, Configuration.Server.CertFile, Configuration.Server.KeyFile, fastrouter.Handler); err != nil {
				zaplog.Fatalf("error in ListenAndServeTLS: %s", err)
			}
		}()
	}
	zaplog.Infof("Serving files from directory %q", Configuration.Server.Addr)
	zaplog.Infof("See stats at http://%s/stats", Configuration.Server.Addr)

	// Wait forever.
	select {}
}

func handlerHeader(ctx *fasthttp.RequestCtx) {
	start := time.Now()
	ctx.Response.Header.Set("Server", "baz")

	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
	postBody := ctx.Request.Body()
	type (
		JsonPost struct {
			Tsingson string `json:"tsingson"`
			Test     string `json:"test"`
		}
	)
	data := JsonPost{"test", "test"}
	//spew.Dump(data)
	jsoniter.Unmarshal(postBody, &data)
	v := map[string]string{"hello": "json"}
	m, _ := jsoniter.Marshal(v)
	fmt.Println("1**********************************")
	spew.Dump(m)
	fmt.Println("**********************************")
	fmt.Fprintf(ctx, "%s", postBody)
	fmt.Fprintf(ctx, "%s", data)
	spew.Dump(data)
	spew.Dump(postBody)
	end := time.Now()
	latency := end.Sub(start)
	//if utc {
	end = end.UTC()
	//}
	zaplog.Info(ctx.Path(),

		//	zap.String("method", ctx.Request.Method()),
		zap.String("path", string(ctx.Path())),
		//	zap.String("ip", c.ClientIP()),
		//	zap.String("user-agent", c.Request.UserAgent()),
		zap.String("time", end.Format(time.RFC3339)),
		zap.Duration("latency", latency),
	)
	return
}
func ctxHandler(ctx *fasthttp.RequestCtx) {
	// set some headers and status code first
	ctx.SetContentType("foo/bar")
	ctx.SetStatusCode(fasthttp.StatusOK)

	// then write the first part of body
	fmt.Fprintf(ctx, "this is the first part of body\n")

	// then set more headers
	ctx.Response.Header.Set("Foo-Bar", "baz")

	// then write more body
	fmt.Fprintf(ctx, "this is the second part of body\n")

	// then override already written body
	ctx.SetBody([]byte("this is completely new body contents"))

	// then update status code
	ctx.SetStatusCode(fasthttp.StatusNotFound)

	// basically, anything may be updated many times before
	// returning from RequestHandler.
	//
	// Unlike net/http fasthttp doesn't put response to the wire until
	// returning from RequestHandler.
}

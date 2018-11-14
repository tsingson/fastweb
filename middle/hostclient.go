package middle

/**
import (
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

var (
	ProxyAddr string
	Log       zerolog.Logger
)

func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response

	uri := ctx.RequestURI()
	nreq := fasthttp.AcquireRequest()
	req.CopyTo(nreq)

	nreq.URI().UpdateBytes(uri[4:])
	//prepareRequest(req)
	var hostClient *fasthttp.HostClient
	hostClient = &fasthttp.HostClient{
		IsTLS:    false,
		Addr:     ProxyAddr,
		MaxConns: fasthttp.DefaultMaxConnsPerHost * 4,
		//	MaxConnDuration: 4 * time.Hour,
	}
	if err := hostClient.Do(nreq, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
		Log.Error().Err(err)
		resp.Header.SetServer("EPG")
		ctx.SetStatusCode(500)
		ctx.SetBodyString("error")
		return
	}
	defer fasthttp.ReleaseRequest(req)
	//postprocessResponse(resp)
	//litter.Dump(utils.Byte2String(resp.Header.Header()))
	resp.Header.SetServer("EPG")
	//	resp.Header.Set("Content-Location", "tsingson") //utils.Byte2String(uri))
	//litter.Dump(utils.Byte2String(resp.Header.Header()))
	return
}

func PostProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	uri := ctx.RequestURI()
	req.SetRequestURIBytes(uri[4:])
	req.Header.Del("Connection")
	var proxyClient *fasthttp.HostClient
	proxyClient = &fasthttp.HostClient{
		IsTLS:    false,
		Addr:     ProxyAddr,
		MaxConns: fasthttp.DefaultMaxConnsPerHost * 4,
		//	MaxConnDuration: 4 * time.Hour,
	}
	if err := proxyClient.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
		Log.Error().Err(err)
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		ctx.SetBodyString("error hostclient")
		return
	}
	//	litter.Dump(utils.Byte2String(resp.Header.Header()))
	resp.Header.Del("Connection")
	ctx.SetStatusCode(200)
	ctx.SetBodyString("post OK")
	return
}

func prepareRequest(req *fasthttp.Request) {
	// do not proxy "Connection" header.
	req.Header.Del("Connection")
	// strip other unneeded headers.

	// alter other request params before sending them to upstream host
}

func postprocessResponse(resp *fasthttp.Response) {
	// do not proxy "Connection" header
	resp.Header.Del("Connection")

	//litter.Dump(utils.Byte2String(resp.Header.Header()))
	// strip other unneeded headers

	// alter other response data if needed
}
*/

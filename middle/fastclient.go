package middle

import (
	"fmt"
	"time"

	"github.com/json-iterator/go"
	"github.com/tsingson/fastx/utils"

	"github.com/valyala/fasthttp"
)

func FastGetJson(url string, timeOut time.Duration) (*fasthttp.Response, error) {
	// init http client
	client := &fasthttp.Client{}
	client.MaxIdleConnDuration = timeOut

	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	//  defer fasthttp.ReleaseRequest(request)
	// 	defer fasthttp.ReleaseResponse(response)

	// 	request.SetConnectionClose()
	request.SetRequestURI(url)
	// 	request.Header.SetContentType("application/json; charset=utf-8")
	// 	request.Header.Add("Accept", "application/json")

	err := client.DoTimeout(request, response, timeOut)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func FastGetWithSession(url string, timeOut time.Duration, sid string) (*fasthttp.Response, error) {
	// init http client
	client := new(fasthttp.Client)
	client.MaxIdleConnDuration = timeOut
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	request.SetRequestURI(url)
	request.Header.Add("Accept", "application/json")
	if len(sid) > 0 {
		request.Header.Add("Sid", sid)
	}

	err := client.DoTimeout(request, response, timeOut)

	if err != nil {
		return nil, err
	}
	// 	fmt.Println(string(response.Header.Header()))
	// 	fmt.Println(string(response.Body()))

	return response, nil
}

func FastPostJson(url string, body []byte, timeOut time.Duration) (*fasthttp.Response, error) {

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	// 	request.Header.Add("User-Agent", "Test-Agent")
	// request.Header.SetContentType("application/json; charset=utf-8")
	request.Header.SetContentType("application/json")
	// Accept: application/vnd.pgrst.object+json
	request.Header.Add("Accept", "application/json")
	// 	request.Header.Add("Accept", "application/vnd.pgrst.object+json")
	request.Header.SetMethod("POST")
	request.SetBody(body)

	responce := fasthttp.AcquireResponse()
	client := &fasthttp.Client{
		MaxIdleConnDuration: 5 * time.Second,
	}

	err := client.DoTimeout(request, responce, timeOut)

	if err != nil {
		return nil, err
	}
	return responce, nil
	// bodyBytes := resp.Body()
	// 	println(string(bodyBytes))

}

// req.Header.Add("User-Agent", "Test-Agent")
// req.Header.SetContentType("application/json; charset=utf-8")
// Accept: application/vnd.pgrst.object+json
// request.Header.Add("Accept", "application/vnd.pgrst.object+json")
// vnd.pgrst.object+json for postgrest only

func PostStruct(url string, body interface{}, timeOut time.Duration) (*fasthttp.Response, error) {

	fmt.Println("PostStruct Url: ", url)
	postBodyByte, err := jsoniter.Marshal(body)
	if err != nil {
		return nil, err
	}
	{ // debug only
		postString := utils.BytesToStringUnsafe(postBodyByte)
		fmt.Println("Payload of PostStruct:", postString)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.SetMethod("POST")
	req.SetBody(postBodyByte)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{
		MaxIdleConnDuration: 5 * time.Second,
	}

	err1 := client.DoTimeout(req, resp, timeOut)

	if err1 != nil {
		return nil, err1
	}
	return resp, nil
	// 	payloadByte := responce.Body()
	// return payloadByte

}

func HostClient() {
	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:8080",
	}

	// Fetch google page via local proxy.
	statusCode, body, err := c.Get(nil, "http://httpbin.org/ip")
	if err != nil {
		Log.Fatal().Err(err)
	}
	if statusCode != fasthttp.StatusOK {
		Log.Error().Int("httpStatus", statusCode)
	}
	useResponseBody(body)

	// Fetch foobar page via local proxy. Reuse body buffer.
	statusCode, body, err = c.Get(body, "http://httpbin.org/")
	if err != nil {
		Log.Fatal().Err(err)
	}
	if statusCode != fasthttp.StatusOK {
		Log.Error().Int("httpStatus", statusCode)
	}
	useResponseBody(body)
}

func useResponseBody(body []byte) {
	// Do something with body :)
	println(string(body))
}

// design and code by tsingson

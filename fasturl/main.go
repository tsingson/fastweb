package fasturl

import (
	//	"github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

func fastGet(url string, timeOut time.Duration) (*fasthttp.Response, error) {
	// init http client
	client := &fasthttp.Client{}
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	//defer fasthttp.ReleaseRequest(request)
	//	defer fasthttp.ReleaseResponse(response)

	//	request.SetConnectionClose()
	request.SetRequestURI(url)
	request.Header.Add("Accept", "application/json")

	err := client.DoTimeout(request, response, timeOut)

	if err != nil {
		return nil, err
	}
	//	fmt.Println(string(response.Header.Header()))
	//	fmt.Println(string(response.Body()))

	return response, nil
}

func fastPost(url string, body []byte, timeOut time.Duration) (*fasthttp.Response, error) {
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(url)
	//	request.Header.Add("User-Agent", "Test-Agent")
	request.Header.Add("Accept", "application/json")

	// GET http://127.0.0.1:61765 HTTP/1.1
	// User-Agent: fasthttp
	// User-Agent: Test-Agent

	request.Header.SetMethod("POST")
	request.SetBody(body)

	responce := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	err := client.DoTimeout(request, responce, timeOut)

	if err != nil {
		//println("Error:", err.Error())
		return nil, err
	}
	return responce, nil
	//bodyBytes := resp.Body()
	//	println(string(bodyBytes))

}

func hostClient() {
	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:8080",
	}

	// Fetch google page via local proxy.
	statusCode, body, err := c.Get(nil, "http://httpbin.org/ip")
	if err != nil {
		log.Fatalf("Error when loading google page through local proxy: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	useResponseBody(body)

	// Fetch foobar page via local proxy. Reuse body buffer.
	statusCode, body, err = c.Get(body, "http://httpbin.org/")
	if err != nil {
		log.Fatalf("Error when loading foobar page through local proxy: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	useResponseBody(body)
}

func useResponseBody(body []byte) {
	// Do something with body :)
	println(string(body))
}

// design and code by tsingson

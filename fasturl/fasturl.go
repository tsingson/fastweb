package fasturl

import (
	//	"github.com/json-iterator/go"
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

//
func FastGet(url string, timeOut time.Duration) (*fasthttp.Response, error) {
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

func FastPost(url string, body []byte, timeOut time.Duration) (*fasthttp.Response, error) {
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

//PostFile 上传文件
func PostFile(fieldname, filename, uri string) ([]byte, error) {
	fields := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: fieldname,
			Filename:  filename,
		},
	}
	return PostMultipartForm(fields, uri)
}

//MultipartFormField 保存文件或其他字段信息
type MultipartFormField struct {
	IsFile    bool
	Fieldname string
	Value     []byte
	Filename  string
}

//PostMultipartForm 上传文件或其他多个字段
func PostMultipartForm(fields []MultipartFormField, uri string) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, field := range fields {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.Fieldname, field.Filename)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			fh, e := os.Open(field.Filename)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%v", e)
				return
			}
			defer fh.Close()

			if _, err = io.Copy(fileWriter, fh); err != nil {
				return
			}
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.Fieldname)
			if e != nil {
				err = e
				return
			}
			valueReader := bytes.NewReader(field.Value)
			if _, err = io.Copy(partWriter, valueReader); err != nil {
				return
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	return
}

// design and code by tsingson

package fasturl

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"strings"
)

func test_main() {

	var client *http.Client
	var remoteURL string
	{
		// setup a mocked http client.
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := httputil.DumpRequest(r, true)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%s", b)
		}))
		defer ts.Close()
		client = ts.Client()
		remoteURL = ts.URL
	}

	// prepare the reader instances to encode
	values := map[string]io.Reader{
		"file":  mustOpen("main.go"), // lets assume its this file
		"other": strings.NewReader("hello world!"),
	}
	err := Upload(client, remoteURL, values)
	if err != nil {
		panic(err)
	}
}

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func uploadPipe(url, filename string) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()
		part, err := m.CreateFormFile("myFile", "foo.txt")
		if err != nil {
			return
		}
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()
	http.Post(url, m.FormDataContentType(), r)

	//
	// url := `http://httpbin.org/post?key=123`
	//
	// req := fasthttp.AcquireRequest()
	// resp := fasthttp.AcquireResponse()
	// defer func(){
	// 	// 用完需要释放资源
	// 	fasthttp.ReleaseResponse(resp)
	// 	fasthttp.ReleaseRequest(req)
	// }()
	//
	// // 默认是application/x-www-form-urlencoded
	// req.Header.SetContentType("application/json")
	// req.Header.SetMethod("POST")
	//
	// req.SetRequestURI(url)
	//
	// requestBody := []byte(`{"request":"test"}`)
	// req.SetBody(requestBody)
	//
	// if err := fasthttp.Do(req, resp); err != nil {
	// 	fmt.Println("请求失败:", err.Error())
	// 	return
	// }
	//
	// b := resp.Body()
	//
	// fmt.Println("result:\r\n", string(b))

}

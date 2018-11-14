package fasturl

import (
	"github.com/json-iterator/go"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var (
	httpbin_header HttpBinHeaders
)

func TestFastGet(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("test fasthttp client function GET ", t, func() {

		response, err := fastGet("https://httpbin.org/headers", 5*time.Second)

		//	var ipbody HttpBinHeaders //{"113.87.14.183"}

		if err != nil {

			//	So(ipbody.Host, ShouldEqual, "httpbin.org")
			So(response.Header.StatusCode(), ShouldEqual, "200")
		}
	})
}
func TestFastPost(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("test fasthttp client function GET ", t, func() {
		response, err := fastPost("http://httpbin.org/post", []byte("p=q"), 5*time.Second)

		var ipbody HttpbinPostBody //{"113.87.14.183"}

		if err != nil {
			body := response.Body()

			jsoniter.Unmarshal(body, &ipbody)
			So(ipbody.Data, ShouldEqual, "p=q")
			//So(response.Header.StatusCode(), ShouldEqual, "200")
		}
	})
}

// design and code by tsingson

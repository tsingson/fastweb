package fasturl

// type for http://httpbin.org
// for test only
type (
	HttpBinHeaders struct {
		Headers `json:"headers"`
	}

	Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
		Connection     string `json:"Connection"`
		Cookie         string `json:"Cookie"`
		Dnt            string `json:"Dnt"`
		Host           string `json:"Host"`
		Referer        string `json:"Referer"`
		UserAgent      string `json:"User-Agent"`
	}
	HttpbinPostBody struct {
		Args struct {
		} `json:"args"`
		Data  string `json:"data"`
		Files struct {
		} `json:"files"`
		Form struct {
		} `json:"form"`
		Headers struct {
			Accept         string `json:"Accept"`
			AcceptEncoding string `json:"Accept-Encoding"`
			CacheControl   string `json:"Cache-Control"`
			Connection     string `json:"Connection"`
			ContentLength  string `json:"Content-Length"`
			ContentType    string `json:"Content-Type"`
			Host           string `json:"Host"`
			PostmanToken   string `json:"Postman-Token"`
			UserAgent      string `json:"User-Agent"`
		} `json:"headers"`
		JSON   interface{} `json:"json"`
		Origin string      `json:"origin"`
		URL    string      `json:"url"`
	}

	IPbody struct {
		Origin string `json:"origin"`
	}
)

// design and code by tsingson

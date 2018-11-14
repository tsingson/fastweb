package main

type (
	ServerConfig struct {
		Addr               string `default:":8888"`
		AddrTLS            string `default:""`
		ByteRange          bool   `default:"false"`
		CertFile           string //`defaulltl:"./ssl-cert-snakeoil.pem"`
		Compress           bool   `default:"true"`
		Dir                string `default:"/Users/qinshen/git/g2cn/public"`
		GenerateIndexPages bool   `default:"false"`
		KeyFile            string //`default:"./ssl-cert-snakeoil.key"`
		Vhost              bool   `default:"false"`
		ErrorFile          string `default:"./error-fasthttp.log"`
		AccessFile         string `default:"./access-fasthttp.log"`
		RedisFlag          bool   `default:"false"`
	}
	Config struct {
		Server ServerConfig
	}
)

/**
type singleton struct {
}

var instance *singleton
var once = sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}


*/

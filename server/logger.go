package main

import (
//"github.com/json-iterator/go"
//"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
)

var (
	rawJSON = []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
)

/**
func init() {

	var cfg zap.Config
	if err := jsoniter.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")

}

*/

// design and code by tsingson

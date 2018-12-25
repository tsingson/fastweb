// Package ginzap provides a logging middleware to get
// https://github.com/uber-go/zap as logging library for
// https://github.com/gin-gonic/gin. It can be used as replacement for
// the internal logging middleware
// http://godoc.org/github.com/gin-gonic/gin#Logger.
//
// This package is heavily based on https://github.com/szuecs/gin-glog
//
// Example:
//    package main
//    import (
//        "flag
//        "time"
//        "github.com/golang/glog"
//        "github.com/akath19/gin-zap"
//        "github.com/gin-gonic/gin"
//    )
//    func main() {
//        flag.Parse()
//        router := gin.New()
// 	      logger := zap.NewProduction()
//        router.Use(ginzap.Logger(3 * time.Second, logger))
//        //..
//        router.Use(gin.Recovery())
// 		  logger.Info("Gin bootstrapped with Zap")
//        router.Run(":8080")
//    }
//
package ginzap

import (
	"time"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

// setupLogging setups the logger to use zap
func setupLogging(duration time.Duration, zap *zap.Logger) {
	go func() {
		for range time.Tick(duration) {
			zap.Sync()
		}
	}()
}

// ErrorLogger returns a gin handler func for errors
func ErrorLogger() gin.HandlerFunc {
	return ErrorLoggerT(gin.ErrorTypeAny)
}

// ErrorLoggerT returns a gin handler middleware with the given
// type gin.ErrorType.
func ErrorLoggerT(t gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if !c.Writer.Written() {
			json := c.Errors.ByType(t).JSON()
			if json != nil {
				c.JSON(-1, json)
			}
		}
	}
}

// Logger returns a gin handler func for all logging
func Logger(logger *zap.Logger) gin.HandlerFunc {
	// 	setupLogging(duration, logger)

	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		// statusColor := colorForStatus(statusCode)
		// methodColor := colorForMethod(method)
		path := c.Request.URL.Path

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				logger.Warn("[GIN]",
					// zap.String("statusColor", statusColor),
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					// zap.String("methodColor", methodColor),
					zap.String("method", method),
					zap.String("path", path),
					zap.String("error", c.Errors.String()),
				)
			}
		case statusCode >= 500:
			{
				logger.Error("[GIN]",
					// zap.String("statusColor", statusColor),
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					// zap.String("methodColor", methodColor),
					zap.String("method", method),
					zap.String("path", path),
					zap.String("error", c.Errors.String()),
				)
			}
		default:
			logger.Info("[GIN]",
				// zap.String("statusColor", statusColor),
				zap.Int("statusCode", statusCode),
				zap.String("latency", latency.String()),
				zap.String("clientIP", clientIP),
				// zap.String("methodColor", methodColor),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("error", c.Errors.String()),
			)
		}
	}
}

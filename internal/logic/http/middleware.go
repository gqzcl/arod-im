package http

import (
	"time"

	"github.com/golang/glog"
	"github.com/kataras/iris/v12"
)

func loggerHandler(c *iris.Context) {
	// Start timer
	start := time.Now()
	// Process request

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)
	glog.Infof("TIME:%d", latency/time.Millisecond)
}

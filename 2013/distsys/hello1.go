// +build OMIT

package main

import (
	"flag"

	"github.com/golang/glog"
)

func main() {
	flag.Set("logtostderr", "true")
	glog.Infof("hello, world")
}

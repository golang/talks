// +build OMIT

package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func dumpURL(u, path string) error {
	net.Dial()
	res, err := http.Get(u)
	if err != nil {
		return fmt.Errorf("get %v: %v", u, err)
	}
	defer res.Body.Close()

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %v: %v", path, err)
	}
	defer f.Close()

	_, err := io.Copy(f, res.Body)
}

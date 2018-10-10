// +build ignore,OMIT

package main

import (
	"bytes"
	"net/http"
	"time"

	"github.com/golang/groupcache"
)

func main() {
	// STARTINIT OMIT
	me := "http://10.0.0.1"
	peers := groupcache.NewHTTPPool(me)

	// Whenever peers change:
	peers.Set("http://10.0.0.1", "http://10.0.0.2", "http://10.0.0.3")
	// ENDINIT OMIT

	// STARTGROUP OMIT
	var thumbNails = groupcache.NewGroup("thumbnail", 64<<20, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			fileName := key
			dest.SetBytes(generateThumbnail(fileName))
			return nil
		}))
	// ENDGROUP OMIT

	var ctx groupcache.Context
	var w http.ResponseWriter
	var r *http.Request

	// STARTUSE OMIT
	var data []byte
	err := thumbNails.Get(ctx, "big-file.jpg",
		groupcache.AllocatingByteSliceSink(&data))
	// ...
	_ = err               // OMIT
	var modTime time.Time // OMIT
	http.ServeContent(w, r, "big-file-thumb.jpg", modTime, bytes.NewReader(data))
	// ENDUSE OMIT
}

func generateThumbnail(filename string) []byte {
	// ...
	return nil
}

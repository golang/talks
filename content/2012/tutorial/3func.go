// +build OMIT

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	items, err := Get("golang") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items { // HL
		fmt.Println(item.Title)
	}
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

type Item struct {
	Title string
	URL   string
}

func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit) // HLurl
	resp, err := http.Get(url)                                // HLget
	if err != nil {
		return nil, err // HLreturn
	}
	defer resp.Body.Close()               // HLclose
	if resp.StatusCode != http.StatusOK { // HLstatus
		return nil, errors.New(resp.Status) // HLerrors
	}
	r := new(Response)                         // HLdecode
	err = json.NewDecoder(resp.Body).Decode(r) // HLdecode
	if err != nil {
		return nil, err // HLreturn
	}
	items := make([]Item, len(r.Data.Children)) // HLprepare
	for i, child := range r.Data.Children {     // HLconvert
		items[i] = child.Data // HLconvert
	} // HLconvert
	return items, nil // HLreturn
}

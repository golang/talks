// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Portland")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var w struct {
		Weather []struct {
			Desc string `json:"description"`
		} `json:"weather"`
	}
	if err := json.NewDecoder(res.Body).Decode(&w); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("No need to rush outside, we have %v.", w.Weather[0].Desc)
}

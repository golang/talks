package main

import (
	"fmt"
	"net/http"
)

type Applicant = http.Client

func main() {
	fmt.Printf("%T", Applicant{})
}

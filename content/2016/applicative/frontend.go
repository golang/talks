// +build ignore,OMIT

// The server program issues Google search requests. It serves on port 8080.
//
// The /search endpoint accepts these query params:
//   q=the Google search query
//
// For example, http://localhost:8080/search?q=golang serves the first
// few Google search results for "golang".
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/talks/content/2016/applicative/google"
)

func main() {
	http.HandleFunc("/search", handleSearch) // HL
	fmt.Println("serving on http://localhost:8080/search")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handleSearch handles URLs like "/search?q=golang" by running a
// Google search for "golang" and writing the results as HTML to w.
// The query parameter "output" selects alternate output formats:
// "json" for JSON, "prettyjson" for human-readable JSON.
func handleSearch(w http.ResponseWriter, req *http.Request) { // HL
	log.Println("serving", req.URL)

	// Check the search query.
	query := req.FormValue("q") // HL
	if query == "" {
		http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
		return
	}
	// ENDQUERY OMIT

	// Run the Google search.
	start := time.Now()
	results, err := google.Search(query) // HL
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ENDSEARCH OMIT

	// Create the structured response.
	type response struct {
		Results []google.Result
		Elapsed time.Duration
	}
	resp := response{results, elapsed} // HL
	// ENDRESPONSE OMIT

	// Render the response.
	switch req.FormValue("output") {
	case "json":
		err = json.NewEncoder(w).Encode(resp) // HL
	case "prettyjson":
		var b []byte
		b, err = json.MarshalIndent(resp, "", "  ") // HL
		if err == nil {
			_, err = w.Write(b)
		}
	default: // HTML
		err = responseTemplate.Execute(w, resp) // HL
	}
	// ENDRENDER OMIT
	if err != nil {
		log.Print(err)
		return
	}
}

var responseTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}</p>
</body>
</html>
`))

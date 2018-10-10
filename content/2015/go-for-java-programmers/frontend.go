// +build OMIT

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
	"net/url"
	"time"
)

func main() {
	http.HandleFunc("/search", handleSearch) // HL
	fmt.Println("serving on http://localhost:8080/search")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handleSearch handles URLs like "/search?q=golang" by running a
// Google search for "golang" and writing the results as HTML to w.
func handleSearch(w http.ResponseWriter, req *http.Request) {
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
	results, err := Search(query) // HL
	elapsed := time.Since(start)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ENDSEARCH OMIT

	// Render the results.
	type templateData struct {
		Results []Result
		Elapsed time.Duration
	}
	if err := resultsTemplate.Execute(w, templateData{ // HL
		Results: results,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
	// ENDRENDER OMIT
}

// A Result contains the title and URL of a search result.
type Result struct { // HL
	Title, URL string // HL
} // HL

var resultsTemplate = template.Must(template.New("results").Parse(`
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

// Search sends query to Google search and returns the results.
func Search(query string) ([]Result, error) {
	// Prepare the Google Search API request.
	u, err := url.Parse("https://ajax.googleapis.com/ajax/services/search/web?v=1.0")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("q", query) // HL
	u.RawQuery = q.Encode()

	// Issue the HTTP request and handle the response.
	resp, err := http.Get(u.String()) // HL
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // HL

	// Parse the JSON search result.
	// https://developers.google.com/web-search/docs/#fonje
	var jsonResponse struct {
		ResponseData struct {
			Results []struct {
				TitleNoFormatting, URL string
			}
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil { // HL
		return nil, err
	}

	// Extract the Results from jsonResponse and return them.
	var results []Result
	for _, r := range jsonResponse.ResponseData.Results { // HL
		results = append(results, Result{Title: r.TitleNoFormatting, URL: r.URL})
	}
	return results, nil
}

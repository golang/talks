package google

// START1 OMIT
type Result struct {
	Title, URL string
}

func Search(query string) ([]Result, error) {
	results := []Result{
		Web(query),
		Image(query),
		Video(query),
	}
	return results, nil
}

// STOP1 OMIT

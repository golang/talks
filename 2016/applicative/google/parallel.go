package google

func SearchParallel(query string) ([]Result, error) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	return []Result{<-c, <-c, <-c}, nil
}

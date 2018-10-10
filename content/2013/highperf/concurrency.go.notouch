// +build ignore

package pkg

func serial() {
	var lists []List
	var items []Item
	_, err := datastore.NewQuery("List").GetAll(c, &lists)
	if err != nil { /* ... */ }
	_, err := datastore.NewQuery("Item").GetAll(c, &items)
	if err != nil { /* ... */ }
	// write response
}

func parallel() {
	var lists []List
	var items []Item
	errc := make(chan error)	// HL
	go func() {	// HL
		_, err := datastore.NewQuery("List").GetAll(c, &lists)
		errc <- err
	}()	// HL
	go func() {	// HL
		_, err := datastore.NewQuery("Item").GetAll(c, &items)
		errc <- err
	}()	// HL
	err1, err2 := <-errc, <-errc	// HL
	if err1 != nil || err2 != nil { /* ... */ }
	// write response
}

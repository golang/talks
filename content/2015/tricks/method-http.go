// +build ignore

package main

import "net/http"

func main() {}

type Server struct {
	// Server state.
}

func (s *Server) index(w http.ResponseWriter, r *http.Request)  { /* Implementation. */ }
func (s *Server) edit(w http.ResponseWriter, r *http.Request)   { /* Implementation. */ }
func (s *Server) delete(w http.ResponseWriter, r *http.Request) { /* Implementation. */ }

func (s *Server) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", s.index)
	mux.HandleFunc("/edit/", s.edit)
	mux.HandleFunc("/delete/", s.delete)
}

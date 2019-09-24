package main

import "net/http"

func (s *server) routes() {
	http.HandleFunc("/books", s.handlerBooksIndex())
	http.HandleFunc("/books/show", s.handlerBooksShow())
	http.HandleFunc("/books/create", s.handlerBooksCreate())
	http.ListenAndServe(":3000", nil)
}

package main

func (s *server) routes() {
	s.router.HandleFunc("/books", s.handlerBooksIndex())
	s.router.HandleFunc("/books/show", s.handlerBooksShow())
	s.router.HandleFunc("/books/create", s.handlerBooksCreate())
}


package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/estensen/library/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type server struct {
	db     *models.DB
	router *mux.Router
}

func main() {
	db, err := models.NewDB("root@/library")
	if err != nil {
		log.Panic(err)
	}

	router := mux.NewRouter()

	srv := server{
		db:     db,
		router: router,
	}

	srv.routes()
	http.ListenAndServe(":3000", srv.router)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handlerBooksIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Printf("%s on %s is not allowed\n", r.Method, "/books")
			http.Error(w, http.StatusText(405), 405)
			return
		}

		books, err := s.db.AllBooks()
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		respondWithJSON(w, http.StatusOK, books)
	}
}

func (s *server) handlerBooksShow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Printf("%s on %s is not allowed\n", r.Method, "/books/show")
			http.Error(w, http.StatusText(405), 405)
			return
		}

		isbn := r.FormValue("isbn")
		if isbn == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		book, err := s.db.GetBook(isbn)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("book with isbn %s was not found", isbn)
				http.NotFound(w, r)
				return
			} else {
				log.Println("could not get book", err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
		}

		respondWithJSON(w, http.StatusOK, book)
	}
}

func (s *server) handlerBooksCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			log.Printf("%s on %s is not allowed\n", r.Method, "/books/create")
			http.Error(w, http.StatusText(405), 405)
			return
		}

		isbn := r.FormValue("isbn")
		title := r.FormValue("title")
		author := r.FormValue("author")
		if isbn == "" || title == "" || author == "" {
			log.Println("cannot create a book without a isbn, title or author")
			http.Error(w, http.StatusText(400), 400)
			return
		}

		err := s.db.AddBook(isbn, title, author)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		log.Printf("%s could not be marshaled into JSON: %s", data, err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

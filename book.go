package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Isbn string
	Title string
	Author string
}

var db *sql.DB

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreate)
	http.ListenAndServe(":3000", nil)
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("%s on %s is not allowed\n", r.Method, "/books")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT isbn, title, author FROM books")
	if err != nil {
		log.Println("could not get books: ", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	books := make([]*Book, 0)
	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author)
		if err != nil {
			log.Println("could not scan books: ", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Println("got an error while looping books", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func booksShow(w http.ResponseWriter, r *http.Request) {
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

	row := db.QueryRow("SELECT isbn, title, author FROM books WHERE isbn = ?", isbn)

	book := new(Book)
	err := row.Scan(&book.Isbn, &book.Title, &book.Author)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		fmt.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func booksCreate(w http.ResponseWriter, r *http.Request) {
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

	_, err := db.Exec("INSERT INTO books VALUES(?, ?, ?)", isbn, title, author)
	if err != nil {
		log.Printf("a book could not be inserted into the db: %s", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
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

func init() {
	var err error
	db, err = sql.Open("mysql", "root@/library")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

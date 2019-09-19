package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	isbn string
	title string
	author string
}

func main() {
	db, err := sql.Open("mysql", "root@/library")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := make([]*Book, 0)
	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.isbn, &book.title, &book.author)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("title | author | isbn\n")
	for _, book := range books {
		fmt.Printf("%s, %s, %s\n", book.title, book.author, book.isbn)
	}
}

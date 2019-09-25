package models

import (
	"log"
)

type Book struct {
	Isbn   string
	Title  string
	Author string
}

func (db *DB) AllBooks() ([]*Book, error) {
	rows, err := db.Query("SELECT isbn, title, author FROM books")
	if err != nil {
		log.Println("could not get books: ", err)
		return nil, err
	}
	defer rows.Close()

	books := make([]*Book, 0)
	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author)
		if err != nil {
			log.Println("could not scan books: ", err)
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Println("got an error while looping books", err)
		return nil, err
	}

	return books, nil
}

func (db *DB) GetBook(isbn string) (*Book, error) {
	row := db.QueryRow("SELECT isbn, title, author FROM books WHERE isbn = ?", isbn)

	book := new(Book)
	err := row.Scan(&book.Isbn, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (db *DB) AddBook(isbn, title, author string) error {
	_, err := db.Exec("INSERT INTO books VALUES(?, ?, ?)", isbn, title, author)
	if err != nil {
		log.Printf("a book could not be inserted into the db: %s", err)
		return err
	}
	return nil
}

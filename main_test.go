package main

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/estensen/library/models"
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldGetBooks(t *testing.T) {
	is := is.New(t)

	db, mock, err := sqlmock.New()
	is.NoErr(err)
	defer db.Close()

	router := mux.NewRouter()

	srv := server{
		db:     &models.DB{db},
		router: router,
	}
	srv.routes()

	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"isbn", "title", "author"}).
		AddRow("9781503261969", "Emma", "Jayne Austen").
		AddRow("9781505255607", "The Time Machine", "H. G. Wells").
		AddRow("9781503379640", "The Prince", "Niccolò Machiavelli")

	mock.ExpectQuery("^SELECT (.+) FROM books$").WillReturnRows(rows)

	data := []*models.Book{
		{"9781503261969", "Emma", "Jayne Austen"},
		{"9781505255607", "The Time Machine", "H. G. Wells"},
		{"9781503379640", "The Prince", "Niccolò Machiavelli"},
	}

	req, err := http.NewRequest("GET", "http://localhost:3000/books", nil)
	is.NoErr(err)

	srv.ServeHTTP(w, req)
	is.Equal(w.Code, http.StatusOK)

	assertJSON(w.Body.Bytes(), data, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("got unfulfilled expectations: %s", err)
	}
}

func assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

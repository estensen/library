# library
Learn to use databases with Go

## Install
library requires Go and mysql

Create db:
```bash
$ bash create_db.sh
```

Populate db:
```bash
$ bash populate_db.sh
```

## Run
```bash
$ go run book.go
```

## Query

GET books
```bash
$ curl "localhost:3000/books"
[
 {
  "Isbn": "9781503261969",
  "Title": "Emma",
  "Author": "Jayne Austen"
 },
 {
  "Isbn": "9781505255607",
  "Title": "The Time Machine",
  "Author": "H. G. Wells"
 },
 {
  "Isbn": "9781503379640",
  "Title": "The Prince",
  "Author": "Niccolò Machiavelli"
 }
]
```

GET book by isbn
```bash
$ curl "localhost:3000/books/show?isbn=9781505255607"
{
 "Isbn": "9781505255607",
 "Title": "The Time Machine",
 "Author": "H. G. Wells"
}
```

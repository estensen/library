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
$ go run .
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

PUT book
```bash
$ curl -i -X PUT -d "isbn=9780553294385&title=I, Robot&author=Isaac Asimov" "http://localhost:3000/books/create"
HTTP/1.1 200 OK
Date: Fri, 20 Sep 2019 13:18:12 GMT
Content-Length: 0

```
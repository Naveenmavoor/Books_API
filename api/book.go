package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Book struct
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn" `
}

//ToJSON does the json marshaling function
func (b Book) ToJSON() []byte {
	tojson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return tojson
}

//FromJSON does the json unmarshaling
func fromJSON(data []byte) Book {
	var newbook Book
	err := json.Unmarshal(data, &newbook)
	if err != nil {
		panic(err)
	}
	return newbook
}

//Books is the List of all books
var Books = map[string]Book{
	"012345678": Book{Title: "Alchemist", Author: "PAULO COELO", ISBN: "012345678"},
	"012345679": Book{Title: "ATOMIC HABITS", Author: "JAMES CLEAR", ISBN: "012345679"},
}

//BooksHandleFunc function is used to handle Create and Read operation for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Hello error occured")
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := fromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

//BookHandleFunc function is used to handle Read, Update and Delete operation for a Book
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := fromJSON(body)
		isupdated := UpdateBook(isbn, book)
		if isupdated {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

//AllBooks gives all the available books
func AllBooks() []Book {
	var s []Book
	for _, v := range Books {
		s = append(s, v)
	}
	return s
}

func writeJSON(w http.ResponseWriter, books interface{}) {
	fmt.Fprint(w, books)
}

//CreateBook creates a book
func CreateBook(book Book) (string, bool) {
	Books[book.ISBN] = book
	_, found := Books[book.ISBN]
	fmt.Println(Books)
	return book.ISBN, found
}

//GetBook gets the book from the library
func GetBook(isbn string) (Book, bool) {
	book, found := Books[isbn]

	return book, found
}

//UpdateBook update a book
func UpdateBook(isbn string, book Book) bool {
	_, found := GetBook(isbn)
	if found {
		Books[isbn] = book
		return true
	}
	return false
}

//DeleteBook deletes the requested book
func DeleteBook(isbn string) {
	delete(Books, isbn)
}

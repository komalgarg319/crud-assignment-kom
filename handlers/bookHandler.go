package handlers

import (
	"crud-assignment/database"
	"crud-assignment/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	database.Instance.Find(&books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	bookId := mux.Vars(r)["id"]
	if checkIfBookExists(bookId) == false {
		json.NewEncoder(w).Encode("Book not Found!")
		return
	}
	var book model.Book
	database.Instance.First(&book, bookId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

	json.NewEncoder(w).Encode(GetAuthorById22(book.AuthID))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	database.Instance.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId := mux.Vars(r)["id"]
	if checkIfBookExists(bookId) == false {
		json.NewEncoder(w).Encode("Book Not Found!")
		return
	}
	var book model.Book
	database.Instance.First(&book, bookId)
	json.NewDecoder(r.Body).Decode(&book)
	//fmt.Println(book)
	database.Instance.Save(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := mux.Vars(r)["id"]
	if checkIfBookExists(bookId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Book Not Found!")
		return
	}
	var book model.Book
	database.Instance.Delete(&book, bookId)
	json.NewEncoder(w).Encode("Book Deleted Successfully!")
}

func checkIfBookExists(bookId string) bool {
	var book model.Book
	database.Instance.First(&book, bookId)
	if book.ID == 0 {
		return false
	}
	return true
}

package handlers

import (
	"crud-assignment/database"
	"crud-assignment/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author
	database.Instance.Find(&authors)
	database.Instance.Preload("Books").Find(&authors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authors)
}
func GetAuthorById(w http.ResponseWriter, r *http.Request) {

	authorId := mux.Vars(r)["id"]
	if checkIfAuthorExists(authorId) == false {
		json.NewEncoder(w).Encode("Author not Found!")
		return
	}
	var author model.Author
	database.Instance.First(&author, authorId)
	database.Instance.Preload("Books").Find(&author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func GetAuthorById22(authorId uint) string {

	var author model.Author
	database.Instance.First(&author, authorId)
	return author.Name
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author model.Author
	json.NewDecoder(r.Body).Decode(&author)
	database.Instance.Create(&author)
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := mux.Vars(r)["id"]
	if checkIfAuthorExists(authorId) == false {
		json.NewEncoder(w).Encode("Author Not Found!")
		return
	}
	var author model.Author
	database.Instance.First(&author, authorId)
	json.NewDecoder(r.Body).Decode(&author)
	//fmt.Println(book)
	database.Instance.Save(&author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authorId := mux.Vars(r)["id"]
	if checkIfAuthorExists(authorId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Author Not Found!")
		return
	}
	var author model.Author
	//database.Instance.Association("Books").Delete(&author, authorId)
	// var bk model.Book
	database.Instance.Preload("Books").First(&author, authorId)
	database.Instance.Delete(&author.Books)
	database.Instance.Delete(&author, authorId)
	//database.Instance.Delete(&author, authorId).Association("Books")
	json.NewEncoder(w).Encode("Author Deleted Successfully!")
}

func checkIfAuthorExists(authorId string) bool {
	var author model.Author
	database.Instance.First(&author, authorId)
	if author.ID == 0 {
		return false
	}
	return true
}

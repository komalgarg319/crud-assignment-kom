package main

import (
	"crud-assignment/database"
	"crud-assignment/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc("/api/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", handlers.DeleteBook).Methods("DELETE")

	//Author
	router.HandleFunc("/api/authors", handlers.CreateAuthor).Methods("POST")
	router.HandleFunc("/api/authors/{id}", handlers.GetAuthorById).Methods("GET")
	router.HandleFunc("/api/authors", handlers.GetAuthors).Methods("GET")
	router.HandleFunc("/api/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/api/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")
}
func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()
	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)
	// Register Routes
	RegisterProductRoutes(router)
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"library_assesment/db"
	"library_assesment/handlers"
	"log"
	"net/http"
)

func main() {
	DB := db.Connect()
	handleRequests(DB)
	db.CloseConnection(DB)
}

func handleRequests(DB *sql.DB) {
	requestHandler := handlers.New(DB)
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/books/{id}", requestHandler.GetBook).Methods(http.MethodGet)
	myRouter.HandleFunc("/booksRead/{id}", requestHandler.GetBooksRead).Methods(http.MethodGet)
	myRouter.HandleFunc("/members/{id}", requestHandler.GetMember).Methods(http.MethodGet)

	myRouter.HandleFunc("/borrow", requestHandler.BorrowBook).Methods(http.MethodPost)
	myRouter.HandleFunc("/return", requestHandler.ReturnBook).Methods(http.MethodPatch)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Library REST API!")
}

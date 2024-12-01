package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"library_assesment/models"
	"log"
	"net/http"
)

func (requestHandler Service) GetBook(writer http.ResponseWriter, request *http.Request) {
	errMessage := map[string]string{}
	vars := mux.Vars(request)
	id := vars["id"]

	queryStmt := `SELECT * FROM books WHERE id = $1 ;`
	results, err := requestHandler.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("SQL error", err)
		writer.WriteHeader(500)
		errMessage = map[string]string{"message": err.Error()}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}

	var book models.Book
	for results.Next() {
		err = results.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			log.Println("failed to scan", err)
			writer.WriteHeader(500)
			errMessage = map[string]string{"message": err.Error()}
			json.NewEncoder(writer).Encode(errMessage)
			return
		}
	}

	if book.Id == 0 {
		writer.WriteHeader(404)
		errMessage = map[string]string{"message": "Book not found"}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}
	log.Println("Results: ", book)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(book)
}

package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"library_assesment/models"
	"log"
	"net/http"
)

func (requestHandler Service) GetBooksRead(writer http.ResponseWriter, request *http.Request) {
	errMessage := map[string]string{}
	vars := mux.Vars(request)
	id := vars["id"]
	log.Println("Query books read by member with id: ", id)
	queryStmt := `SELECT DISTINCT(books.id), books.title, books.author FROM members
				LEFT JOIN loans on members.id = loans.member_id
				LEFT JOIN books on books.id = loans.book_id
				WHERE members.id = $1
				AND loans.borrowed IS NOT NULL;`

	results, err := requestHandler.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("SQL error", err)
		writer.WriteHeader(500)
		errMessage = map[string]string{"message": err.Error()}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}

	var booksRead = make([]models.Book, 0)

	for results.Next() {
		var book models.Book
		err = results.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			log.Println("failed to scan", err)
			writer.WriteHeader(500)
			errMessage = map[string]string{"message": err.Error()}
			json.NewEncoder(writer).Encode(errMessage)
			return
		}

		booksRead = append(booksRead, book)
	}
	log.Println("Results:", booksRead)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(booksRead)
}

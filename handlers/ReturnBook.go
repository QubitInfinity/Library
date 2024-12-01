package handlers

import (
	"encoding/json"
	"io/ioutil"
	"library_assesment/models"
	"log"
	"net/http"
	"time"
)

func (requestHandler Service) ReturnBook(writer http.ResponseWriter, request *http.Request) {
	errMessage := map[string]string{}
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println(err)
		writer.WriteHeader(500)
		errMessage = map[string]string{"message": "Failed parse request"}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}
	var loan models.Loans
	err = json.Unmarshal(body, &loan)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(400)
		errMessage = map[string]string{"message": "Failed to unmarshal json"}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}
	var today = time.Now()
	loan.ReturnedDate = &today

	queryStmt := `UPDATE loans SET returned = $1 
					WHERE loans.book_id = $2
					AND loans.member_id = $3
					AND loans.returned IS NULL 
					RETURNING Id;`
	err = requestHandler.DB.QueryRow(queryStmt, &loan.ReturnedDate, &loan.BookId, &loan.MemberId).Scan(&loan.Id)
	if err != nil {
		log.Println("SQL error", err, loan)
		if err.Error() == "sql: no rows in result set" {
			writer.WriteHeader(409)
			errMessage = map[string]string{"message": "Book already returned"}
			json.NewEncoder(writer).Encode(errMessage)
		} else {
			writer.WriteHeader(500)
			errMessage = map[string]string{"message": err.Error()}
			json.NewEncoder(writer).Encode(errMessage)
		}
		return
	}

	log.Println("Returned: ", loan)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(loan)
}

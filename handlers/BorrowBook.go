package handlers

import (
	"encoding/json"
	"io/ioutil"
	"library_assesment/models"
	"log"
	"net/http"
	"time"
)

func (requestHandler Service) BorrowBook(writer http.ResponseWriter, request *http.Request) {
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
	loan.BorrowedDate = &today

	queryStmt := `WITH existing_loan AS (
						SELECT id
						FROM loans
						WHERE book_id = $1
						AND returned IS NULL
						)
						INSERT INTO public.loans (book_id, member_id, borrowed)
						SELECT $2, $3, $4
						WHERE NOT EXISTS (SELECT 1 FROM existing_loan)
						RETURNING id;`
	err = requestHandler.DB.QueryRow(queryStmt, &loan.BookId, &loan.BookId, &loan.MemberId, &loan.BorrowedDate).Scan(&loan.Id)
	if err != nil {
		log.Println("SQL error", err, loan)
		if err.Error() == "sql: no rows in result set" {
			writer.WriteHeader(409)
			errMessage = map[string]string{"message": "Book not available"}
		} else {
			writer.WriteHeader(500)
			errMessage = map[string]string{"message": err.Error()}
		}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}

	log.Println("Borrowed: ", loan)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(loan)

}

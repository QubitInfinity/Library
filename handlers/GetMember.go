package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"library_assesment/models"
	"log"
	"net/http"
)

func (requestHandler Service) GetMember(writer http.ResponseWriter, request *http.Request) {
	errMessage := map[string]string{}
	vars := mux.Vars(request)
	id := vars["id"]

	queryStmt := `SELECT * FROM members WHERE id = $1 ;`
	results, err := requestHandler.DB.Query(queryStmt, id)
	if err != nil {
		log.Println("SQL error", err)
		writer.WriteHeader(500)
		errMessage = map[string]string{"message": err.Error()}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}

	var member models.Member
	for results.Next() {
		err = results.Scan(&member.Id, &member.Full_name)
		if err != nil {
			log.Println("failed to scan", err)
			writer.WriteHeader(500)
			errMessage = map[string]string{"message": err.Error()}
			json.NewEncoder(writer).Encode(errMessage)
			return
		}
	}

	if member.Id == 0 {
		writer.WriteHeader(404)
		errMessage = map[string]string{"message": "Member not found"}
		json.NewEncoder(writer).Encode(errMessage)
		return
	}

	log.Println("Results: ", member)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(member)
}

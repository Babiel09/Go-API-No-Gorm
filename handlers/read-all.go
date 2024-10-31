package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Babiel09/models"
)

func ListAll(w http.ResponseWriter, r *http.Request) {
	people, err := models.GetAll()
	if err != nil {
		log.Printf("Unxpected erro: %v", err)
	}
	w.Header().Add("Context-type", "application/json")
	json.NewEncoder(w).Encode(people)
}

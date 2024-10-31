package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Babiel09/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var people models.People
	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	id, err := models.InsertInformation(people)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Unspected error: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Insert people, id: %d", id),
		}
	}
	w.Header().Add("Context-type", "applicaton/json")
	json.NewEncoder(w).Encode(resp)
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Babiel09/models"
	"github.com/go-chi/chi"
)

func UpdateInformations(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var people models.People
	err = json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rows, err := models.Update(uint(id), people)
	if err != nil {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if rows > 1 {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "Item suceffuly update",
	}
	w.Header().Add("Context-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

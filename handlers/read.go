package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Babiel09/models"
	"github.com/go-chi/chi"
)

func Read(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Unxpected error: %v", err)
		return
	}

	people, err := models.Get(uint(id))
	if err != nil {
		log.Printf("Unxpected error: %v", err)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	json.NewEncoder(w).Encode(people)

}

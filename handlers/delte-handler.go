package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Babiel09/models"
	"github.com/go-chi/chi"
)

func Delte(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rows, err := models.Delete(uint(id))
	if err != nil {
		log.Printf("Unspected error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if rows > 1 {
		log.Printf("More them 1 row were deleted: %d", id)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Row suceffuly deleted",
	}

	w.Header().Add("Context-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

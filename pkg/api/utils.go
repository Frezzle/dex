package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON writes a JSON-marshallable value to the HTTP response-writer.
// If it errors then it'll respond with a generic HTTP error.
func JSON(w http.ResponseWriter, value any) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(value)
	if err != nil {
		http.Error(w, fmt.Errorf("failed writing JSON response: %w", err).Error(), http.StatusInternalServerError)
	}
}

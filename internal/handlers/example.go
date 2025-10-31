package handlers

import (
	"encoding/json"
	"net/http"
)

type Example struct {
	Message string `json:"message"`
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Example{Message: "hello from go-api-starter"})
}

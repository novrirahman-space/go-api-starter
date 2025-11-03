package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Probe struct {
	Status string `json:"status"`
	Time time.Time `json:"time"`
}

func Liveness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Probe{
		Status: "alive",
		Time:   time.Now(),
	})
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	// TODO: tambahkan cek dependecy (DB, cache, dll) bila ada
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Probe{
		Status: "ready",
		Time:   time.Now(),
	})
}
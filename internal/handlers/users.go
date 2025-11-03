package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = make(map[int]User)
	mu    sync.RWMutex
	next  = 1
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	out := make([]User, 0, len(users))
	for _, u := range users { out = append(out, u) }
	_ = json.NewEncoder(w).Encode(out)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var in struct{ Name string `json:"name"` }
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Name == "" {
		http.Error(w, "invalid payload", http.StatusBadRequest); return
	}
	mu.Lock()
	id := next; next++
	u := User{ID: id, Name: in.Name}
	users[id] = u
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(u)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	mu.Lock(); defer mu.Unlock()
	delete(users, id)
	w.WriteHeader(http.StatusNoContent)
}

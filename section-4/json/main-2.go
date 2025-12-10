package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{}

func (s *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		ID:   payload.ID,
		Name: payload.Name,
	}

	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.ID == "" || u.Name == "" {
		return errors.New("invalid user data")
	}
	for _, user := range users {
		if user.ID == u.ID {
			return errors.New("user already exists")
		}
	}

	users = append(users, u)
	return nil
}

func main() { //main
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	srv.ListenAndServe()
}

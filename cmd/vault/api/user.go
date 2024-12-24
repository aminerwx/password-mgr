package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aminerwx/password-mgr/cmd/vault/model"
)

// Handle user auth
func (s *Server) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, _ := io.ReadAll(r.Body)
	w.Header().Set("content-type", "application/json")

	err := json.Unmarshal(body, &user)
	if err != nil || user == (model.User{}) {
		w.WriteHeader(StatusBadRequestJSON.StatusCode)
		json.NewEncoder(w).Encode(StatusBadRequestJSON)
		return
	}

	u, _ := model.NewUser(user.Username, user.MasterPassword)
	err = s.store.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(StatusConflictJSON)
		return
	}
	w.WriteHeader(StatusCreatedJSON.StatusCode)
	json.NewEncoder(w).Encode(StatusCreatedJSON)
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, _ := io.ReadAll(r.Body)
	w.Header().Set("content-type", "application/json")

	err := json.Unmarshal(body, &user)
	if err != nil || user == (model.User{}) {
		w.WriteHeader(StatusBadRequestJSON.StatusCode)
		json.NewEncoder(w).Encode(StatusBadRequestJSON)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(StatusConflictJSON)
		return
	}
	w.WriteHeader(StatusCreatedJSON.StatusCode)
	json.NewEncoder(w).Encode(StatusCreatedJSON)
}

func (s *Server) RemoveUserHandler(w http.ResponseWriter, r *http.Request) {
}

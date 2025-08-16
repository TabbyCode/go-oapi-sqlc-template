package server

import (
	"net/http"

	"github.com/xurvan/go-oapi-sqlc-template/internal/database"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
	httputil2 "github.com/xurvan/go-oapi-sqlc-template/internal/httputil"
)

type Server struct {
	db *database.Database
}

func NewServer(db *database.Database) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.db.ListUsers(r.Context())
	if err != nil {
		httputil2.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputil2.WriteJSON(w, http.StatusOK, users)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCreate oapi.UserCreate
	if !httputil2.ReadJSON(w, r, &userCreate) {
		return
	}

	user, err := s.db.CreateUser(r.Context(), userCreate)
	if err != nil {
		httputil2.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputil2.WriteJSON(w, http.StatusCreated, user)
}

func (s *Server) GetUserById(w http.ResponseWriter, r *http.Request, id int64) {
	user, err := s.db.GetUser(r.Context(), id)
	if err != nil {
		httputil2.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	httputil2.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var userUpdate oapi.UserUpdate
	if !httputil2.ReadJSON(w, r, &userUpdate) {
		return
	}

	user, err := s.db.UpdateUser(r.Context(), id, userUpdate)
	if err != nil {
		httputil2.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputil2.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	err := s.db.DeleteUser(r.Context(), id)
	if err != nil {
		httputil2.WriteError(w, http.StatusInternalServerError, "Failed to delete user: "+err.Error())
		return
	}

	httputil2.WriteJSON(w, http.StatusNoContent, nil)
}

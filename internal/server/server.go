package server

import (
	"net/http"
	"time"

	"github.com/xurvan/go-template/internal/gen/oapi"
	"github.com/xurvan/go-template/pkg/httputil"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := []oapi.User{
		{
			Id:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			CreatedAt: time.Now(),
		},
		{
			Id:        2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			CreatedAt: time.Now(),
		},
	}

	httputil.WriteJSON(w, http.StatusOK, users)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCreate oapi.UserCreate
	if !httputil.ReadJSON(w, r, &userCreate) {
		return
	}

	// This is a mock implementation
	user := oapi.User{
		Id:        3,
		Name:      userCreate.Name,
		Email:     userCreate.Email,
		CreatedAt: time.Now(),
	}

	httputil.WriteJSON(w, http.StatusCreated, user)
}

func (s *Server) GetUserById(w http.ResponseWriter, r *http.Request, id int64) {
	user := oapi.User{
		Id:        id,
		Name:      "John Doe",
		Email:     "john@example.com",
		CreatedAt: time.Now(),
	}

	httputil.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	var userUpdate oapi.UserUpdate
	if !httputil.ReadJSON(w, r, &userUpdate) {
		return
	}

	now := time.Now()
	user := oapi.User{
		Id:        id,
		CreatedAt: time.Now(),
		UpdatedAt: &now,
	}

	if userUpdate.Name != nil {
		user.Name = *userUpdate.Name
	} else {
		user.Name = "Default Name"
	}

	if userUpdate.Email != nil {
		user.Email = *userUpdate.Email
	} else {
		user.Email = "default@example.com"
	}

	httputil.WriteJSON(w, http.StatusOK, user)
}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	httputil.WriteJSON(w, http.StatusNoContent, nil)
}

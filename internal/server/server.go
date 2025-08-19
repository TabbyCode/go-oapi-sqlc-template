package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/xurvan/go-oapi-sqlc-template/internal/database"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
)

type Server struct {
	db *database.Database
}

// Ensure that Server implements the StrictServerInterface interface at compile time
var _ oapi.StrictServerInterface = (*Server)(nil)

func NewServer(db *database.Database) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) ListUsers(ctx context.Context, request oapi.ListUsersRequestObject) (oapi.ListUsersResponseObject, error) {
	users, err := s.db.ListUsers(ctx)
	if err != nil {
		return oapi.ListUsers500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}
	return oapi.ListUsers200JSONResponse(users), nil
}

func (s *Server) CreateUser(ctx context.Context, request oapi.CreateUserRequestObject) (oapi.CreateUserResponseObject, error) {
	if request.Body == nil {
		return oapi.CreateUser400JSONResponse{
			BadRequestJSONResponse: oapi.BadRequestJSONResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
		}, nil
	}

	user, err := s.db.CreateUser(ctx, *request.Body)
	if err != nil {
		return oapi.CreateUser500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.CreateUser201JSONResponse(*user), nil
}

func (s *Server) GetUserById(ctx context.Context, request oapi.GetUserByIdRequestObject) (oapi.GetUserByIdResponseObject, error) {
	user, err := s.db.GetUser(ctx, request.Id)
	if err != nil {
		return oapi.GetUserById404JSONResponse{
			NotFoundJSONResponse: oapi.NotFoundJSONResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.GetUserById200JSONResponse(*user), nil
}

func (s *Server) UpdateUser(ctx context.Context, request oapi.UpdateUserRequestObject) (oapi.UpdateUserResponseObject, error) {
	if request.Body == nil {
		return oapi.UpdateUser400JSONResponse{
			BadRequestJSONResponse: oapi.BadRequestJSONResponse{
				Code:    http.StatusBadRequest,
				Message: "invalid request body",
			},
		}, nil
	}

	user, err := s.db.UpdateUser(ctx, request.Id, *request.Body)
	if err != nil {
		return oapi.UpdateUser500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		}, nil
	}

	return oapi.UpdateUser200JSONResponse(*user), nil
}

func (s *Server) DeleteUser(ctx context.Context, request oapi.DeleteUserRequestObject) (oapi.DeleteUserResponseObject, error) {
	err := s.db.DeleteUser(ctx, request.Id)
	if err != nil {
		if errors.Is(err, database.ErrUserNotFound) {
			return oapi.DeleteUser404JSONResponse{
				NotFoundJSONResponse: oapi.NotFoundJSONResponse{
					Code:    http.StatusNotFound,
					Message: "User not found",
				},
			}, nil
		}

		return oapi.DeleteUser500JSONResponse{
			ErrorJSONResponse: oapi.ErrorJSONResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to delete user: " + err.Error(),
			},
		}, nil
	}

	return oapi.DeleteUser204Response{}, nil
}

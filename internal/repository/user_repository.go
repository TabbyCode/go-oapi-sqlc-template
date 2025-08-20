package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jinzhu/copier"

	"github.com/xurvan/go-oapi-sqlc-template/internal/config"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/db"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
)

type UserRepository struct {
	qr *db.Queries
}

func NewUserRepository(cfg *config.Config) (*UserRepository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("database: %w", err)
	}

	postgres := UserRepository{
		qr: db.New(conn),
	}

	return &postgres, nil
}

func (d *UserRepository) Create(ctx context.Context, user oapi.UserCreate) (*oapi.User, error) {
	var params db.CreateUserParams

	err := copier.Copy(&params, user)
	if err != nil {
		return nil, err
	}

	row, err := d.qr.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	var res oapi.User

	err = copier.Copy(&res, row)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (d *UserRepository) Get(ctx context.Context, id int32) (*oapi.User, error) {
	row, err := d.qr.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	var user oapi.User

	err = copier.Copy(&user, row)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *UserRepository) List(ctx context.Context, user oapi.ListUsersParams) ([]oapi.User, error) {
	var params db.ListUsersParams

	err := copier.Copy(&params, user)
	if err != nil {
		return nil, err
	}

	if user.Limit == nil {
		params.Limit = 10
	}

	rows, err := d.qr.ListUsers(ctx, params)
	if err != nil {
		return nil, err
	}

	var users []oapi.User
	for _, row := range rows {
		var user oapi.User

		err = copier.Copy(&user, row)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (d *UserRepository) Update(ctx context.Context, id int32, update oapi.UserUpdate) (*oapi.User, error) {
	var params db.UpdateUserParams

	err := copier.Copy(&params, update)
	if err != nil {
		return nil, err
	}

	params.ID = id

	row, err := d.qr.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	var user oapi.User

	err = copier.Copy(&user, row)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *UserRepository) Delete(ctx context.Context, id int32) error {
	rowsAffected, err := d.qr.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

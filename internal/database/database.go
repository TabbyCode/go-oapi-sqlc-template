package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jinzhu/copier"

	"github.com/xurvan/go-oapi-sqlc-template/internal/config"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/db"
	"github.com/xurvan/go-oapi-sqlc-template/internal/gen/oapi"
)

type Database struct {
	qr *db.Queries
}

func New(cfg *config.Config) (*Database, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("database: %w", err)
	}

	postgres := Database{
		qr: db.New(conn),
	}

	return &postgres, nil
}

func (d *Database) CreateUser(ctx context.Context, user oapi.UserCreate) (*oapi.User, error) {
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

func (d *Database) GetUser(ctx context.Context, id int64) (*oapi.User, error) {
	row, err := d.qr.GetUser(ctx, int32(id))
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

func (d *Database) ListUsers(ctx context.Context) ([]oapi.User, error) {
	rows, err := d.qr.ListUsers(ctx)
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

func (d *Database) UpdateUser(ctx context.Context, id int64, update oapi.UserUpdate) (*oapi.User, error) {
	var params db.UpdateUserParams
	err := copier.Copy(&params, update)
	if err != nil {
		return nil, err
	}

	params.ID = int32(id)
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

func (d *Database) DeleteUser(ctx context.Context, id int64) error {
	return d.qr.DeleteUser(ctx, int32(id))
}

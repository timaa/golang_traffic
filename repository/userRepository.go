package repository

import "github.com/jmoiron/sqlx"

type UserRepo struct {
	Db *sqlx.DB
}

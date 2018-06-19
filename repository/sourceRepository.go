package repository

import "github.com/jmoiron/sqlx"

type SourceRepo struct {
	Db *sqlx.DB
}


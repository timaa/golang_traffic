package db

import (
	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
	)



const dbType = "postgres"

type Config struct {
	ConnectionString string
}

func InitDb(cfg Config) (*sqlx.DB, error) {
	if pgConn, err := sqlx.Connect(dbType, cfg.ConnectionString); err != nil {
		return nil, err
	} else {

		if err := pgConn.Ping(); err != nil {
			return nil, err
		}
		return pgConn, nil
	}
}



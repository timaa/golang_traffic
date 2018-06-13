package daemon

import (
	"log"
	"github.com/timaa/trafficStat/db"
	"github.com/jmoiron/sqlx"
	"github.com/timaa/trafficStat/rabbit"
)

type Config struct {
	Db db.Config
	Rabbit rabbit.Config
}

func Run(cfg *Config) (*sqlx.DB, error) {
	dbConnect, err := db.InitDb(cfg.Db)
	if err != nil {
		log.Printf("Error Initializing database %v \n", err)
	}
	return  dbConnect, nil
}
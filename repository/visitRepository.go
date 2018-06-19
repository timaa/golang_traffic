package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/timaa/trafficStat/models"
	"log"
)

type VisitRepo struct {
	DB *sqlx.DB
}

func NewVisitRepository(db *sqlx.DB) (*VisitRepo) {
	return &VisitRepo{DB:db}
}


func (dr *VisitRepo) Save (visit models.Visit) {
	query := `INSERT INTO visit(id, ip, page, created_at, device_id, source_id)`
	_, err := dr.DB.NamedExec(query, visit)

	if err != nil {
		log.Fatalln(err)
	}
}
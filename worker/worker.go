package worker

import (
	"github.com/jmoiron/sqlx"
	"github.com/timaa/trafficStat/DTO"
	"fmt"
)

type Worker struct {
	Db *sqlx.DB
}



func (w *Worker) Run(db *sqlx.DB, in <-chan *DTO.TrafficDto, i int) {
	w.Db = db
	for input:= range in {
		fmt.Printf("worker #%d, value= %v \n", i, input)
		//runtime.Gosched()
	}
}
package worker

import (
	"github.com/jmoiron/sqlx"
	"github.com/timaa/trafficStat/DTO"
	"github.com/timaa/trafficStat/repository"
	"github.com/timaa/trafficStat/models"
	"log"
)

type Worker struct {
	Db *sqlx.DB
}



func (w *Worker) Run(db *sqlx.DB, in <-chan *DTO.TrafficDto, i int) {
	w.Db = db
	for input:= range in {
		deviceRepository := repository.NewDeviceRepository(db)

		device, err := deviceRepository.FindByCookie(input.Cookie)
		if err != nil {
			log.Printf("Device not found by Cookie - %s, err =%v",input.Cookie, err)
			continue
		}

		visit          := &models.Visit{}
		visit.SourceId  = input.SourceId
		visit.CreatedAt = input.CreatedAt.Time
		visit.Page      = input.Page
		visit.DeviceId  = device.Id
		visit.
		//fmt.Printf("worker #%d, value= %v \n", i, input)
		//runtime.Gosched()
		visitRepository := repository.NewDeviceRepository(db)

	}
}
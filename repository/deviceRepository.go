package repository

import (

	"github.com/jmoiron/sqlx"
	"github.com/timaa/trafficStat/models"
	"fmt"
	"log"
)

type DeviceRepo struct {
	DB *sqlx.DB
}

func NewDeviceRepository(db *sqlx.DB) (*DeviceRepo) {
	return &DeviceRepo{DB:db}
}

func (dr *DeviceRepo) GetAll() ([]*models.Device, error){
	deviceSlice := make([]*models.Device, 0)
	err := dr.DB.Select(&deviceSlice, "SELECT * from device")
	if err != nil {
		return nil, err
	}

	fmt.Printf("%v",deviceSlice)
	return deviceSlice, nil
}


func (dr *DeviceRepo) FindByCookie(cookie string) (*models.Device, error) {
	device := models.Device{}
	err := dr.DB.Get(&device, "SELECT * FROM device where cookie_id = $1", cookie)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (dr *DeviceRepo) Save (device models.Device) {
	query := `INSERT INTO device(id, user_id, source_id, cookie_id, user_agent, created_at)`
	_, err := dr.DB.NamedExec(query, device)

	if err != nil {
		log.Fatalln(err)
	}
}
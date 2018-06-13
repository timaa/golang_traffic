package repository

import (

	"github.com/jmoiron/sqlx"
	"github.com/timaa/trafficStat/models"
	"fmt"
)

type DeviceRepo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) (*DeviceRepo) {
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

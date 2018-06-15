package models

import "time"

type Visit struct {
	Id int64  `db:"id"`
	Ip int64  `db:"ip"`
	Page string `page:"page"`
	CreatedAt time.Time
	DeviceId int64  `db:"device_id"`
	SourceId int64  `db:"source_id"`
}

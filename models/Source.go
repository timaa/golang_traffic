package models

import "time"

type Source struct {
	Id 				int64     `db:"id"`
	Name 			string    `db:"name"`
	PhoneSelector   string    `db:"phone_selector"`
	EmailSelector   string    `db:"email_selector"`
	IsActive 		bool      `db:"is_active"`
	CreatedAt		time.Time `db:"created_at"`


}

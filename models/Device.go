package models

import "time"

type Device struct {
	Id 			int64 		`db:"id"`
	UserId      int64   	`db:"user_id"`
	SourceId    int64   	`db:"source_id"`
	CookieId    string  	`db:"cookie_id"`
	UserAgent   string  	`db:"user_agent"`
	CreatedAt   time.Time	`db:"created_at"`
}

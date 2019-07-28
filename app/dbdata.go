package app

import "time"

type DBData struct {
	Id   int       `json:"id"`
	Date time.Time `json:date`
	Name string    `json:name`
}

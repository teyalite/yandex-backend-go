package db

import "github.com/lib/pq"

type Courier struct {
	Id           int `gorm:"primaryKey"`
	CourierType  string
	Regions      pq.Int32Array  `gorm:"type:integer[]"`
	WorkingHours pq.StringArray `gorm:"type:text[]"`
}

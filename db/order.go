package db

import "github.com/lib/pq"

type Order struct {
	Id            int `gorm:"primaryKey"`
	Weight        float64
	Regions       int32
	DeliveryHours pq.StringArray `gorm:"type:text[]"`
	Cost          int32
	CompletedTime string `gorm:"default:custom_null"`
}

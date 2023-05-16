package main

import "enrollments/22366-khaidara.a-phystech.edu-89/db"

func main() {
	db.LoadEnv()
	db.ConnectDB()
	db.DB.AutoMigrate(&db.Order{}, &db.Courier{}, &db.OrderCourier{})
}

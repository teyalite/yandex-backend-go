package main

import (
	"enrollments/22366-khaidara.a-phystech.edu-89/controllers"
	"enrollments/22366-khaidara.a-phystech.edu-89/db"

	"github.com/gin-gonic/gin"
)

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func main() {
	db.LoadEnv()
	db.ConnectDB()
	db.DB.AutoMigrate(&db.Order{}, &db.Courier{}, &db.OrderCourier{})

	router := gin.Default()

	router.Use(JSONMiddleware())

	orders := router.Group("/orders")
	{
		orders.GET("/", controllers.GetOrders)

		orders.POST("/", controllers.CreateOrders)

		orders.GET("/:orderId", controllers.GetOrder)

		orders.POST("/complete", controllers.CompleteOrder)

		// orders.POST("/assign", controllers.OrdersAssign)

	}

	couriers := router.Group("/couriers")
	{

		couriers.GET("/", controllers.GetCouriers)

		couriers.POST("/", controllers.CreateCourier)

		couriers.GET("/:courierId", controllers.GetCourier)

		couriers.GET("/meta-info/:courierId", controllers.GetCourierMetaInfo)

		// couriers.GET("/assignments", couriersAssignments)

	}

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}

package controllers

import (
	"enrollments/22366-khaidara.a-phystech.edu-89/db"
	"enrollments/22366-khaidara.a-phystech.edu-89/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func convertOrder(order db.Order) gin.H {
	m_order := gin.H{"regions": order.Regions, "delivery_hours": order.DeliveryHours, "order_id": order.Id, "cost": order.Cost, "weight": order.Weight}

	if order.CompletedTime != db.CompletedTimeDefault {
		m_order["completed_time"] = order.CompletedTime
	}

	return m_order
}

func GetOrders(c *gin.Context) {
	var limit int
	var offset int

	s, err := strconv.ParseInt(c.Query("limit"), 10, 64)

	if err != nil {
		limit = 1
	} else {
		limit = int(s)
	}

	s_, err_ := strconv.ParseInt(c.Query("offset"), 10, 64)

	if err_ != nil {
		offset = 0
	} else {
		offset = int(s_)
	}

	orders_query := []db.Order{}

	db.DB.Limit(limit).Offset(offset).Find(&orders_query)

	orders := []gin.H{}

	for _, order := range orders_query {
		orders = append(orders, convertOrder(order))
	}

	c.JSON(http.StatusOK, orders)
}

func CreateOrders(c *gin.Context) {
	var ordersReq schemas.CreateOrderRequest

	if err := c.BindJSON(&ordersReq); err != nil {
		var badRequestResponse schemas.BadRequestResponse

		badRequestResponse.Message = "Invalid request body"
		c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
		return
	}

	var orders_query []db.Order

	for _, order := range ordersReq.Orders {
		orders_query = append(orders_query, db.Order{Regions: order.Regions, Cost: order.Cost, Weight: order.Weight, DeliveryHours: pq.StringArray(order.DeliveryHours)})
	}

	db.DB.Create(&orders_query)

	var orders []gin.H

	for _, order := range orders_query {
		orders = append(orders, convertOrder(order))
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	var order_id int64
	var err error

	order_id, err = strconv.ParseInt(c.Params.ByName("orderId"), 10, 64)

	if err != nil {

		var badRequestResponse schemas.BadRequestResponse

		badRequestResponse.Message = "Invalid request body"
		c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
		return
	}

	var order db.Order

	result := db.DB.First(&order, order_id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, convertOrder(order))
}

func CompleteOrder(c *gin.Context) {
	var completeOrderReq schemas.CompleteOrderRequestDto

	if err := c.BindJSON(&completeOrderReq); err != nil {
		var badRequestResponse schemas.BadRequestResponse

		badRequestResponse.Message = "Invalid request body"
		c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
		return
	}

	orders_query := []db.Order{}

	db.DB.Transaction(func(tx *gorm.DB) error {

		var oc_query []db.OrderCourier

		for _, info := range completeOrderReq.CompleteInfo {
			oc_query = append(oc_query, db.OrderCourier{OrderId: int(info.OrderId), CourierId: int(info.CourierId)})

			var order db.Order

			if err := db.DB.First(&order, int(info.OrderId)).Error; err != nil {
				return err
			}

			order.CompletedTime = info.CompleteTime

			if err := db.DB.Save(&order).Error; err != nil {
				return err
			}

			orders_query = append(orders_query, order)
		}

		db.DB.Create(&oc_query)

		return nil
	})

	var orders []gin.H

	for _, order := range orders_query {
		orders = append(orders, convertOrder(order))
	}

	c.JSON(http.StatusOK, orders)
}

func OrdersAssign(c *gin.Context) {}

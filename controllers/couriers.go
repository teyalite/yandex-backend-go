package controllers

import (
	"enrollments/22366-khaidara.a-phystech.edu-89/db"
	"enrollments/22366-khaidara.a-phystech.edu-89/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func converCourier(courier db.Courier) gin.H {

	m_courier := gin.H{"regions": courier.Regions, "working_hours": courier.WorkingHours, "courier_id": courier.Id, "courier_type": courier.CourierType}

	return m_courier
}

func GetCouriers(c *gin.Context) {
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

	var couriers_query []db.Courier

	db.DB.Limit(limit).Offset(offset).Find(&couriers_query)

	couriers := []gin.H{}

	for _, courier := range couriers_query {
		couriers = append(couriers, converCourier(courier))
	}

	c.JSON(http.StatusOK, gin.H{"couriers": couriers, "limit": limit, "offset": offset})
}

func CreateCourier(c *gin.Context) {
	var couriersReq schemas.CreateCourierRequest

	if err := c.BindJSON(&couriersReq); err != nil {
		var badRequestResponse schemas.BadRequestResponse

		badRequestResponse.Message = "Invalid request body"
		c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
		return
	}

	var couriers_query []db.Courier

	for _, courier := range couriersReq.Couriers {
		if courier.CourierType != "FOOT" && courier.CourierType != "BIKE" && courier.CourierType != "AUTO" {
			var badRequestResponse schemas.BadRequestResponse

			badRequestResponse.Message = "Invalid request body"
			c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
			return
		}

		couriers_query = append(couriers_query, db.Courier{Regions: pq.Int32Array(courier.Regions), CourierType: courier.CourierType, WorkingHours: pq.StringArray(courier.WorkingHours)})
	}

	db.DB.Create(&couriers_query)

	// if err := db.DB.Create(&couriers_query).Error; err != nil {

	// 	return
	// }

	var couriers []gin.H

	for _, order := range couriers_query {
		couriers = append(couriers, converCourier(order))
	}

	c.JSON(http.StatusOK, couriers)
}

func GetCourier(c *gin.Context) {
	var courier_id int64
	var err error

	courier_id, err = strconv.ParseInt(c.Params.ByName("courierId"), 10, 64)

	if err != nil {

		var badRequestResponse schemas.BadRequestResponse

		badRequestResponse.Message = "Invalid request body"
		c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
		return
	}

	var courier db.Courier

	result := db.DB.First(&courier, courier_id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, converCourier(courier))
}

func GetCourierMetaInfo(c *gin.Context) {

	var courier_id int64
	var err error

	courier_id, err = strconv.ParseInt(c.Params.ByName("courierId"), 10, 64)

	if err != nil {

		var badRequestResponse schemas.BadRequestResponse

		badRequestResponse.Message = "Invalid request body"
		c.IndentedJSON(http.StatusBadRequest, badRequestResponse)
		return
	}

	var courier db.Courier

	result := db.DB.First(&courier, courier_id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	s := converCourier(courier)

	s["rating"] = 0
	s["earnings"] = 0

	c.JSON(http.StatusOK, s)
}

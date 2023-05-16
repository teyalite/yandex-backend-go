package schemas

type BadRequestResponse struct {
	Message string `json:"message"`
}

type CreateOrderDto struct {
	Weight        float64  `json:"weight" binding:"required"`
	Regions       int32    `json:"regions" binding:"required"`
	DeliveryHours []string `json:"delivery_hours" binding:"required"`
	Cost          int32    `json:"cost" binding:"required"`
}

type CreateOrderRequest struct {
	Orders []CreateOrderDto `json:"orders" binding:"required,dive"`
}

type OrderDto struct {
	Weight        float64  `json:"weight" binding:"required"`
	Regions       int32    `json:"regions" binding:"required"`
	DeliveryHours []string `json:"delivery_hours" binding:"required"`
	Cost          int32    `json:"cost" binding:"required"`
	CompletedTime string   `json:"completed_time"`
}

//
type CreateCourierDto struct {
	Regions      []int32  `json:"regions" binding:"required"`
	WorkingHours []string `json:"working_hours" binding:"required"`
	CourierType  string   `json:"courier_type" binding:"required"`
}

type CreateCourierRequest struct {
	Couriers []CreateCourierDto `json:"couriers" binding:"required,dive"`
}

//
type CompleteOrder struct {
	CourierId    int64  `json:"courier_id" binding:"required"`
	OrderId      int64  `json:"order_id" binding:"required"`
	CompleteTime string `json:"complete_time" binding:"required"`
}

type CompleteOrderRequestDto struct {
	CompleteInfo []CompleteOrder `json:"complete_info" binding:"required,dive"`
}

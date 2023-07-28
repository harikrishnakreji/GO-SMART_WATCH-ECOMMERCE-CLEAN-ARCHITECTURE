package models

type Cart struct {
	ProductID  uint    `json:"product_id"`
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CartResponse struct {
	UserName   string
	TotalPrice float64
	Cart       []Cart
}
type CartTotal struct {
	UserName   string  `json:"user_name"`
	TotalPrice float64 `json:"total_price"`
}

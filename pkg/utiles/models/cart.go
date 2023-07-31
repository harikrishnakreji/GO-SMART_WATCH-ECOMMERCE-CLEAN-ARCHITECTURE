package models

type Carts struct {
	ProductID  uint    `json:"product_id"`
	Name       string  `json:"name"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CartsResponse struct {
	UserName   string
	TotalPrice float64
	Carts      []Carts
}
type CartsTotal struct {
	UserName   string  `json:"user_name"`
	TotalPrice float64 `json:"total_price"`
}

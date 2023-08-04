package models

type ProductResponse struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	CategoryName        string  `json:"category_name"`
	ProductsDescription string  `json:"product_description"`
	BrandName           string  `json:"brand_name"`
	Quantity            int     `json:"quantity"`
	Price               float64 `json:"price"`
}

type ShowProducts struct {
	TottalCount   int
	ProductsBrief []ProductsBrief
}

type ProductsReceiver struct {
	Name                string  `json:"name" binding:"required"`
	CategoryID          uint    `json:"category_id" binding:"required"`
	ProductsDescription string  `json:"products_description" binding:"required"`
	BrandID             uint    `json:"brand_id" binding:"required"`
	Quantity            int     `json:"quantity" binding:"required"`
	Price               float64 `json:"price" binding:"required"`
}

type ProductsBrief struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	ProductStatus string  `json:"product_status"`
}

type CategoryUpdate struct {
	Category string `json:"category" binding:"required"`
}

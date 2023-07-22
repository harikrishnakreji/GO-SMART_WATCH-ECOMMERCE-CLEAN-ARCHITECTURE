package models

type ProductResponse struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	GenreName           string  `json:"genre_name"`
	ProductsDescription string  `json:"product_description"`
	BrandName           string  `json:"brand_namess"`
	Quantity            int     `json:"quantity"`
	Price               float64 `json:"price"`
}

type ProductsReceiver struct {
	Name                string  `json:"name" binding:"required"`
	GenreID             uint    `json:"genre_id" binding:"required"`
	ProductsDescription string  `json:"products_description" binding:"required"`
	BrandID             uint    `json:"brand_id" binding:"required"`
	Quantity            int     `json:"quantity" binding:"required"`
	Price               float64 `json:"price" binding:"required"`
}

type ProductsBrief struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Genre         string  `json:"genre"`
	Price         float64 `json:"price"`
	Quantity      int     `json:"quantity"`
	ProductStatus string  `json:"product_status"`
}

type CategoryUpdate struct {
	Genre string `json:"genre" binding:"required"`
}

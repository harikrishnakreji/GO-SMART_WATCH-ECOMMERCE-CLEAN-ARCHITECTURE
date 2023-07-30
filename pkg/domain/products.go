package domain

type Products struct {
	ID                  uint     `json:"id" gorm:"unique;not null"`
	Sku                 string   `json:"sku"`
	Name                string   `json:"name"`
	CategoryID          uint     `json:"category_id"`
	Category            Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductsDescription string   `json:"products_discription"`
	BrandID             uint     `json:"brand_id"`
	Brand               Brand    `json:"-" gorm:"foreignkey:BrandID;constraint:OnDelete:CASCADE"`
	Quantity            int      `json:"quantity"`
	Price               float64  `json:"price"`
	Delete              bool     `json:"delete" gorm:"default:false"`
}

type Category struct {
	ID           uint   `json:"id" gorm:"unique; not null"`
	CategoryName string `json:"category_name"`
}

type Brand struct {
	ID        uint   `json:"id" gorm:"unique; not null"`
	BrandName string `json:"brand_name"`
}

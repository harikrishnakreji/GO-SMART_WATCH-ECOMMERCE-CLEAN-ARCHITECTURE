package repository

import (
	"errors"
	"strconv"

	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/domain"
	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

func (p *productDatabase) ShowAllProducts(page int, count int) ([]models.ProductsBrief, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var productsBrief []models.ProductsBrief
	err := p.DB.Raw(`
		SELECT products.id, products.name,genres.genre_name AS genre,products.price,products.quantity
		FROM products
		JOIN genres ON products.genre_id = genres.id
		 limit ? offset ?
	`, count, offset).Scan(&productsBrief).Error

	if err != nil {
		return nil, err
	}

	return productsBrief, nil

}

func (p *productDatabase) AddProduct(product models.ProductsReceiver) (models.ProductResponse, error) {

	var id int

	sku := product.Name
	err := p.DB.Raw("INSERT INTO products (name, genre_id, products_description, brand_id, quantity, price, sku) VALUES (?,?, ?, ?, ?, ?, ?) RETURNING id", product.Name, product.GenreID, product.ProductsDescription, product.BrandID, product.Quantity, product.Price, sku).Scan(&id).Error
	if err != nil {
		return models.ProductResponse{}, err
	}

	var productResponse models.ProductResponse
	err = p.DB.Raw(`
	SELECT
		p.id,
		p.sku,
		p.name,
		g.genre_name,
		p.products_description,
		s.brand_name,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			genres g ON p.genre_id = g.id
		JOIN
			brands s ON p.brand_id = s.id 
		WHERE
			p.id = ?
			`, id).Scan(&productResponse).Error

	if err != nil {
		return models.ProductResponse{}, err
	}

	return productResponse, nil

}

func (p *productDatabase) DeleteProduct(product_id string) error {

	id, _ := strconv.Atoi(product_id)
	result := p.DB.Exec("delete from products where id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records were of that id exists")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (pr *productDatabase) GetGenres() ([]domain.Genre, error) {

	var genres []domain.Genre
	if err := pr.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

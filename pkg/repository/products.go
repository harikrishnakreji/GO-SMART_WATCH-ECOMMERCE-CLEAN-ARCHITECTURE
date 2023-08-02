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
		SELECT products.id, products.name,categories.category_name AS category,products.price,products.quantity
		FROM products
		JOIN categories ON products.category_id = categories.id
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
	err := p.DB.Raw("INSERT INTO products (name, category_id, products_description, brand_id, quantity, price, sku) VALUES (?,?, ?, ?, ?, ?, ?) RETURNING id", product.Name, product.CategoryID, product.ProductsDescription, product.BrandID, product.Quantity, product.Price, sku).Scan(&id).Error
	if err != nil {
		return models.ProductResponse{}, err
	}

	var productResponse models.ProductResponse
	err = p.DB.Raw(`
	SELECT
		p.id,
		p.sku,
		p.name,
		g.category_name,
		p.products_description,
		s.brand_name,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			categories g ON p.category_id = g.id
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

func (p *productDatabase) GetQuantityFromProductID(id int) (int, error) {

	var quantity int
	err := p.DB.Raw("select quantity from products where id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}

	return quantity, nil

}

func (p *productDatabase) DoesProductExist(productID int) (bool, error) {

	var count int
	err := p.DB.Raw("select count(*) from products where id = ?", productID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (pr *productDatabase) GetCategorys() ([]domain.Category, error) {

	var categories []domain.Category
	if err := pr.DB.Raw("select * from categories").Scan(&categories).Error; err != nil {
		return []domain.Category{}, err
	}

	return categories, nil

}

func (pr *productDatabase) GetPriceOfProductFromID(productID int) (float64, error) {

	var productPrice float64
	if err := pr.DB.Raw("select price from products where id = ?", productID).Scan(&productPrice).Error; err != nil {
		pr.DB.Rollback()
		return 0.0, err
	}

	return productPrice, nil

}

// detailed product details
func (p *productDatabase) ShowIndividualProducts(product_id string) (models.ProductResponse, error) {
	id, _ := strconv.Atoi(product_id)
	var product models.ProductResponse
	err := p.DB.Raw(`
	SELECT
		p.id,
		p.name,
		g.category_name,
		p.products_description,
		s.brand_name,
		p.quantity,
		p.price
		FROM
			products p
		JOIN
			categories g ON p.category_id = g.id
		JOIN
			brands s ON p.brand_id = s.id 
		WHERE
			p.id = ?
			`, id).Scan(&product).Error

	if err != nil {
		return models.ProductResponse{}, errors.New("error retrieved record")
	}

	return product, nil

}

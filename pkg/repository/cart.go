package repository

import (
	"errors"

	interfaces "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/repository/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"gorm.io/gorm"
)

type cartsRepository struct {
	DB *gorm.DB
}

func NewCartsRepository(DB *gorm.DB) interfaces.CartsRepository {
	return &cartsRepository{
		DB: DB,
	}
}

func (cr *cartsRepository) QuantityOfProductInCarts(userID int, product_id int) (int, error) {
	var cartsQuantity int
	if err := cr.DB.Raw("select quantity from carts where user_id=? and product_id=?", userID, product_id).Scan(&cartsQuantity).Error; err != nil {
		return 0, err
	}
	return cartsQuantity, nil
}
func (cr *cartsRepository) AddItemToCarts(userID int, product_id int, quantity int, productPrice float64) error {

	if err := cr.DB.Exec("insert into carts(user_id,product_id,quantity,total_price)values(?,?,?,?)", userID, product_id, quantity, productPrice).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartsRepository) TotalPriceForProductInCarts(userID int, productID int) (float64, error) {

	var totalPrice float64
	if err := cr.DB.Raw("select sum(total_price) as total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&totalPrice).Error; err != nil {
		return 0.0, err
	}

	return totalPrice, nil
}

func (cr *cartsRepository) UpdateCarts(quantity int, price float64, userID int, product_id int) error {

	if err := cr.DB.Exec("update carts set quantity = ?, total_price = ? where user_id = ? and product_id = ?", quantity, price, userID, product_id).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartsRepository) GetTotalPrice(userID int) (models.CartsTotal, error) {

	var cartsTotal models.CartsTotal
	err := cr.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&cartsTotal.TotalPrice).Error
	if err != nil {
		return models.CartsTotal{}, err
	}

	err = cr.DB.Raw("select name as user_name from users where id = ?", userID).Scan(&cartsTotal.UserName).Error
	if err != nil {
		return models.CartsTotal{}, err
	}

	return cartsTotal, nil

}

func (cr *cartsRepository) GetQuantityAndTotalPrice(userID int, productID int, cartsDetails struct {
	Quantity   int
	TotalPrice float64
}) (struct {
	Quantity   int
	TotalPrice float64
}, error) {

	// select quantity and totalprice = quantity * indiviualproductpriice from carts
	if err := cr.DB.Raw("select quantity,total_price from carts where user_id = ? and product_id = ?", userID, productID).Scan(&cartsDetails).Error; err != nil {
		return struct {
			Quantity   int
			TotalPrice float64
		}{}, err
	}

	return cartsDetails, nil

}

func (cr *cartsRepository) RemoveProductFromCarts(userID int, product_id int) error {

	if err := cr.DB.Exec("delete from carts where user_id = ? and product_id = ?", uint(userID), uint(product_id)).Error; err != nil {
		return err
	}

	return nil
}

func (cr *cartsRepository) UpdateCartsDetails(cartsDetails struct {
	Quantity   int
	TotalPrice float64
}, userID int, productID int) error {

	if err := cr.DB.Exec("update carts set quantity = ?,total_price = ? where user_id = ? and product_id = ?", cartsDetails.Quantity, cartsDetails.TotalPrice, userID, productID).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartsRepository) RemoveFromCarts(userID int) ([]models.Carts, error) {

	var cartsResponse []models.Carts
	if err := cr.DB.Raw("select carts.product_id,products.name as name,carts.quantity,carts.total_price from carts inner join products on carts.product_id = products.id where carts.user_id = ?", userID).First(&cartsResponse).Error; err != nil {
		return []models.Carts{}, err
	}

	return cartsResponse, nil

}

func (cr *cartsRepository) DisplayCarts(userID int) ([]models.Carts, error) {

	var count int
	if err := cr.DB.Raw("select count(*) from carts where user_id = ? ", userID).First(&count).Error; err != nil {
		return []models.Carts{}, err
	}

	if count == 0 {
		return []models.Carts{}, nil
	}

	var cartsResponse []models.Carts

	if err := cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.name as name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartsResponse).Error; err != nil {
		return []models.Carts{}, err
	}

	return cartsResponse, nil

}

func (cr *cartsRepository) EmptyCarts(userID int) error {

	if err := cr.DB.Exec("delete from carts where user_id = ? ", userID).Error; err != nil {
		return err
	}

	return nil

}

func (cr *cartsRepository) GetAllItemFromCarts(userID int) ([]models.Carts, error) {

	var count int

	var cartsResponse []models.Carts
	err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Carts{}, err
	}

	if count == 0 {
		return []models.Carts{}, nil
	}

	err = cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.name as name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartsResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartsResponse) == 0 {
				return []models.Carts{}, nil
			}
			return []models.Carts{}, err
		}
		return []models.Carts{}, err
	}

	return cartsResponse, nil

}

func (cr *cartsRepository) CheckProduct(product_id int) (bool, string, error) {

	var count int
	err := cr.DB.Raw("select count(*) from products where id = ?", product_id).Scan(&count).Error
	if err != nil {
		return false, "", err
	}

	var category string
	if count > 0 {
		err := cr.DB.Raw("select categories.category_name from categories inner join products on products.category_id = categories.id where products.id = ?", product_id).Scan(&category).Error
		if err != nil {
			return false, "", err
		}
		return true, category, nil
	}
	return false, "", nil

}

func (cr *cartsRepository) ProductExist(product_id int, userID int) (bool, error) {

	var count int
	err := cr.DB.Raw("select count(*) from carts where user_id = ? and product_id = ?", userID, product_id).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func (cr *cartsRepository) GetTotalPriceFromCarts(userID int) (float64, error) {

	var totalPrice float64
	err := cr.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&totalPrice).Error
	if err != nil {
		return 0.0, err
	}

	return totalPrice, nil

}

func (cr *cartsRepository) DoesCartsExist(userID int) (bool, error) {

	count := 0
	err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}

	return true, nil
}

func (cr *cartsRepository) GetAllItemsFromCarts(userID int) ([]models.Carts, error) {

	var count int

	var cartsResponse []models.Carts
	err := cr.DB.Raw("select count(*) from carts where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Carts{}, err
	}

	if count == 0 {
		return []models.Carts{}, nil
	}

	err = cr.DB.Raw("select carts.user_id,users.name as user_name,carts.product_id,products.name as name,carts.quantity,carts.total_price from carts inner join users on carts.user_id = users.id inner join products on carts.product_id = products.id where user_id = ?", userID).First(&cartsResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartsResponse) == 0 {
				return []models.Carts{}, nil
			}
			return []models.Carts{}, err
		}
		return []models.Carts{}, err
	}

	return cartsResponse, nil

}

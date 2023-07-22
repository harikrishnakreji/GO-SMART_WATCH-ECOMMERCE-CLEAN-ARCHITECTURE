package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/response"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(useCase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: useCase,
	}
}

// @Summary Get Products Details to users
// @Description Retrieve all product Details with pagination to users
// @Tags User Product
// @Accept json
// @Produce json
// @Param page path string true "Page number"
// @Param count query string true "Page Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /products/page/{page} [get]
func (pr *ProductHandler) ShowAllProducts(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	count, err := strconv.Atoi(c.Query("count"))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := pr.productUseCase.ShowAllProducts(page, count)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Product Details To Admin
// @Description Retrieve product Details with pagination to Admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Param count query string true "Products Count Per Page"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products [get]
func (pr *ProductHandler) SeeAllProductToAdmin(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)

	count, _ := strconv.Atoi(c.Query("count"))

	products, err := pr.productUseCase.ShowAllProductsToAdmin(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	fmt.Println(products)

	successRes := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products to admin side", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add Products
// @Description Add product from admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param product body models.ProductsReceiver true "Product details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/add-product/ [post]
func (pr *ProductHandler) AddProduct(c *gin.Context) {

	var product models.ProductsReceiver

	if err := c.ShouldBindJSON(&product); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	fmt.Println(product)

	productResponse, err := pr.productUseCase.AddProduct(product)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not add the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added products", productResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Delete product
// @Description Delete a product from the admin side
// @Tags Admin Product Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/products/delete-product/{id} [delete]
func (pr *ProductHandler) DeleteProduct(c *gin.Context) {

	product_id := c.Param("id")
	err := pr.productUseCase.DeleteProduct(product_id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the item", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

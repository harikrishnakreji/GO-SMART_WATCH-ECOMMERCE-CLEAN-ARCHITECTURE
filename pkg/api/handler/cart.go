package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/response"
)

type CartsHandler struct {
	cartsUseCase services.CartsUseCase
}

func NewCartsHandler(usecase services.CartsUseCase) *CartsHandler {
	return &CartsHandler{
		cartsUseCase: usecase,
	}
}

// @Summary: Add to Carts
// @Description Adding products to carts using product id
// @Tags User Carts
// @Accept json
// @Produce json
// @Param id path string true "product_id"
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /carts/addtocarts/{id} [post]
func (cr *CartsHandler) AddToCarts(c *gin.Context) {
	id := c.Param("id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product id not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	userID, _ := c.Get("user_id")
	cartsResponse, err := cr.cartsUseCase.AddToCarts(product_id, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add product to carts", cartsResponse, nil)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "successfully added product to the carts", cartsResponse, nil)
	c.JSON(http.StatusCreated, successRes)
}

// @Summary: Removing product from carts
// @Description Remove specified product of quantity 1 from carts using product id
// @Tags User Carts
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /carts/removefromcarts/{id} [delete]
func (cr *CartsHandler) RemoveFromCarts(c *gin.Context) {

	id := c.Param("id")
	product_id, err := strconv.Atoi(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "product not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, _ := c.Get("user_id")
	updatedCarts, err := cr.cartsUseCase.RemoveFromCarts(product_id, userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not delete carts items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Carts item deleted", updatedCarts, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Display Carts
// @Description Display all products of the carts along with price of the product and grand total
// @Tags User Carts
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /carts [get]
func (cr *CartsHandler) DisplayCarts(c *gin.Context) {

	userID, _ := c.Get("user_id")
	carts, err := cr.cartsUseCase.DisplayCarts(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not display carts items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Carts items displayed successfully", carts, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Delete all Items Present inside the Carts
// @Description Remove all product from carts
// @Tags User Carts
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /carts [delete]
func (cr *CartsHandler) EmptyCarts(c *gin.Context) {

	userID, _ := c.Get("user_id")
	emptyCarts, err := cr.cartsUseCase.EmptyCarts(userID.(int))

	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not display carts items", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "carts items displayed successfully", emptyCarts, nil)
	c.JSON(http.StatusOK, successRes)
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/usecase/interface"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/models"
	"github.com/harikrishnakreji/GO-SMART_WATCH-ECOMMERCE-CLEAN-ARCHITECTURE/pkg/utiles/response"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// @Summary Admin Login
// @Description Login handler for admin
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Param  admin body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/adminlogin [post]
func (cr *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := cr.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Admin Signup
// @Description Signup handler for admin
// @Tags Admin Authentication
// @Accept json
// @Produce json
// @Security Bearer
// @Param  admin body models.AdminSignUp true "Admin login details"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/createadmin [post]
func (cr *AdminHandler) CreateAdmin(c *gin.Context) {

	var admin models.AdminSignUp
	if err := c.ShouldBindJSON(&admin); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	adminDetails, err := cr.adminUseCase.CreateAdmin(admin)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully signed up the user", adminDetails, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Get Users Details To Admin
// @Description Retrieve users with pagination to admin side
// @Tags Admin User Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page path string true "Page number"
// @Param pageSize query string true "page size"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/users/{page} [get]
func (ad *AdminHandler) GetUsers(c *gin.Context) {

	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.adminUseCase.GetUsers(page, pageSize)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Category Details to admin side
// @Description Display Category details on the admin side
// @Tags Admin Category Management
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/categories [get]
func (ad *AdminHandler) GetCategorys(c *gin.Context) {

	categories, err := ad.adminUseCase.GetCategorys()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Add a new Categorys ( Category )
// @Description Add a new Category So that movie of that category can be added
// @Tags Admin Category Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param  category body models.CategoryUpdate true "Update Category"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/categories/add_category [POST]
func (ad *AdminHandler) AddCategorys(c *gin.Context) {

	var category models.CategoryUpdate
	if err := c.ShouldBindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := ad.adminUseCase.AddCategorys(category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "The category could not be added", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Successfully added the category", nil, nil)
	c.JSON(http.StatusCreated, successRes)

}

// @Summary Delete Category ( Category )
// @Description Delete Category for existing films and delete the films along with it
// @Tags Admin Category Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "category-id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/categories/delete_category/{id} [POST]
func (ad *AdminHandler) DeleteCategory(c *gin.Context) {

	category_id := c.Param("id")
	err := ad.adminUseCase.Delete(category_id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not delete the specified category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the product", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Block  user
// @Description Block an existing user using user id
// @Tags Admin User Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "user-id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/users/block-users/{id} [POST]
func (ad *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Param("id")
	err := ad.adminUseCase.BlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Unblock  User
// @Description Unblock an already blocked user using user id
// @Tags Admin User Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "user-id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/users/unblock-users/{id} [POST]
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {

	id := c.Param("id")
	err := ad.adminUseCase.UnBlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Admin Dashboard
// @Description Get Amin Home Page with Complete Details
// @Tags Admin Dash Board
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/dashboard [GET]
func (ad *AdminHandler) DashBoard(c *gin.Context) {

	adminDashBoard, err := ad.adminUseCase.DashBoard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "admin dashboard displayed fine", adminDashBoard, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Filtered Sales Report
// @Description Get Filtered sales report by week, month and year
// @Tags Admin Dash Board
// @Accept json
// @Produce json
// @Security Bearer
// @Param period path string true "sales report"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/sales-report/{period} [GET]
func (ad *AdminHandler) FilteredSalesReport(c *gin.Context) {

	timePeriod := c.Param("period")
	salesReport, err := ad.adminUseCase.FilteredSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", salesReport, nil)
	c.JSON(http.StatusOK, successRes)

}

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vitalicher97/psychologist_app/internal/app/db/models"
)

// CreateCustomer handles the creation of a new customer.
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.Error(err)
		return
	}

	if err := customer.Create(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, customer)
}

// GetCustomer handles retrieving a customer by ID.
func GetCustomer(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	customer := &models.Customer{ID: id}

	customer, err = customer.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer handles updating a customer's data by ID.
func UpdateCustomer(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	existingCustomer := &models.Customer{ID: id}
	existingCustomer, err = existingCustomer.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.Error(err)
		return
	}

	customer.CreatedAt = existingCustomer.CreatedAt
	customer.ID = id

	if err := customer.Update(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

// DeleteCustomer handles deleting a customer by ID.
func DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	customer := &models.Customer{ID: id}

	if err := customer.DeleteByID(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetAllCustomers handles retrieving a list of all customers.
func GetAllCustomers(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		customer, err := models.GetCustomerByEmail(c, email)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, customer)
		return
	}

	customers, err := models.GetAllCustomers(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

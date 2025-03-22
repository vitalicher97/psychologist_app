package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"

	"github.com/vitalicher97/psychologist_app/internal/app/db/models"
)

// GetCustomerPsychologistPrices handles retrieving a customer-psychologist price entry by customer and psychologist ID.
func GetCustomerPsychologistPrices(c *gin.Context) {
	customerID := c.Query("customer_id")
	psychologistID := c.Query("psychologist_id")

	customerIDInt, err := strconv.Atoi(customerID)
	if err != nil {
		c.Error(err)
		return
	}

	psychologistIDInt, err := strconv.Atoi(psychologistID)
	if err != nil {
		c.Error(err)
		return
	}

	price, err := models.GetFixedPrice(customerIDInt, psychologistIDInt)
	if err != nil && err != pg.ErrNoRows {
		c.Error(err)
		return
	}

	if err == pg.ErrNoRows {
		price = &models.CustomerPsychologistPrices{}
	}

	c.JSON(http.StatusOK, price)
}

// CreateCustomerPsychologistPrices handles the creation of a new customer-psychologist price entry.
func CreateCustomerPsychologistPrices(c *gin.Context) {
	var customerPsychologistPrices models.CustomerPsychologistPrices
	err := c.ShouldBindJSON(&customerPsychologistPrices)
	if err != nil {
		c.Error(err)
		return
	}

	existingPrice, err := models.GetFixedPrice(customerPsychologistPrices.CustomerID, customerPsychologistPrices.PsychologistID)
	if err != nil && err != pg.ErrNoRows {
		c.Error(err)
		return
	}

	if existingPrice != nil {
		c.JSON(http.StatusCreated, customerPsychologistPrices)
		return
	}

	err = customerPsychologistPrices.Create()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, customerPsychologistPrices)
}

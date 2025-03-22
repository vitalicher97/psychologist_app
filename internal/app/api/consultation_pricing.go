package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vitalicher97/psychologist_app/internal/app/db/models"
)

// CreateConsultationPricing handles the creation of a new consultation pricing entry.
func CreateConsultationPricing(c *gin.Context) {
	var consultationPricing models.ConsultationPricing
	if err := c.ShouldBindJSON(&consultationPricing); err != nil {
		c.Error(err)
		return
	}

	if err := consultationPricing.Create(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, consultationPricing)
}

// GetConsultationPricing handles retrieving a consultation pricing entry by ID.
func GetConsultationPricing(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	consultationPricing := &models.ConsultationPricing{ID: id}

	consultationPricing, err = consultationPricing.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, consultationPricing)
}

// UpdateConsultationPricing handles updating a consultation pricing entry by ID.
func UpdateConsultationPricing(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	existingConsultationPricing := &models.ConsultationPricing{ID: id}
	existingConsultationPricing, err = existingConsultationPricing.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var consultationPricing models.ConsultationPricing
	if err := c.ShouldBindJSON(&consultationPricing); err != nil {
		c.Error(err)
		return
	}

	consultationPricing.CreatedAt = existingConsultationPricing.CreatedAt
	consultationPricing.ID = id

	if err := consultationPricing.Update(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, consultationPricing)
}

// DeleteConsultationPricing handles deleting a consultation pricing entry by ID.
func DeleteConsultationPricing(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	consultationPricing := &models.ConsultationPricing{ID: id}

	if err := consultationPricing.DeleteByID(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetAllConsultationPricing handles retrieving a list of all consultation pricing entries.
func GetAllConsultationPricing(c *gin.Context) {
	psychologistIDStr := c.Query("psychologist")

	var pricingList []models.ConsultationPricing
	var psychologistID int
	var err error

	if psychologistIDStr != "" {
		psychologistID, err = strconv.Atoi(psychologistIDStr)
		if err != nil {
			c.Error(err)
			return
		}

		pricingList, err = models.GetConsultationPricingByPsychologistID(c, psychologistID)
		if err != nil {
			c.Error(err)
			return
		}
	} else {
		pricingList, err = models.GetAllConsultationPricing(c)
		if err != nil {
			c.Error(err)
			return
		}
	}

	if len(pricingList) == 0 {
		pricingList = []models.ConsultationPricing{}
	}

	c.JSON(http.StatusOK, pricingList)
}

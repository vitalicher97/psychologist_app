package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vitalicher97/psychologist_app/internal/app/db/models"
)

// CreateAvailability handles the creation of a new availability entry.
func CreateAvailability(c *gin.Context) {
	var availability models.Availability
	if err := c.ShouldBindJSON(&availability); err != nil {
		c.Error(err)
		return
	}

	if err := availability.Create(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, availability)
}

// GetAvailability handles retrieving an availability entry by ID.
func GetAvailability(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	availability := &models.Availability{ID: id}

	availability, err = availability.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, availability)
}

// UpdateAvailability handles updating an availability entry by ID.
func UpdateAvailability(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	existingAvailability := &models.Availability{ID: id}
	existingAvailability, err = existingAvailability.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var availability models.Availability
	if err := c.ShouldBindJSON(&availability); err != nil {
		c.Error(err)
		return
	}

	availability.CreatedAt = existingAvailability.CreatedAt
	availability.ID = id

	if err := availability.Update(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, availability)
}

// DeleteAvailability handles deleting an availability entry by ID.
func DeleteAvailability(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	availability := &models.Availability{ID: id}

	if err := availability.DeleteByID(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ListAvailability handles retrieving a list of all availability entries.
func GetAllAvailability(c *gin.Context) {
	psychologistIDStr := c.Query("psychologist")

	var availabilityList []models.Availability
	var err error
	if psychologistIDStr != "" {
		psychologistID, err := strconv.Atoi(psychologistIDStr)
		if err != nil {
			c.Error(err)
			return
		}

		availabilityList, err = models.GetAvailabilityByPsychologist(c, psychologistID)
		if err != nil {
			c.Error(err)
			return
		}

	} else {
		availabilityList, err = models.GetAllAvailabilities(c)
		if err != nil {
			c.Error(err)
			return
		}
	}

	if len(availabilityList) == 0 {
		availabilityList = []models.Availability{}
	}

	c.JSON(http.StatusOK, availabilityList)
}

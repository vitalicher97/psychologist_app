package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vitalicher97/psychologist_app/internal/app/db/models"
)

// CreatePsychologist handles the creation of a new psychologist.
func CreatePsychologist(c *gin.Context) {
	var psychologist models.Psychologist
	if err := c.ShouldBindJSON(&psychologist); err != nil {
		c.Error(err)
		return
	}

	if err := psychologist.Create(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, psychologist)
}

// GetPsychologist handles retrieving a psychologist by ID.
func GetPsychologist(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	psychologist := &models.Psychologist{ID: id}

	psychologist, err = psychologist.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, psychologist)
}

// UpdatePsychologist handles updating a psychologist's data by ID.
func UpdatePsychologist(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	existingPsychologist := &models.Psychologist{ID: id}
	existingPsychologist, err = existingPsychologist.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var psychologist models.Psychologist
	if err := c.ShouldBindJSON(&psychologist); err != nil {
		c.Error(err)
		return
	}

	psychologist.CreatedAt = existingPsychologist.CreatedAt
	psychologist.ID = id

	if err := psychologist.Update(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, psychologist)
}

// DeletePsychologist handles deleting a psychologist by ID.
func DeletePsychologist(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	psychologist := &models.Psychologist{ID: id}

	if err := psychologist.DeleteByID(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ListPsychologists handles retrieving a list of all psychologists.
func GetAllPsychologists(c *gin.Context) {

	psychologists, err := models.GetAllPsychologists(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, psychologists)
}

package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vitalicher97/psychologist_app/internal/app/db/models"
)

// CreateAppointment handles the creation of a new appointment.
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.Error(err)
		return
	}

	if err := appointment.Create(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

// GetAppointment handles retrieving an appointment by ID.
func GetAppointment(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	appointment := &models.Appointment{ID: id}

	appointment, err = appointment.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, appointment)
}

// UpdateAppointment handles updating an appointment's data by ID.
func UpdateAppointment(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	existingAppointment := &models.Appointment{ID: id}
	existingAppointment, err = existingAppointment.GetByID(c)
	if err != nil {
		c.Error(err)
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.Error(err)
		return
	}

	appointment.CreatedAt = existingAppointment.CreatedAt
	appointment.ID = id

	if err := appointment.Update(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, appointment)
}

// DeleteAppointment handles deleting an appointment by ID.
func DeleteAppointment(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Error(err)
		return
	}

	appointment := &models.Appointment{ID: id}

	if err := appointment.DeleteByID(c); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ListAppointments handles retrieving a list of all appointments.
func GetAllAppointments(c *gin.Context) {
	psychologistIDStr := c.Query("psychologist")

	var appointmentList []models.Appointment
	var err error
	if psychologistIDStr != "" {
		psychologistID, err := strconv.Atoi(psychologistIDStr)
		if err != nil {
			c.Error(err)
			return
		}

		appointmentList, err = models.GetAppointmentsByPsychologistID(c, psychologistID)
		if err != nil {
			c.Error(err)
			return
		}
	} else {
		appointmentList, err = models.GetAllAppointments(c)
		if err != nil {
			c.Error(err)
			return
		}
	}

	if len(appointmentList) == 0 {
		appointmentList = []models.Appointment{}
	}

	c.JSON(http.StatusOK, appointmentList)
}

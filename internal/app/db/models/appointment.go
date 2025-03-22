package models

import (
	"context"
	"time"

	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

// Appointment represents the appointments table in the database.
type Appointment struct {
	ID             int       `json:"id" binding:"-" pg:",pk"`
	PsychologistID int       `json:"psychologist_id" binding:"required" pg:",notnull"`
	CustomerID     int       `json:"customer_id" binding:"required" pg:",notnull"`
	StartTime      time.Time `json:"start_time" binding:"required" pg:",notnull"`
	EndTime        time.Time `json:"end_time" binding:"required" pg:",notnull"`
	Status         string    `json:"status" binding:"-" pg:",notnull"`
	CreatedBy      int       `json:"created_by" binding:"-" pg:",notnull"`
	UpdatedBy      int       `json:"updated_by" binding:"-" pg:",notnull"`
	CreatedAt      time.Time `json:"created_at" binding:"-" pg:",default:now()"`
	UpdatedAt      time.Time `json:"updated_at" binding:"-" pg:",default:now()"`
}

// BeforeInsert is a method for performing additional changes to the appointment table when INSERT query executes. It adds time in created_at and updated_at columns.
func (a *Appointment) BeforeInsert(ctx context.Context) (context.Context, error) {
	a.CreatedAt = time.Now()
	a.UpdatedAt = a.CreatedAt

	return ctx, nil
}

// BeforeUpdate is a method for performing additional changes to the appointment table when UPDATE query executes. It updates time in updated_at column.
func (a *Appointment) BeforeUpdate(ctx context.Context) (context.Context, error) {
	a.UpdatedAt = time.Now()

	return ctx, nil
}

// Create inserts a new appointment into the database.
func (a *Appointment) Create(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(a).Insert()

	return err
}

// GetByID retrieves an appointment by its ID.
func (a *Appointment) GetByID(ctx context.Context) (*Appointment, error) {
	conn := db.GetConnection()
	err := conn.WithContext(ctx).Model(a).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return a, nil
}

// GetAllAppointments retrieves a list of all appointments.
func GetAllAppointments(ctx context.Context) ([]Appointment, error) {
	conn := db.GetConnection()
	var appointments []Appointment
	err := conn.WithContext(ctx).Model(&appointments).Select()
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

// Update modifies an existing appointment's data.
func (a *Appointment) Update(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(a).WherePK().Update()

	return err
}

// DeleteByID removes an appointment from the database by its ID.
func (a *Appointment) DeleteByID(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(a).WherePK().Delete()

	return err
}

// GetAppointmentsByPsychologistID retrieves a list of all appointments for a given psychologist.
func GetAppointmentsByPsychologistID(ctx context.Context, psychologistID int) ([]Appointment, error) {
	conn := db.GetConnection()
	var appointments []Appointment
	err := conn.WithContext(ctx).Model(&appointments).Where("psychologist_id = ?", psychologistID).Select()
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

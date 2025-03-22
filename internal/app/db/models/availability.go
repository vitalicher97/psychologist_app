package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

// Availability represents the availability table in the database.
type Availability struct {
	ID             int          `json:"id" binding:"-" pg:",pk"`
	PsychologistID int          `json:"psychologist_id" binding:"required" pg:",notnull"`
	DayOfWeek      time.Weekday `json:"day_of_week" binding:"required" pg:",notnull"`
	StartTime      TimeOnly     `json:"start_time" binding:"required" pg:",notnull"`
	EndTime        TimeOnly     `json:"end_time" binding:"required" pg:",notnull"`
	CreatedBy      int          `json:"created_by" binding:"-" pg:",notnull"`
	UpdatedBy      int          `json:"updated_by" binding:"-" pg:",notnull"`
	CreatedAt      time.Time    `json:"created_at" binding:"-" pg:",default:now()"`
	UpdatedAt      time.Time    `json:"updated_at" binding:"-" pg:",default:now()"`
}

// Custom TimeOnly type to handle "TIME WITHOUT TIME ZONE"
type TimeOnly struct {
	time.Time
}

// Scan method for PostgreSQL compatibility
func (t *TimeOnly) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		return t.parseTime(v)
	case []byte:
		return t.parseTime(string(v))
	default:
		return fmt.Errorf("invalid time format: %T", value)
	}
}

// Value method for PostgreSQL INSERT/UPDATE
func (t TimeOnly) Value() (driver.Value, error) {
	return t.Format("15:04:05"), nil
}

// MarshalJSON ensures JSON output remains "HH:MM:SS"
func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format("15:04:05"))
}

// UnmarshalJSON allows parsing "HH:MM" or "HH:MM:SS"
func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	var strTime string
	if err := json.Unmarshal(data, &strTime); err != nil {
		return err
	}
	return t.parseTime(strTime)
}

// Helper function to parse time and handle missing seconds
func (t *TimeOnly) parseTime(strTime string) error {
	if len(strTime) == 5 {
		strTime += ":00"
	} else if len(strTime) != 8 {
		return fmt.Errorf("invalid time format: %s", strTime)
	}

	parsedTime, err := time.Parse("15:04:05", strTime)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

// BeforeInsert is a method for performing additional changes to the availability table when INSERT query executes. It adds time in created_at and updated_at columns.
func (a *Availability) BeforeInsert(ctx context.Context) (context.Context, error) {
	a.CreatedAt = time.Now()
	a.UpdatedAt = a.CreatedAt

	return ctx, nil
}

// BeforeUpdate is a method for performing additional changes to the availability table when UPDATE query executes. It updates the time in updated_at column.
func (a *Availability) BeforeUpdate(ctx context.Context) (context.Context, error) {
	a.UpdatedAt = time.Now()

	return ctx, nil
}

// List retrieves all availability records.
func GetAllAvailabilities(ctx context.Context) ([]Availability, error) {
	conn := db.GetConnection()
	var availabilities []Availability
	err := conn.WithContext(ctx).Model(&availabilities).Order("id").Select()
	if err != nil {
		return nil, err
	}

	return availabilities, nil
}

// GetByID retrieves an availability record by its ID.
func (a *Availability) GetByID(ctx context.Context) (*Availability, error) {
	conn := db.GetConnection()
	err := conn.WithContext(ctx).Model(a).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Create inserts a new availability record into the database.
func (a *Availability) Create(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(a).Insert()

	return err
}

// Update modifies an existing availability record's data.
func (a *Availability) Update(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(a).WherePK().Update()

	return err
}

// Delete removes an availability record from the database by its ID.
func (a *Availability) DeleteByID(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(a).WherePK().Delete()

	return err
}

// GetAvailabilityByPsychologist retrieves all availability records for a psychologist by their ID.
func GetAvailabilityByPsychologist(ctx context.Context, psychologistID int) ([]Availability, error) {
	conn := db.GetConnection()
	var availability []Availability
	err := conn.WithContext(ctx).Model(&availability).Where("psychologist_id = ?", psychologistID).OrderExpr("day_of_week, start_time").Select()
	if err != nil {
		return nil, err
	}

	return availability, nil
}

package models

import (
	"context"
	"time"

	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

// ConsultationPricing represents the consultation_pricing table in the database.
type ConsultationPricing struct {
	tableName struct{} `pg:"consultation_pricing"`

	ID             int       `json:"id" binding:"-" pg:",pk"`
	PsychologistID int       `json:"psychologist_id" binding:"required" pg:",notnull"`
	Price          float64   `json:"price" binding:"required" pg:",notnull"`
	Currency       string    `json:"currency" binding:"required" pg:",notnull"`
	CreatedBy      int       `json:"created_by" binding:"-" pg:",notnull"`
	UpdatedBy      int       `json:"updated_by" binding:"-" pg:",notnull"`
	CreatedAt      time.Time `json:"created_at" binding:"-" pg:",default:now()"`
	UpdatedAt      time.Time `json:"updated_at" binding:"-" pg:",default:now()"`
}

// BeforeInsert is a method for performing additional changes to the consultation_pricing table when INSERT query executes.
// It adds time to created_at and updated_at columns
func (c *ConsultationPricing) BeforeInsert(ctx context.Context) (context.Context, error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt

	return ctx, nil
}

// BeforeUpdate is a method for performing additional changes to the consultation_pricing table when UPDATE query executes.
// It updates time in updated_at column
func (c *ConsultationPricing) BeforeUpdate(ctx context.Context) (context.Context, error) {
	c.UpdatedAt = time.Now()

	return ctx, nil
}

// Create inserts a new consultation pricing entry into the database.
func (c *ConsultationPricing) Create(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(c).Insert()

	return err
}

// Update modifies an existing consultation pricing entry.
func (c *ConsultationPricing) Update(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(c).WherePK().Update()

	return err
}

// DeleteByID removes a consultation pricing entry from the database by its ID.
func (c *ConsultationPricing) DeleteByID(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(c).WherePK().Delete()

	return err
}

// GetByID retrieves a consultation pricing entry by its ID.
func (c *ConsultationPricing) GetByID(ctx context.Context) (*ConsultationPricing, error) {
	conn := db.GetConnection()
	err := conn.WithContext(ctx).Model(c).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetAll retrieves all consultation pricing entries.
func GetAllConsultationPricing(ctx context.Context) ([]ConsultationPricing, error) {
	conn := db.GetConnection()
	var pricingList []ConsultationPricing
	err := conn.WithContext(ctx).Model(&pricingList).Select()
	if err != nil {
		return nil, err
	}

	return pricingList, nil
}

// GetConsultationPricingByPsychologistID retrieves consultation pricing entries by psychologist ID.
func GetConsultationPricingByPsychologistID(ctx context.Context, psychologistID int) ([]ConsultationPricing, error) {
	conn := db.GetConnection()
	var pricingList []ConsultationPricing
	err := conn.WithContext(ctx).Model(&pricingList).Where("psychologist_id = ?", psychologistID).Select()
	if err != nil {
		return nil, err
	}

	return pricingList, nil
}

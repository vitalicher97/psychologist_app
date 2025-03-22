package models

import (
	"context"
	"time"

	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

// Psychologist represents the psychologists table in the database.
type Psychologist struct {
	ID             int       `json:"id" binding:"-" pg:",pk"`
	FirstName      string    `json:"first_name" binding:"required" pg:",notnull"`
	LastName       string    `json:"last_name" binding:"required" pg:",notnull"`
	Email          string    `json:"email" binding:"required,email" pg:",unique,notnull"`
	ProfilePicture string    `json:"profile_picture" binding:"-"`
	Bio            string    `json:"bio" binding:"-"`
	CreatedBy      int       `json:"created_by" binding:"-" pg:",notnull"`
	UpdatedBy      int       `json:"updated_by" binding:"-" pg:",notnull"`
	CreatedAt      time.Time `json:"created_at" binding:"-" pg:",default:now()"`
	UpdatedAt      time.Time `json:"updated_at" binding:"-" pg:",default:now()"`
}

// BeforeInsert is a method for performing additional changes to the psychologists table when INSERT query executes. It add time in created_at and updated_at column
func (p *Psychologist) BeforeInsert(ctx context.Context) (context.Context, error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt

	return ctx, nil
}

// BeforeUpdate is a method for performing additional changes to the psychologists table when UPDATE query executes. It updates time in updated_at column
func (p *Psychologist) BeforeUpdate(ctx context.Context) (context.Context, error) {
	p.UpdatedAt = time.Now()

	return ctx, nil
}

// List retrieves all psychologists.
func GetAllPsychologists(ctx context.Context) ([]Psychologist, error) {
	conn := db.GetConnection()
	var psychologists []Psychologist
	err := conn.WithContext(ctx).Model(&psychologists).Order("id").Select()
	if err != nil {
		return nil, err
	}

	return psychologists, nil
}

// GetByID retrieves a psychologist by their ID.
func (p *Psychologist) GetByID(ctx context.Context) (*Psychologist, error) {
	conn := db.GetConnection()
	err := conn.WithContext(ctx).Model(p).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Create inserts a new psychologist into the database.
func (p *Psychologist) Create(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(p).Returning("*").Insert()

	return err
}

// Update modifies an existing psychologist's data.
func (p *Psychologist) Update(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(p).WherePK().Update()

	return err
}

// Delete removes a psychologist from the database by their ID.
func (p *Psychologist) DeleteByID(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(p).WherePK().Delete()

	return err
}

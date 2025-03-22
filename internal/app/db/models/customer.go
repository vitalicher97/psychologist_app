package models

import (
	"context"
	"time"

	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

// Customer represents the customers table in the database.
type Customer struct {
	ID        int       `json:"id" binding:"-" pg:",pk"`
	FirstName string    `json:"first_name" binding:"required" pg:",notnull"`
	LastName  string    `json:"last_name" binding:"required" pg:",notnull"`
	Email     string    `json:"email" binding:"required,email" pg:",unique,notnull"`
	Phone     string    `json:"phone" binding:"-" pg:",notnull"`
	CreatedBy int       `json:"created_by" binding:"-" pg:",notnull"`
	UpdatedBy int       `json:"updated_by" binding:"-" pg:",notnull"`
	CreatedAt time.Time `json:"created_at" binding:"-" pg:",default:now()"`
	UpdatedAt time.Time `json:"updated_at" binding:"-" pg:",default:now()"`
}

// BeforeInsert is a method for performing additional changes to the customers table when an INSERT query executes.
// It adds time in created_at and updated_at columns
func (c *Customer) BeforeInsert(ctx context.Context) (context.Context, error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt

	return ctx, nil
}

// BeforeUpdate is a method for performing additional changes to the customers table when an UPDATE query executes.
// It updates the time in updated_at column
func (c *Customer) BeforeUpdate(ctx context.Context) (context.Context, error) {
	c.UpdatedAt = time.Now()

	return ctx, nil
}

// Create inserts a new customer into the database.
func (c *Customer) Create(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(c).Insert()

	return err
}

// GetByID retrieves a customer by their ID.
func (c *Customer) GetByID(ctx context.Context) (*Customer, error) {
	conn := db.GetConnection()
	err := conn.WithContext(ctx).Model(c).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetByEmail retrieves a customer by their email.
func GetCustomerByEmail(ctx context.Context, email string) (*Customer, error) {
	conn := db.GetConnection()
	var customer Customer
	err := conn.WithContext(ctx).Model(&customer).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// GetAll retrieves all customers.
func GetAllCustomers(ctx context.Context) ([]Customer, error) {
	conn := db.GetConnection()
	var customers []Customer
	err := conn.WithContext(ctx).Model(&customers).Select()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// Update modifies an existing customer's data.
func (c *Customer) Update(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(c).WherePK().Update()

	return err
}

// DeleteByID removes a customer from the database by their ID.
func (c *Customer) DeleteByID(ctx context.Context) error {
	conn := db.GetConnection()
	_, err := conn.WithContext(ctx).Model(c).WherePK().Delete()

	return err
}

package models

import (
	"time"

	"github.com/vitalicher97/psychologist_app/internal/app/db"
)

// CustomerPsychologistPrices represents the fixed price assigned to a customer for a psychologist.
type CustomerPsychologistPrices struct {
	ID             int       `json:"id" pg:",pk"`
	CustomerID     int       `json:"customer_id" pg:",notnull"`
	PsychologistID int       `json:"psychologist_id" pg:",notnull"`
	FixedPrice     float64   `json:"fixed_price" pg:",notnull,use_zero"`
	CreatedAt      time.Time `json:"created_at" pg:"default:now()"`
}

// GetFixedPrice retrieves the fixed price for a specific customer and psychologist.
func GetFixedPrice(customerID, psychologistID int) (*CustomerPsychologistPrices, error) {
	conn := db.GetConnection()

	priceRecord := &CustomerPsychologistPrices{}

	err := conn.Model(priceRecord).
		Where("customer_id = ? AND psychologist_id = ?", customerID, psychologistID).
		Select()

	if err != nil {
		return nil, err
	}

	return priceRecord, nil
}

// CreateFixedPrice inserts a new record if it does not exist.
func (cpp *CustomerPsychologistPrices) Create() error {
	conn := db.GetConnection()

	_, err := conn.Model(cpp).Insert()
	if err != nil {
		return err
	}

	return err
}

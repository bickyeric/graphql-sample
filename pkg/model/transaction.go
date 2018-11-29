package model

import (
	"database/sql"
	"time"

	"github.com/bickyeric/graphql-sample/pkg/wrapper/database"
)

// Transaction ...
type Transaction struct {
	ID            int
	ReceiptNumber string
	EmployeeID    int
	CustomerID    int
	Comment       string
	Status        string
	DiscountID    int
	TaxID         int
	CreatedAt     time.Time
	UpdatedAt     time.Time

	receiptNumber sql.NullString
	customerID    sql.NullInt64
	comment       sql.NullString
	discountID    sql.NullInt64
	taxID         sql.NullInt64
}

// Create ...
func (t Transaction) Create() (Transaction, error) {
	row, err := database.Connection.Exec(
		`INSERT INTO transaction(receipt_number, employee_id customer_id, comment, status, discount_id, tax_id)
		VALUES (?, ?, ?)`, NewNullString(t.ReceiptNumber), t.EmployeeID, NewNullInt64(t.CustomerID), NewNullString(t.Comment), t.Status, NewNullInt64(t.DiscountID), NewNullInt64(t.TaxID),
	)
	if err != nil {
		return t, err
	}
	id, _ := row.LastInsertId()
	t.ID = int(id)
	return t, nil
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TransactionStatus string

const (
	TransactionStatusPENDING   TransactionStatus = "PENDING"
	TransactionStatusPAID      TransactionStatus = "PAID"
	TransactionStatusFAILED    TransactionStatus = "FAILED"
	TransactionStatusCANCELLED TransactionStatus = "CANCELLED"
)

func (e *TransactionStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TransactionStatus(s)
	case string:
		*e = TransactionStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for TransactionStatus: %T", src)
	}
	return nil
}

type NullTransactionStatus struct {
	TransactionStatus TransactionStatus `json:"transaction_status"`
	Valid             bool              `json:"valid"` // Valid is true if TransactionStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTransactionStatus) Scan(value interface{}) error {
	if value == nil {
		ns.TransactionStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TransactionStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTransactionStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TransactionStatus), nil
}

type File struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	Size      int32     `json:"size"`
	Path      string    `json:"path"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Transaction struct {
	ID         uuid.UUID         `json:"id"`
	ExternalID string            `json:"external_id"`
	Amount     pgtype.Numeric    `json:"amount"`
	Date       time.Time         `json:"date"`
	Status     TransactionStatus `json:"status"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

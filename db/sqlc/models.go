// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"database/sql"

	"github.com/google/uuid"
)

type Affiliate struct {
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	MasterAffiliate uuid.NullUUID   `json:"master_affiliate"`
	Balance         string          `json:"balance"`
	Percent         sql.NullFloat64 `json:"percent"`
}

type Commission struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	AffiliateID uuid.UUID `json:"affiliate_id"`
	Amount      string    `json:"amount"`
}

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Quantity int32     `json:"quantity"`
	Price    string    `json:"price"`
}

type User struct {
	ID          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Balance     string        `json:"balance"`
	AffiliateID uuid.NullUUID `json:"affiliate_id"`
}

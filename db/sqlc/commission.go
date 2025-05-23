package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type CommissionRepository interface {
	WithTx(ctx context.Context, fn func(q *Queries) error) error
	ListCommissions(ctx context.Context) ([]Commission, error)
}

// --------- struct ที่ implement interface ---------

type Commissions struct {
	db *sql.DB  // เก็บ connection pool
	q  *Queries // sqlc-generated methods
}

func NewCommissions(db *sql.DB) *Commissions {
	return &Commissions{
		db: db,
		q:  New(db), // <- New(db) คือฟังก์ชัน sqlc สร้างให้ คืน *Queries
	}
}

func (p *Commissions) WithTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := p.q.WithTx(tx) // sqlc-generated helper

	if err := fn(qtx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (p *Commissions) ListCommissions(ctx context.Context) ([]Commission, error) {
	return p.q.ListCommission(ctx)
}

func (p *Commissions) GetCommission(ctx context.Context, id uuid.UUID) (Commission, error) {
	return p.q.GetCommission(ctx, id)
}

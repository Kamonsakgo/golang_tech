package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type AffiliateRepository interface {
	WithTx(ctx context.Context, fn func(q *Queries) error) error
	Createaffiliate(ctx context.Context, arg CreateaffiliateParams) (Affiliate, error)
	ListAffiliates(ctx context.Context) ([]Affiliate, error)
	GetAffiliate(ctx context.Context, id uuid.UUID) (Affiliate, error)
}

// --------- struct ที่ implement interface ---------

type Affiliates struct {
	db *sql.DB  // เก็บ connection pool
	q  *Queries // sqlc-generated methods
}

func NewAffiliates(db *sql.DB) *Affiliates {
	return &Affiliates{
		db: db,
		q:  New(db), // <- New(db) คือฟังก์ชัน sqlc สร้างให้ คืน *Queries
	}
}
func (p *Affiliates) ListAffiliates(ctx context.Context) ([]Affiliate, error) {
	return p.q.Listaffiliate(ctx)
}
func (p *Affiliates) WithTx(ctx context.Context, fn func(q *Queries) error) error {
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

func (p *Affiliates) Createaffiliate(ctx context.Context, arg CreateaffiliateParams) (Affiliate, error) {
	return p.q.Createaffiliate(ctx, arg)
}

func (p *Affiliates) Listaffiliate(ctx context.Context) ([]Affiliate, error) {
	return p.q.Listaffiliate(ctx)
}

func (p *Affiliates) Getaffiliate(ctx context.Context, id uuid.UUID) (Affiliate, error) {
	return p.q.Getaffiliate(ctx, id)
}

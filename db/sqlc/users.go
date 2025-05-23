package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// --------- interface ที่ service layer จะใช้ ---------

type UserRepository interface {
	CreateUser(ctx context.Context, arg CreateUsersParams) (User, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	ListUsers(ctx context.Context, limit, offset int32) ([]User, error)
	DeductBalance(ctx context.Context, id uuid.UUID, balance string) (User, error)
	AddBalanceParams(ctx context.Context, id uuid.UUID, balance string) (User, error)
	WithTx(ctx context.Context, fn func(q *Queries) error) error
}

// --------- struct ที่ implement interface ---------

type Users struct {
	db *sql.DB  // เก็บ connection pool
	q  *Queries // sqlc-generated methods
}

func NewUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
		q:  New(db), // <- New(db) คือฟังก์ชัน sqlc สร้างให้ คืน *Queries
	}
}

// --------- เมธอด passthrough ถึง *Queries ---------

func (u *Users) CreateUser(ctx context.Context, arg CreateUsersParams) (User, error) {
	// ถ้า AffiliateID.UUID เป็น uuid.Nil ให้ตั้ง Valid เป็น false เพื่อเก็บเป็น NULL ใน DB
	if arg.AffiliateID.UUID == uuid.Nil {
		arg.AffiliateID = uuid.NullUUID{
			UUID:  uuid.Nil,
			Valid: false,
		}
	} else {
		// ถ้า AffiliateID มีค่า ให้ตั้ง Valid เป็น true
		arg.AffiliateID.Valid = true
		u.q.Createaffiliate(ctx, CreateaffiliateParams{
			Name:            arg.Username,
			MasterAffiliate: arg.AffiliateID,
			Balance:         "0.00",
		})
	}

	// สร้าง user ตามข้อมูลที่ได้รับ
	user, err := u.q.CreateUsers(ctx, arg)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *Users) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	return u.q.GetUser(ctx, id)
}

func (u *Users) ListUsers(ctx context.Context, limit, offset int32) ([]User, int64, error) {
	// ดึงข้อมูลผู้ใช้
	users, err := u.q.ListUsers(ctx, ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, 0, err
	}

	// ดึงจำนวนทั้งหมด
	totalCount, err := u.q.CountUsers(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (u *Users) DeductBalance(ctx context.Context, id uuid.UUID, balance float64) (User, error) {
	balanceStr := fmt.Sprintf("%.2f", balance)
	return u.q.DeductBalance(ctx, DeductBalanceParams{
		ID:      id,
		Balance: balanceStr,
	})
}
func (u *Users) AddBalanceParams(ctx context.Context, id uuid.UUID, balance float64) (User, error) {
	balanceStr := fmt.Sprintf("%.2f", balance)
	return u.q.AddBalance(ctx, AddBalanceParams{
		ID:      id,
		Balance: balanceStr,
	})
}

// --------- helper รันหลายคำสั่งในหนึ่ง transaction ---------

func (u *Users) WithTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := u.q.WithTx(tx) // sqlc-generated helper

	if err := fn(qtx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

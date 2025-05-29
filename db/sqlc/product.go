package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	WithTx(ctx context.Context, fn func(q *Queries) error) error
	ListProducts(ctx context.Context) ([]Product, error)
	BuyProduct(ctx context.Context, data Product_order) (Product_order_Output, error)
}

// --------- struct ที่ implement interface ---------

type Products struct {
	db         *sql.DB  // เก็บ connection pool
	q          *Queries // sqlc-generated methods
	users      *Users
	affiliate  *Affiliates
	commission *Commissions
}

func NewProducts(db *sql.DB, users *Users, affiliate *Affiliates, commission *Commissions) *Products {
	return &Products{
		db:         db,
		q:          New(db),
		users:      users, // <- New(db) คือฟังก์ชัน sqlc สร้างให้ คืน *Queries
		affiliate:  affiliate,
		commission: commission,
	}
}

func (p *Products) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	return p.q.CreateProduct(ctx, arg)
}
func (p *Products) ListProducts(ctx context.Context) ([]Product, error) {
	return p.q.ListProduct(ctx)
}
func (p *Products) WithTx(ctx context.Context, fn func(q *Queries) error) error {
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
func (p *Products) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	return p.q.GetProduct(ctx, id)
}

type Product_order struct {
	User_id    uuid.UUID
	Name       string
	Product_ID uuid.UUID
	Amount     string
	Price      string
}
type Product_order_Output struct {
	User_id    uuid.UUID
	Name       string
	Product_ID uuid.UUID
	Amount     string
	Price      string
}

func (p *Products) BuyProduct(ctx context.Context, data Product_order) (Product_order_Output, error) {
	data_quantity, err := p.q.GetProduct_quantity(ctx, data.Product_ID)
	if err != nil {
		return Product_order_Output{}, fmt.Errorf("product not found")
	}
	user_balance, err := p.users.q.GetBalance(ctx, data.User_id)
	if err != nil {
		return Product_order_Output{}, fmt.Errorf("user not found")
	}
	userFloat, _ := strconv.ParseFloat(user_balance, 64)
	priceFloat, _ := strconv.ParseFloat(data.Price, 64)

	Quantity_int, _ := strconv.Atoi(data.Amount)

	if userFloat < priceFloat {
		return Product_order_Output{}, fmt.Errorf("not enough balance")
	}
	if data_quantity < int32(Quantity_int) {
		return Product_order_Output{}, fmt.Errorf("not enough quantity")
	}
	err = p.q.DeleteProduct(ctx, DeleteProductParams{
		ID:       data.Product_ID,
		Quantity: int32(Quantity_int),
	})

	_, err = p.q.DeductBalance(ctx, DeductBalanceParams{
		ID:      data.User_id,
		Balance: data.Price,
	})

	user, err := p.q.GetUser(ctx, data.User_id)

	if user.AffiliateID.Valid {
		// ดึง chain เดิม
		affiliateChain, err := p.affiliate.q.GetAffiliateChain(ctx, user.AffiliateID.UUID)
		if err != nil {
			return Product_order_Output{}, err
		}
		// useraff, err := p.affiliate.q.GetaffiliateByname(ctx, user.Username)
		// if err != nil {
		// 	return Product_order_Output{}, err
		// }
		// affiliateChain = append(
		// 	[]GetAffiliateChainRow{{ID: useraff.ID, MasterAffiliate: useraff.MasterAffiliate, Balance: useraff.Balance}},
		// 	affiliateChain...,
		// )

		productPrice, err := strconv.ParseFloat(data.Price, 64)
		if err != nil {
			return Product_order_Output{}, err
		}

		for i := 0; i < len(affiliateChain); i++ {
			fmt.Println(affiliateChain[i].ID)
			commissionToAdd := productPrice * affiliateChain[i].Percent.Float64
			commissionToAddstr := strconv.FormatFloat(commissionToAdd, 'f', -1, 64)

			err = p.q.AddBalance_affiliate(ctx, AddBalance_affiliateParams{
				ID:      affiliateChain[i].ID,
				Balance: commissionToAddstr,
			})
			if err != nil {
				return Product_order_Output{}, err
			}

			_, err = p.q.Createcommission(ctx, CreatecommissionParams{
				AffiliateID: affiliateChain[i].ID,
				OrderID:     data.Product_ID,
				Amount:      commissionToAddstr,
			})
			if err != nil {
				return Product_order_Output{}, err
			}
		}
	}

	return Product_order_Output{
		User_id:    data.User_id,
		Name:       data.Name,
		Product_ID: data.Product_ID,
		Amount:     data.Amount,
		Price:      data.Price,
	}, err
}

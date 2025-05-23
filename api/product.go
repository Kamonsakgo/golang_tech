package api

import (
	"fmt"
	"net/http"
	db "simple/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name    string  `json:"name"`
	Quntity int     `json:"quantity"`
	Price   float64 `json:"price"`
}

func (server *server) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	Pricestr := fmt.Sprintf("%.2f", req.Price)
	arg := db.CreateProductParams{
		Name:     req.Name,
		Quantity: int32(req.Quntity),
		Price:    Pricestr,
	}
	product, err := server.products.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (server *server) ListProduct(ctx *gin.Context) {

	products, err := server.products.ListProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, products)
}
func (server *server) GetProduct(ctx *gin.Context) {

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	product, err := server.products.GetProduct(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (server *server) BuyProduct(ctx *gin.Context) {
	userIDStr := ctx.PostForm("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productIDStr := ctx.PostForm("product_id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	price := ctx.PostForm("price")
	amount := ctx.PostForm("amount")

	data := db.Product_order{ // ต้อง Export struct ใน db: ProductOrder
		User_id:    userID,
		Product_ID: productID,
		Name:       ctx.PostForm("name"),
		Amount:     amount,
		Price:      price,
	}

	product, err := server.products.BuyProduct(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

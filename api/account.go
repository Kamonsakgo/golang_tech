package api

import (
	"fmt"
	"net/http"
	db "simple/db/sqlc"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUsersRequest struct {
	Username    string        `json:"username"`
	AffiliateID uuid.NullUUID `json:"affiliate_id"`
}
type PaginatedResponse[T any] struct {
	Page       int32 `json:"page"`
	TotalPage  int32 `json:"total_page"`
	Count      int32 `json:"count"`
	TotalCount int32 `json:"total_count"`
	Data       []T   `json:"data"`
}

func (server *server) createuser(ctx *gin.Context) {
	var req CreateUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUsersParams{
		Username:    req.Username,
		Balance:     "0",
		AffiliateID: req.AffiliateID,
	}
	user, err := server.users.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (server *server) ListUsers(ctx *gin.Context) {
	// รับ query param
	limitStr := ctx.Query("limit")
	pageStr := ctx.Query("page")

	// แปลงเป็น int
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	// คำนวณ offset
	offset := (page - 1) * limit

	// ดึงข้อมูล
	users, totalCount, err := server.users.ListUsers(ctx, int32(limit), int32(offset))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	totalPage := int32((totalCount + int64(limit) - 1) / int64(limit))
	response := PaginatedResponse[db.User]{
		Page:       int32(page),
		TotalPage:  totalPage,
		Count:      int32(len(users)),
		TotalCount: int32(totalCount),
		Data:       users,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *server) GetUser(ctx *gin.Context) {

	id, err := uuid.Parse(ctx.Param("id"))
	fmt.Println(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.users.GetUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type UpdatebalanceParams struct {
	ID      uuid.UUID `json:"id"`
	Balance float64   `json:"balance"`
}

func (server *server) deductBalance(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req UpdatebalanceParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.users.DeductBalance(ctx, id, req.Balance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	type UserResponse struct {
		ID      uuid.UUID `json:"id"`
		Balance string    `json:"balance"`
	}

	ctx.JSON(http.StatusOK, UserResponse{
		ID:      user.ID,
		Balance: user.Balance,
	})
}
func (server *server) addBalance(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var req UpdatebalanceParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.users.AddBalanceParams(ctx, id, req.Balance)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	type UserResponse struct {
		ID      uuid.UUID `json:"id"`
		Balance string    `json:"balance"`
	}

	ctx.JSON(http.StatusOK, UserResponse{
		ID:      user.ID,
		Balance: user.Balance,
	})
}

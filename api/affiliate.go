package api

import (
	"net/http"
	db "simple/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateaffiliateRequest struct {
	Name            string        `json:"name"`
	MasterAffiliate uuid.NullUUID `json:"master_affiliate"`
}

func (server *server) CreateAffiliate(ctx *gin.Context) {
	var req CreateaffiliateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var masterAff uuid.NullUUID
	if req.MasterAffiliate.UUID == uuid.Nil {
		masterAff = uuid.NullUUID{Valid: false}
	} else {
		masterAff = uuid.NullUUID{
			UUID:  req.MasterAffiliate.UUID,
			Valid: true,
		}
	}

	arg := db.CreateaffiliateParams{
		Name:            req.Name,
		MasterAffiliate: masterAff,
		Balance:         "0.00",
	}

	affiliate, err := server.Affiliates.Createaffiliate(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, affiliate)
}

func (server *server) ListAffiliates(ctx *gin.Context) {
	affiliates, err := server.Affiliates.Listaffiliate(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, affiliates)
}
func (server *server) GetAffiliate(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	affiliate, err := server.Affiliates.Getaffiliate(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, affiliate)
}

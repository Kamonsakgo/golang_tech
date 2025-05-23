package api

import (
	db "simple/db/sqlc"

	"github.com/gin-gonic/gin"
)

type server struct {
	users        *db.Users
	router       *gin.Engine
	products     *db.Products
	Affiliates   *db.Affiliates
	commisstions *db.Commissions
}

func NewServer(users *db.Users, products *db.Products, Affiliates *db.Affiliates, commisstions *db.Commissions) *server {
	server := &server{
		users:        users,
		products:     products,
		Affiliates:   Affiliates,
		commisstions: commisstions,
	}
	router := gin.Default()

	router.POST("/createuser", server.createuser)
	router.GET("/user/all", server.ListUsers)
	router.GET("/user/:id", server.GetUser)
	router.PATCH("/user/deduct/balance/:id", server.deductBalance)
	router.PATCH("/user/add/balance/:id", server.addBalance)
	router.POST("/product", server.CreateProduct)
	router.GET("/product/list", server.ListProduct)
	router.GET("/product/:id", server.GetProduct)
	router.POST("/affiliate", server.CreateAffiliate)
	router.GET("/affiliate/list", server.ListAffiliates)
	router.GET("/affiliate/:id", server.GetAffiliate)
	router.GET("/commission/:id", server.Getcommission)
	router.GET("/commission/list", server.Listcommission)
	router.POST("/buyproduct", server.BuyProduct)
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

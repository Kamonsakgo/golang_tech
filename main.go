package main

import (
	"database/sql"
	"log"
	"simple/api"
	db "simple/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:1234@localhost:5432/simple?sslmode=disable"
	ServerAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	users := db.NewUsers(conn)

	affiliate := db.NewAffiliates(conn)

	commissions := db.NewCommissions(conn)
	products := db.NewProducts(conn, users, affiliate, commissions)
	server := api.NewServer(users, products, affiliate, commissions)
	err = server.Start(ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

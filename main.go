package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vsevicky/simplebank/api"
	db "github.com/vsevicky/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to the databse", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}

}

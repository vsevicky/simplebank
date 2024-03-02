package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vsevicky/simplebank/api"
	db "github.com/vsevicky/simplebank/db/sqlc"
	"github.com/vsevicky/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load Configuration")
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the databse", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}

}

package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vsevicky/simplebank/util"
)

var testQueries Store

// func newTestServer(t *testing.T, store db.Store) *Server {
// 	config := util.Config{
// 		TokenSymmetricKey:   util.RandomString(32),
// 		AccessTokenDuration: time.Minute,
// 	}

// 	server, err := NewServer(config, store)
// 	require.NoError(t, err)

// 	return server
// }

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load Configuration")
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the databse", err)
	}

	testQueries = NewStore(connPool)

	os.Exit(m.Run())
}

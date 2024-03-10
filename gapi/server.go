package gapi

import (
	"fmt"

	db "github.com/vsevicky/simplebank/db/sqlc"
	"github.com/vsevicky/simplebank/pb"
	"github.com/vsevicky/simplebank/token"
	"github.com/vsevicky/simplebank/util"
	// "github.com/vsevicky/simplebank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	// taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		// taskDistributor: taskDistributor,
	}

	return server, nil
}

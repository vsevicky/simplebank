package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vsevicky/simplebank/api"
	db "github.com/vsevicky/simplebank/db/sqlc"
	"github.com/vsevicky/simplebank/gapi"
	"github.com/vsevicky/simplebank/pb"
	"github.com/vsevicky/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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
	go runGatewayServer(config, store)
	runGrpcServer(config, store)

}

func runGatewayServer(
	// ctx context.Context,
	// waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,
	// taskDistributor worker.TaskDistributor,
) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	// waitGroup.Go(func() error {
	log.Printf("start Htttp server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("gRPC server failed to serve")
	}
}

func runGrpcServer(
	// ctx context.Context,
	// waitGroup *errgroup.Group,
	config util.Config,
	store db.Store,
	// taskDistributor worker.TaskDistributor,
) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	// gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	// waitGroup.Go(func() error {
	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("gRPC server failed to serve")
		// return err
	}

	// 	return nil
	// })

	// waitGroup.Go(func() error {
	// 	<-ctx.Done()
	// 	log.Info("graceful shutdown gRPC server")

	// 	grpcServer.GracefulStop()
	// 	log.Info("gRPC server is stopped")

	// 	return nil
	// })
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot Create the server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}

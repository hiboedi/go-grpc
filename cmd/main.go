package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/hiboedi/go-grpc/cmd/config"
	"github.com/hiboedi/go-grpc/cmd/services"
	productPb "github.com/hiboedi/go-grpc/pb/product"
)

const (
	port = ":50051"
)

func main() {

	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", err.Error())
	}

	db := config.ConnectDB()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server start on %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatal("failed to serve", err.Error())
	}
}

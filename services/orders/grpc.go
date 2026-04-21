package main

import (
	"log"
	"net"

	handler "github.com/Nihal1203/golang-grpc/services/orders/handler/orders"
	"github.com/Nihal1203/golang-grpc/services/orders/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGRPCServer(addr string) *grpcServer {
	return &grpcServer{addr}
}

func (s *grpcServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	grpcServer := grpc.NewServer()

	//register grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcServer, orderService)

	log.Println("Starting grpc server on port", s.addr)
	return grpcServer.Serve(lis)
}

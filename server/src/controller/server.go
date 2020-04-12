package controller

import (
	"log"
	"net"

	svc "server/src/controller/service"
	pb "server/src/generated/services"

	"google.golang.org/grpc"
)

// StartServer start server
func StartServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterHelloServicesServer(s, &svc.HelloService{})

	log.Println("Starting server", port)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal("Failed to open", port)
	}
}

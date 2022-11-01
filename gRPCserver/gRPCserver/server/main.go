package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCserver/gRPCserver/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type EchoService struct {
	pb.EchoServiceServer
}

func NewEchoService() *EchoService {
	return &EchoService{}
}
func (e *EchoService) Echo(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println("Echo server function called")

	return &pb.EchoResponse{
		Response: fmt.Sprintf("%s-pong\n", req.Request),
	}, nil
}

const (
	addr = "127.0.0.1:55051"
)

func main() {

	if lstnr, err := net.Listen("tcp", addr); err != nil {
		log.Fatalf("Error listening to port %s \n %+v \n",
			addr,
			err)
	} else {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic error from %+v \n", r)
			}
		}()

		log.Printf("gRPCServer listening at port %s", addr)
		echoService := NewEchoService()
		s := grpc.NewServer()
		pb.RegisterEchoServiceServer(s, echoService)

		if err = s.Serve(lstnr); err != nil {
			log.Fatalf("Error listening at gprc server\n %s\n",
				err.Error())
		}
	}
}

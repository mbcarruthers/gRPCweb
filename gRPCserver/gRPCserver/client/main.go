package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCserver/gRPCserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	addr = "localhost:50051"
)

func echoRequest(client pb.EchoServiceClient, req *pb.EchoRequest) {
	log.Println("EchoRequest(client-side) called")
	res, err := client.Echo(context.Background(), req)
	if err != nil {
		log.Printf("%+v \n", err)
		return
	}
	log.Printf("Go client code seems to be working \n %s\n",
		res.Response)
}

func main() {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %s\n %s\n",
			addr,
			err.Error())
	}
	defer func() {
		if r := recover(); r != nil {
			log.Printf("$+v \n", err)
		}
	}()
	echoClient := pb.NewEchoServiceClient(conn)
	echoRequest(echoClient, &pb.EchoRequest{
		Request: fmt.Sprintf("Ping"),
	})
	defer conn.Close()

}

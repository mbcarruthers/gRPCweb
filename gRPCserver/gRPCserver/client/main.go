package main

import (
	"context"
	"fmt"
	pb "github.com/mbcarruthers/gRPCserver/gRPCserver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)
// NOTE: This code is only needed to test that the gRPCserver/server is working with solely gRPC, not grpc-web. Compilation of this code is not 
// needed to run the gRPCserver project at all. Just kept it here for server verification in case you need to see if it's actually working standalone.
//if you feel like you need it you should probobly just restart envoy container with `docker start -a ${CONTAINER_ID} instead so you can see what envoy is
//doing internally.
const (
	addr = "127.0.0.1:55051"
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

package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"seeker/grpcdemo/pb"
	"time"
)

const (
	address = "localhost:50051"
	defaultName = "hapi"
)

func main()  {
	conn , err :=grpc.Dial(address ,  grpc.WithInsecure())
	if err !=nil{
		fmt.Println("did not connect")
	}
	defer conn.Close()
	c :=pb.NewGreeterClient(conn)

	name :=defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx , cancel :=context.WithTimeout(context.Background() , time.Second)
	defer cancel()
	r ,err :=c.SayHello(ctx , &pb.HelloRequest{Name : name})
	if err !=nil {
		fmt.Println("34")
	}
	fmt.Println(r.Message)

}

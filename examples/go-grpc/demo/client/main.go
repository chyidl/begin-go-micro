package main

import (
	"context"
	"github.com/chyidl/begin-go-micro/examples/go-grpc/demo/proto/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close conn: %s", err)
		}
	}(conn)

	c := user.NewUserClient(conn)

	// Contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUserInfo(ctx, &user.UserID{
		ID: 9527,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("ID: %d, Name: %s, Age: %d, Gender: %s", r.GetID(), r.GetName(), r.GetAge(), r.GetGender())
}

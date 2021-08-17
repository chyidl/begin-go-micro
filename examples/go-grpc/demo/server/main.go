package main

import (
	"context"
	"github.com/chyidl/begin-go-micro/examples/go-grpc/demo/proto/article"
	"github.com/chyidl/begin-go-micro/examples/go-grpc/demo/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement
type server struct {
	user.UnimplementedUserServer
}

// GetUserInfo implements UserServer
func (s *server)GetUserInfo(_ context.Context, in *user.UserID) (*user.UserInfo, error) {
	log.Printf("Received: %v", in.GetID())
	return &user.UserInfo{
		ID: in.GetID(),
		Name: "chyidl",
		Age: 28,
		Gender: user.UserInfo_MALE,
	}, nil
}


func (s *server)GetUserFavArticle(_ context.Context, in *user.UserID) (*article.Articles_Article, error) {
	log.Printf("Received: %v", in.GetID())
	return &article.Articles_Article{
		ID:    in.GetID(),
		Title: "do or do not there is no try",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf(
			"failed to listen: %v",
			err,
		)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"fmt"
	"github.com/DenisCom3/m-auth/internal/config"
	pb "github.com/DenisCom3/m-auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

type server struct {
	pb.UnimplementedUserV1Server
}

func (s *server) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	t := time.Now()
	return &pb.GetResponse{
		Id:        1,
		Name:      "sdsf",
		Email:     "email",
		Role:      1,
		CreatedAt: timestamppb.New(t),
		UpdatedAt: timestamppb.New(t),
	}, nil
}
func main() {
	err := config.MustLoad()
	if err != nil {
		log.Fatalf("failed to init config. %v", err)
	}
	fmt.Println(config.GetPostgres().Dsn())

	lis, err := net.Listen("tcp", ":4300")

	if err != nil {
		log.Fatalf("failed to listen. %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUserV1Server(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

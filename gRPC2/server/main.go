package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listner, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err.Error())
	}
	srv := grpc.NewServer()
	proto.RegisterIsAvailableServer(srv, server{})
	reflection.Register(srv)
	if e := srv.Serve(listner); e != nil {
		panic(e)
	}

}

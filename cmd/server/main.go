package main

import (
	"net"
	"log"

	"google.golang.org/grpc"

	pd "github.com/NeuraSoul/chat-service/pkg/grpc"
	"github.com/NeuraSoul/chat-service/internal/domain/services"
)


func main() {
	port := ":8080"

	listner,err := net.Listen("tcp",port)
	if err != nil {
		log.Panic("fail to serve in %v",port)
	}

	s := grpc.NewServer()
	pd.RegisterChatServiceServer(s,services.NewChatService())

	go func () {
		log.Printf("serv in port %v",port)
		s.Serve(listner)
	}()



}
package services


import (

	pd "github.com/NeuraSoul/chat-service/pkg/grpc"
	"context"

)

type chatService struct {
	pd.UnimplementedChatServiceServer
}

func NewChatService() *chatService {
	return &chatService{}
}

func (s *chatService) SendMessage (ctx context.Context, req *pd.SendMessageRequest) (*pd.SendMessageResponse,error) {

	return &pd.SendMessageResponse{},nil
}
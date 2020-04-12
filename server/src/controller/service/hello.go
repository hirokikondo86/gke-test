package service

import (
	"context"
	"server/src/generated/enums"
	msg "server/src/generated/messages"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HelloService is gRPC service
type HelloService struct{}

// Hello return text
func (s *HelloService) Hello(ctx context.Context, in *msg.HelloRequest) (*msg.HelloReply, error) {

	return &msg.HelloReply{
		Status:     true,
		StatusCode: enums.StatusCodes_SUCCESS,
		Text:       "Hello World!!",
	}, status.Error(codes.OK, "")
}

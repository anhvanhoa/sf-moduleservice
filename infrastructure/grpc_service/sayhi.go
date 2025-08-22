package grpcservice

import (
	"context"

	proto "github.com/anhvanhoa/module-service/proto/gen/exam/v1"
)

func (s *examService) SayHi(ctx context.Context, req *proto.SayHiRequest) (*proto.SayHiResponse, error) {
	return &proto.SayHiResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

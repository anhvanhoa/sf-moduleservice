package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) CountByUserID(ctx context.Context, req *proto_user_role.CountByUserIDRequest) (*proto_user_role.CountByUserIDResponse, error) {
	count, err := s.userRoleUsecase.CountByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.CountByUserIDResponse{
		Count: count,
	}, nil
}

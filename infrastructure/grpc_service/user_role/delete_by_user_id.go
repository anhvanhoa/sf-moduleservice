package user_role_service

import (
	"context"

	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
)

func (s *userRoleService) DeleteByUserID(ctx context.Context, req *proto_user_role.DeleteByUserIDRequest) (*proto_user_role.DeleteByUserIDResponse, error) {
	err := s.userRoleUsecase.DeleteByUserID(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto_user_role.DeleteByUserIDResponse{}, nil
}

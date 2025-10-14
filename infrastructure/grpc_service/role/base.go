package role_service

import (
	"context"
	"module-service/domain/entity"
	role "module-service/domain/usecase/role"
	"module-service/infrastructure/repo"

	proto_role "github.com/anhvanhoa/sf-proto/gen/role/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type roleService struct {
	proto_role.UnimplementedRoleServiceServer
	roleUsecase role.RoleUsecaseI
}

func NewRoleServer(repos repo.Repositories) proto_role.RoleServiceServer {
	roleRepo := repos.RoleRepository()
	roleUC := role.NewRoleUsecase(roleRepo)
	return &roleService{
		roleUsecase: roleUC,
	}
}

func (s *roleService) GetAllRoles(ctx context.Context, req *proto_role.GetAllRolesRequest) (*proto_role.GetAllRolesResponse, error) {
	roles, err := s.roleUsecase.GetAllRoles()
	if err != nil {
		return nil, err
	}
	return &proto_role.GetAllRolesResponse{
		Roles: s.createProtoRoles(roles),
	}, nil
}

func (s *roleService) GetRoleById(ctx context.Context, req *proto_role.GetRoleByIdRequest) (*proto_role.GetRoleByIdResponse, error) {
	role, err := s.roleUsecase.GetRoleById(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_role.GetRoleByIdResponse{
		Role: s.createProtoRole(role),
	}, nil
}

func (s *roleService) CreateRole(ctx context.Context, req *proto_role.CreateRoleRequest) (*proto_role.CreateRoleResponse, error) {
	role := s.createEntityRole(req.Name, req.Description, req.Variant)
	err := s.roleUsecase.CreateRole(role)
	if err != nil {
		return nil, err
	}
	return &proto_role.CreateRoleResponse{
		Message: "Tạo vai trò thành công",
		Success: true,
	}, nil
}

func (s *roleService) UpdateRole(ctx context.Context, req *proto_role.UpdateRoleRequest) (*proto_role.UpdateRoleResponse, error) {
	role := s.createEntityRole(req.Name, req.Description, req.Variant)
	updatedRole, err := s.roleUsecase.UpdateRole(req.Id, role)
	if err != nil {
		return nil, err
	}
	return &proto_role.UpdateRoleResponse{
		Role: s.createProtoRole(updatedRole),
	}, nil
}

func (s *roleService) DeleteRole(ctx context.Context, req *proto_role.DeleteRoleRequest) (*proto_role.DeleteRoleResponse, error) {

	err := s.roleUsecase.DeleteRole(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto_role.DeleteRoleResponse{
		Message: "Xoá vai trò thành công",
		Success: true,
	}, nil
}

func (s *roleService) createEntityRole(
	name string,
	description string,
	variant string,
) entity.Role {
	return entity.Role{
		Name:        name,
		Description: description,
		Variant:     variant,
	}
}

func (s *roleService) createProtoRoles(roles []entity.Role) []*proto_role.Role {
	protoRoles := make([]*proto_role.Role, len(roles))
	for i, role := range roles {
		protoRoles[i] = s.createProtoRole(role)
	}
	return protoRoles
}

func (s *roleService) createProtoRole(role entity.Role) *proto_role.Role {
	r := &proto_role.Role{
		Id:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Variant:     role.Variant,
		Status:      string(role.Status),
		CreatedBy:   role.CreatedBy,
		CreatedAt:   timestamppb.New(role.CreatedAt),
	}

	if role.UpdatedAt != nil {
		r.UpdatedAt = timestamppb.New(*role.UpdatedAt)
	}

	return r
}

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/anhvanhoa/service-core/domain/cache"
	"github.com/anhvanhoa/service-core/domain/user_context"
	proto_permission "github.com/anhvanhoa/sf-proto/gen/permission/v1"
	proto_resource_permission "github.com/anhvanhoa/sf-proto/gen/resource_permission/v1"
	proto_role "github.com/anhvanhoa/sf-proto/gen/role/v1"
	proto_role_permission "github.com/anhvanhoa/sf-proto/gen/role_permission/v1"
	proto_user_role "github.com/anhvanhoa/sf-proto/gen/user_role/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yaml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

type RoleServiceClient struct {
	ctx                      context.Context
	roleClient               proto_role.RoleServiceClient
	permissionClient         proto_permission.PermissionServiceClient
	rolePermissionClient     proto_role_permission.RolePermissionServiceClient
	resourcePermissionClient proto_resource_permission.ResourcePermissionServiceClient
	userRoleClient           proto_user_role.UserRoleServiceClient
	conn                     *grpc.ClientConn
}

func NewRoleServiceClient(address string) (*RoleServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	ctx := context.Background()
	md := metadata.Pairs("user_id", "1")
	ctx = metadata.NewOutgoingContext(ctx, md)

	return &RoleServiceClient{
		roleClient:               proto_role.NewRoleServiceClient(conn),
		permissionClient:         proto_permission.NewPermissionServiceClient(conn),
		rolePermissionClient:     proto_role_permission.NewRolePermissionServiceClient(conn),
		resourcePermissionClient: proto_resource_permission.NewResourcePermissionServiceClient(conn),
		userRoleClient:           proto_user_role.NewUserRoleServiceClient(conn),
		conn:                     conn,
		ctx:                      ctx,
	}, nil
}

func (c *RoleServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Role Service Tests ==================

func (c *RoleServiceClient) TestCreateRole() {
	fmt.Println("\n=== Test Create Role ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Enter variant: ")
	variant, _ := reader.ReadString('\n')
	variant = cleanInput(variant)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.roleClient.CreateRole(ctx, &proto_role.CreateRoleRequest{
		Name:        name,
		Description: description,
		Variant:     variant,
	})
	if err != nil {
		fmt.Printf("Error calling CreateRole: %v\n", err)
		return
	}

	fmt.Printf("Create Role result:\n")
	fmt.Printf("Message: %s\n", resp.Message)
	fmt.Printf("Success: %t\n", resp.Success)
}

func (c *RoleServiceClient) TestGetRoleById() {
	fmt.Println("\n=== Test Get Role By ID ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.roleClient.GetRoleById(ctx, &proto_role.GetRoleByIdRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetRoleById: %v\n", err)
		return
	}

	fmt.Printf("Get Role result:\n")
	fmt.Printf("ID: %s\n", resp.Role.Id)
	fmt.Printf("Name: %s\n", resp.Role.Name)
	fmt.Printf("Description: %s\n", resp.Role.Description)
	fmt.Printf("Variant: %s\n", resp.Role.Variant)
	fmt.Printf("Status: %s\n", resp.Role.Status)
}

func (c *RoleServiceClient) TestGetAllRoles() {
	fmt.Println("\n=== Test Get All Roles ===")

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.roleClient.GetAllRoles(ctx, &proto_role.GetAllRolesRequest{})
	if err != nil {
		fmt.Printf("Error calling GetAllRoles: %v\n", err)
		return
	}

	fmt.Printf("Get All Roles result:\n")
	fmt.Printf("Total Roles: %d\n", len(resp.Roles))
	for i, role := range resp.Roles {
		fmt.Printf("  [%d] ID: %s, Name: %s, Description: %s, Variant: %s\n", i+1, role.Id, role.Name, role.Description, role.Variant)
	}
}

func (c *RoleServiceClient) TestUpdateRole() {
	fmt.Println("\n=== Test Update Role ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter new description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Enter new variant: ")
	variant, _ := reader.ReadString('\n')
	variant = cleanInput(variant)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.roleClient.UpdateRole(ctx, &proto_role.UpdateRoleRequest{
		Id:          id,
		Name:        name,
		Description: description,
		Variant:     variant,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateRole: %v\n", err)
		return
	}

	fmt.Printf("Update Role result:\n")
	fmt.Printf("ID: %s\n", resp.Role.Id)
	fmt.Printf("Name: %s\n", resp.Role.Name)
	fmt.Printf("Description: %s\n", resp.Role.Description)
	fmt.Printf("Variant: %s\n", resp.Role.Variant)
}

func (c *RoleServiceClient) TestDeleteRole() {
	fmt.Println("\n=== Test Delete Role ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.roleClient.DeleteRole(ctx, &proto_role.DeleteRoleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteRole: %v\n", err)
		return
	}

	fmt.Printf("Delete Role result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
}

// ================== Permission Service Tests ==================

func (c *RoleServiceClient) TestCreatePermission() {
	fmt.Println("\n=== Test Create Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter resource: ")
	resource, _ := reader.ReadString('\n')
	resource = cleanInput(resource)

	fmt.Print("Enter action: ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.permissionClient.CreatePermission(ctx, &proto_permission.CreatePermissionRequest{
		Resource:    resource,
		Action:      action,
		Description: description,
	})
	if err != nil {
		fmt.Printf("Error calling CreatePermission: %v\n", err)
		return
	}

	fmt.Printf("Create Permission result:\n")
	fmt.Printf("ID: %s\n", resp.Permission.Id)
	fmt.Printf("Resource: %s\n", resp.Permission.Resource)
	fmt.Printf("Action: %s\n", resp.Permission.Action)
	fmt.Printf("Description: %s\n", resp.Permission.Description)
}

func (c *RoleServiceClient) TestGetPermission() {
	fmt.Println("\n=== Test Get Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter permission ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.permissionClient.GetPermission(ctx, &proto_permission.GetPermissionRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetPermission: %v\n", err)
		return
	}

	fmt.Printf("Get Permission result:\n")
	fmt.Printf("ID: %s\n", resp.Permission.Id)
	fmt.Printf("Resource: %s\n", resp.Permission.Resource)
	fmt.Printf("Action: %s\n", resp.Permission.Action)
	fmt.Printf("Description: %s\n", resp.Permission.Description)
}

func (c *RoleServiceClient) TestListPermissions() {
	fmt.Println("\n=== Test List Permissions ===")

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.permissionClient.ListPermissions(ctx, &proto_permission.ListPermissionsRequest{})
	if err != nil {
		fmt.Printf("Error calling ListPermissions: %v\n", err)
		return
	}

	fmt.Printf("List Permissions result:\n")
	fmt.Printf("Total Permissions: %d\n", len(resp.Permissions))
	fmt.Printf("Permissions:\n")
	for i, permission := range resp.Permissions {
		fmt.Printf("  [%d] ID: %s, Resource: %s, Action: %s, Description: %s\n", i+1, permission.Id, permission.Resource, permission.Action, permission.Description)
	}
}

func (c *RoleServiceClient) TestUpdatePermission() {
	fmt.Println("\n=== Test Update Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter permission ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new resource: ")
	resource, _ := reader.ReadString('\n')
	resource = cleanInput(resource)

	fmt.Print("Enter new action: ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	fmt.Print("Enter new description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.permissionClient.UpdatePermission(ctx, &proto_permission.UpdatePermissionRequest{
		Id:          id,
		Resource:    resource,
		Action:      action,
		Description: description,
	})
	if err != nil {
		fmt.Printf("Error calling UpdatePermission: %v\n", err)
		return
	}

	fmt.Printf("Update Permission result:\n")
	fmt.Printf("ID: %s\n", resp.Permission.Id)
	fmt.Printf("Resource: %s\n", resp.Permission.Resource)
	fmt.Printf("Action: %s\n", resp.Permission.Action)
	fmt.Printf("Description: %s\n", resp.Permission.Description)
}

func (c *RoleServiceClient) TestDeletePermission() {
	fmt.Println("\n=== Test Delete Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter permission ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	_, err := c.permissionClient.DeletePermission(ctx, &proto_permission.DeletePermissionRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeletePermission: %v\n", err)
		return
	}

	fmt.Printf("Delete Permission result:\n")
	fmt.Printf("Permission deleted successfully\n")
}

// ================== Role Permission Service Tests ==================

func (c *RoleServiceClient) TestCreateRolePermission() {
	fmt.Println("\n=== Test Create Role Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role ID: ")
	roleId, _ := reader.ReadString('\n')
	roleId = cleanInput(roleId)

	fmt.Print("Enter permission ID: ")
	permissionId, _ := reader.ReadString('\n')
	permissionId = cleanInput(permissionId)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.rolePermissionClient.CreateRolePermission(ctx, &proto_role_permission.CreateRolePermissionRequest{
		RoleId:       roleId,
		PermissionId: permissionId,
	})
	if err != nil {
		fmt.Printf("Error calling CreateRolePermission: %v\n", err)
		return
	}

	fmt.Printf("Create Role Permission result:\n")
	fmt.Printf("Role ID: %s\n", resp.RolePermission.RoleId)
	fmt.Printf("Permission ID: %s\n", resp.RolePermission.PermissionId)
}

func (c *RoleServiceClient) TestListRolePermissions() {
	fmt.Println("\n=== Test List Role Permissions ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role ID (optional, press Enter to skip): ")
	roleId, _ := reader.ReadString('\n')
	roleId = cleanInput(roleId)

	fmt.Print("Enter permission ID (optional, press Enter to skip): ")
	permissionId, _ := reader.ReadString('\n')
	permissionId = cleanInput(permissionId)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	req := &proto_role_permission.ListRolePermissionsRequest{}
	if roleId != "" {
		req.Filter = &proto_role_permission.RolePermissionFilter{
			RoleId: roleId,
		}
	}
	if permissionId != "" && req.Filter != nil {
		req.Filter.PermissionId = permissionId
	} else if permissionId != "" {
		req.Filter = &proto_role_permission.RolePermissionFilter{
			PermissionId: permissionId,
		}
	}

	resp, err := c.rolePermissionClient.ListRolePermissions(ctx, req)
	if err != nil {
		fmt.Printf("Error calling ListRolePermissions: %v\n", err)
		return
	}

	fmt.Printf("List Role Permissions result:\n")
	fmt.Printf("Total Role Permissions: %d\n", len(resp.RolePermissions))
	for i, rp := range resp.RolePermissions {
		fmt.Printf("  [%d] Role ID: %s, Permission ID: %s\n", i+1, rp.RoleId, rp.PermissionId)
	}
}

func (c *RoleServiceClient) TestDeleteRolePermission() {
	fmt.Println("\n=== Test Delete Role Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter role ID: ")
	roleId, _ := reader.ReadString('\n')
	roleId = cleanInput(roleId)

	fmt.Print("Enter permission ID: ")
	permissionId, _ := reader.ReadString('\n')
	permissionId = cleanInput(permissionId)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	_, err := c.rolePermissionClient.DeleteRolePermission(ctx, &proto_role_permission.DeleteRolePermissionRequest{
		RoleId:       roleId,
		PermissionId: permissionId,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteRolePermission: %v\n", err)
		return
	}

	fmt.Printf("Delete Role Permission result:\n")
	fmt.Printf("Role Permission deleted successfully\n")
}

// ================== Resource Permission Service Tests ==================

func (c *RoleServiceClient) TestCreateResourcePermission() {
	fmt.Println("\n=== Test Create Resource Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter user ID: ")
	userId, _ := reader.ReadString('\n')
	userId = cleanInput(userId)

	fmt.Print("Enter resource type: ")
	resourceType, _ := reader.ReadString('\n')
	resourceType = cleanInput(resourceType)

	fmt.Print("Enter resource data json: ")
	resourceData, _ := reader.ReadString('\n')
	resourceData = cleanInput(resourceData)
	var resourceDataMap map[string]string
	err := json.Unmarshal([]byte(resourceData), &resourceDataMap)
	if err != nil {
		fmt.Printf("Error unmarshalling resource data: %v\n", err)
		return
	}

	fmt.Print("Enter action: ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.resourcePermissionClient.CreateResourcePermission(ctx, &proto_resource_permission.CreateResourcePermissionRequest{
		UserId:       userId,
		ResourceType: resourceType,
		ResourceData: resourceDataMap,
		Action:       action,
	})
	if err != nil {
		fmt.Printf("Error calling CreateResourcePermission: %v\n", err)
		return
	}

	fmt.Printf("Create Resource Permission result:\n")
	fmt.Printf("ID: %s\n", resp.ResourcePermission.Id)
	fmt.Printf("User ID: %s\n", resp.ResourcePermission.UserId)
	fmt.Printf("Resource Type: %s\n", resp.ResourcePermission.ResourceType)
	fmt.Printf("Resource Data: %s\n", resp.ResourcePermission.ResourceData)
	fmt.Printf("Action: %s\n", resp.ResourcePermission.Action)
}

func (c *RoleServiceClient) TestGetResourcePermission() {
	fmt.Println("\n=== Test Get Resource Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter resource permission ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.resourcePermissionClient.GetResourcePermission(ctx, &proto_resource_permission.GetResourcePermissionRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetResourcePermission: %v\n", err)
		return
	}

	fmt.Printf("Get Resource Permission result:\n")
	fmt.Printf("ID: %s\n", resp.ResourcePermission.Id)
	fmt.Printf("User ID: %s\n", resp.ResourcePermission.UserId)
	fmt.Printf("Resource Type: %s\n", resp.ResourcePermission.ResourceType)
	fmt.Printf("Resource Data: %s\n", resp.ResourcePermission.ResourceData)
	fmt.Printf("Action: %s\n", resp.ResourcePermission.Action)
}

func (c *RoleServiceClient) TestListResourcePermissions() {
	fmt.Println("\n=== Test List Resource Permissions ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter user ID (optional, press Enter to skip): ")
	userId, _ := reader.ReadString('\n')
	userId = cleanInput(userId)

	fmt.Print("Enter resource type (optional, press Enter to skip): ")
	resourceType, _ := reader.ReadString('\n')
	resourceType = cleanInput(resourceType)

	fmt.Print("Enter action (optional, press Enter to skip): ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	req := &proto_resource_permission.ListResourcePermissionsRequest{}
	if userId != "" || resourceType != "" || action != "" {
		req.Filter = &proto_resource_permission.ResourcePermissionFilter{}
		if userId != "" {
			req.Filter.UserId = userId
		}
		if resourceType != "" {
			req.Filter.ResourceType = resourceType
		}
		if action != "" {
			req.Filter.Action = action
		}
	}

	resp, err := c.resourcePermissionClient.ListResourcePermissions(ctx, req)
	if err != nil {
		fmt.Printf("Error calling ListResourcePermissions: %v\n", err)
		return
	}

	fmt.Printf("List Resource Permissions result:\n")
	fmt.Printf("Total Resource Permissions: %d\n", len(resp.ResourcePermissions))
	for i, rp := range resp.ResourcePermissions {
		fmt.Printf("  [%d] ID: %s, User ID: %s, Resource Type: %s, Resource Data: %s, Action: %s\n", i+1, rp.Id, rp.UserId, rp.ResourceType, rp.ResourceData, rp.Action)
	}
}

func (c *RoleServiceClient) TestUpdateResourcePermission() {
	fmt.Println("\n=== Test Update Resource Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter resource permission ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new user ID: ")
	userId, _ := reader.ReadString('\n')
	userId = cleanInput(userId)

	fmt.Print("Enter new resource type: ")
	resourceType, _ := reader.ReadString('\n')
	resourceType = cleanInput(resourceType)

	fmt.Print("Enter new resource data json: ")
	resourceData, _ := reader.ReadString('\n')
	resourceData = cleanInput(resourceData)
	var resourceDataMap map[string]string
	err := json.Unmarshal([]byte(resourceData), &resourceDataMap)
	if err != nil {
		fmt.Printf("Error unmarshalling resource data: %v\n", err)
		return
	}

	fmt.Print("Enter new action: ")
	action, _ := reader.ReadString('\n')
	action = cleanInput(action)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.resourcePermissionClient.UpdateResourcePermission(ctx, &proto_resource_permission.UpdateResourcePermissionRequest{
		Id:           id,
		UserId:       userId,
		ResourceType: resourceType,
		ResourceData: resourceDataMap,
		Action:       action,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateResourcePermission: %v\n", err)
		return
	}

	fmt.Printf("Update Resource Permission result:\n")
	fmt.Printf("ID: %s\n", resp.ResourcePermission.Id)
	fmt.Printf("User ID: %s\n", resp.ResourcePermission.UserId)
	fmt.Printf("Resource Type: %s\n", resp.ResourcePermission.ResourceType)
	fmt.Printf("Resource Data: %s\n", resp.ResourcePermission.ResourceData)
	fmt.Printf("Action: %s\n", resp.ResourcePermission.Action)
}

func (c *RoleServiceClient) TestDeleteResourcePermission() {
	fmt.Println("\n=== Test Delete Resource Permission ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter resource permission ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	_, err := c.resourcePermissionClient.DeleteResourcePermission(ctx, &proto_resource_permission.DeleteResourcePermissionRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteResourcePermission: %v\n", err)
		return
	}

	fmt.Printf("Delete Resource Permission result:\n")
	fmt.Printf("Resource Permission deleted successfully\n")
}

// ================== User Role Service Tests ==================

func (c *RoleServiceClient) TestCreateUserRole() {
	fmt.Println("\n=== Test Create User Role ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter user ID: ")
	userId, _ := reader.ReadString('\n')
	userId = cleanInput(userId)

	fmt.Print("Enter role ID: ")
	roleId, _ := reader.ReadString('\n')
	roleId = cleanInput(roleId)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	resp, err := c.userRoleClient.CreateUserRole(ctx, &proto_user_role.CreateUserRoleRequest{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		fmt.Printf("Error calling CreateUserRole: %v\n", err)
		return
	}

	fmt.Printf("Create User Role result:\n")
	fmt.Printf("User ID: %s\n", resp.Data.UserId)
	fmt.Printf("Role ID: %s\n", resp.Data.RoleId)
}

func (c *RoleServiceClient) TestListUserRoles() {
	fmt.Println("\n=== Test List User Roles ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter user ID (optional, press Enter to skip): ")
	userId, _ := reader.ReadString('\n')
	userId = cleanInput(userId)

	fmt.Print("Enter role ID (optional, press Enter to skip): ")
	roleId, _ := reader.ReadString('\n')
	roleId = cleanInput(roleId)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	req := &proto_user_role.ListUserRolesRequest{}
	if userId != "" || roleId != "" {
		req.Filter = &proto_user_role.UserRoleFilter{}
		if userId != "" {
			req.Filter.UserId = userId
		}
		if roleId != "" {
			req.Filter.RoleId = roleId
		}
	}

	resp, err := c.userRoleClient.ListUserRoles(ctx, req)
	if err != nil {
		fmt.Printf("Error calling ListUserRoles: %v\n", err)
		return
	}

	fmt.Printf("List User Roles result:\n")
	fmt.Printf("Total User Roles: %d\n", len(resp.UserRoles))
	for i, ur := range resp.UserRoles {
		fmt.Printf("  [%d] User ID: %s, Role ID: %s\n", i+1, ur.UserId, ur.RoleId)
	}
}

func (c *RoleServiceClient) TestDeleteUserRole() {
	fmt.Println("\n=== Test Delete User Role ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter user ID: ")
	userId, _ := reader.ReadString('\n')
	userId = cleanInput(userId)

	fmt.Print("Enter role ID: ")
	roleId, _ := reader.ReadString('\n')
	roleId = cleanInput(roleId)

	ctx, cancel := context.WithTimeout(c.ctx, 10*time.Second)
	defer cancel()

	_, err := c.userRoleClient.DeleteUserRole(ctx, &proto_user_role.DeleteUserRoleRequest{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteUserRole: %v\n", err)
		return
	}

	fmt.Printf("Delete User Role result:\n")
	fmt.Printf("User Role deleted successfully\n")
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== gRPC Role Service Test Client ===")
	fmt.Println("1. Role Service")
	fmt.Println("2. Permission Service")
	fmt.Println("3. Role Permission Service")
	fmt.Println("4. Resource Permission Service")
	fmt.Println("5. User Role Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printRoleMenu() {
	fmt.Println("\n=== Role Service ===")
	fmt.Println("1. Create Role")
	fmt.Println("2. Get Role By ID")
	fmt.Println("3. Get All Roles")
	fmt.Println("4. Update Role")
	fmt.Println("5. Delete Role")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printPermissionMenu() {
	fmt.Println("\n=== Permission Service ===")
	fmt.Println("1. Create Permission")
	fmt.Println("2. Get Permission")
	fmt.Println("3. List Permissions")
	fmt.Println("4. Update Permission")
	fmt.Println("5. Delete Permission")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printRolePermissionMenu() {
	fmt.Println("\n=== Role Permission Service ===")
	fmt.Println("1. Create Role Permission")
	fmt.Println("2. List Role Permissions")
	fmt.Println("3. Delete Role Permission")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printResourcePermissionMenu() {
	fmt.Println("\n=== Resource Permission Service ===")
	fmt.Println("1. Create Resource Permission")
	fmt.Println("2. Get Resource Permission")
	fmt.Println("3. List Resource Permissions")
	fmt.Println("4. Update Resource Permission")
	fmt.Println("5. Delete Resource Permission")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printUserRoleMenu() {
	fmt.Println("\n=== User Role Service ===")
	fmt.Println("1. Create User Role")
	fmt.Println("2. List User Roles")
	fmt.Println("3. Delete User Role")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	cacher := cache.NewCache(cache.ConfigCache{
		Addr:        viper.GetString("cache_addr"),
		DB:          viper.GetInt("cache_db"),
		Password:    viper.GetString("cache_password"),
		MaxIdle:     viper.GetInt("cache_max_idle"),
		MaxActive:   viper.GetInt("cache_max_active"),
		IdleTimeout: viper.GetInt("cache_idle_timeout"),
	})

	uCtx := user_context.NewUserContext()
	uCtx.UserID = "1"
	uCtx.Roles = []string{"admin", "user"}
	uCtx.Permissions = []user_context.Permission{
		{
			Resource: "admin",
			Action:   "read",
		},
	}
	userData, err := uCtx.ToBytes()
	if err != nil {
		log.Fatalf("Failed to convert user context to bytes: %v", err)
	}
	cacher.Set("1", userData, 0)

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewRoleServiceClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected successfully!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Role Service
			for {
				printRoleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateRole()
				case "2":
					client.TestGetRoleById()
				case "3":
					client.TestGetAllRoles()
				case "4":
					client.TestUpdateRole()
				case "5":
					client.TestDeleteRole()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "2":
			// Permission Service
			for {
				printPermissionMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreatePermission()
				case "2":
					client.TestGetPermission()
				case "3":
					client.TestListPermissions()
				case "4":
					client.TestUpdatePermission()
				case "5":
					client.TestDeletePermission()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "3":
			// Role Permission Service
			for {
				printRolePermissionMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateRolePermission()
				case "2":
					client.TestListRolePermissions()
				case "3":
					client.TestDeleteRolePermission()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "4":
			// Resource Permission Service
			for {
				printResourcePermissionMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateResourcePermission()
				case "2":
					client.TestGetResourcePermission()
				case "3":
					client.TestListResourcePermissions()
				case "4":
					client.TestUpdateResourcePermission()
				case "5":
					client.TestDeleteResourcePermission()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "5":
			// User Role Service
			for {
				printUserRoleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateUserRole()
				case "2":
					client.TestListUserRoles()
				case "3":
					client.TestDeleteUserRole()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

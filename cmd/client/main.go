package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	proto_module "github.com/anhvanhoa/sf-proto/gen/module/v1"
	proto_module_child "github.com/anhvanhoa/sf-proto/gen/module_child/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddress = "localhost:50052"
)

type GRPCClient struct {
	moduleClient      proto_module.ModuleServiceClient
	moduleChildClient proto_module_child.ModuleChildServiceClient
	conn              *grpc.ClientConn
}

func NewGRPCClient(address string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &GRPCClient{
		moduleClient:      proto_module.NewModuleServiceClient(conn),
		moduleChildClient: proto_module_child.NewModuleChildServiceClient(conn),
		conn:              conn,
	}, nil
}

func (c *GRPCClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// Module Service Tests
func (c *GRPCClient) TestCreateModule() {
	fmt.Println("\n=== Test CreateModule ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter module description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Enter module status (active/inactive): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.CreateModule(ctx, &proto_module.CreateModuleRequest{
		Name:        name,
		Description: description,
		Status:      status,
	})
	if err != nil {
		fmt.Printf("Error calling CreateModule: %v\n", err)
		return
	}

	fmt.Printf("Module created successfully:\n")
	fmt.Printf("ID: %s\n", resp.Module.Id)
	fmt.Printf("Name: %s\n", resp.Module.Name)
	fmt.Printf("Description: %s\n", resp.Module.Description)
	fmt.Printf("Status: %s\n", resp.Module.Status)
}

func (c *GRPCClient) TestGetModule() {
	fmt.Println("\n=== Test GetModule ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter module ID: ")
	moduleID, _ := reader.ReadString('\n')
	moduleID = strings.TrimSpace(moduleID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.GetModule(ctx, &proto_module.GetModuleRequest{
		Id: moduleID,
	})
	if err != nil {
		fmt.Printf("Error calling GetModule: %v\n", err)
		return
	}

	fmt.Printf("Module found:\n")
	fmt.Printf("ID: %s\n", resp.Module.Id)
	fmt.Printf("Name: %s\n", resp.Module.Name)
	fmt.Printf("Description: %s\n", resp.Module.Description)
	fmt.Printf("Status: %s\n", resp.Module.Status)
	fmt.Printf("Created At: %s\n", resp.Module.CreatedAt)
	if resp.Module.UpdatedAt != "" {
		fmt.Printf("Updated At: %s\n", resp.Module.UpdatedAt)
	}
}

func (c *GRPCClient) TestListModules() {
	fmt.Println("\n=== Test ListModules ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter page number (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = strings.TrimSpace(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.ParseInt(pageStr, 10, 32); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter page size (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = strings.TrimSpace(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.ParseInt(limitStr, 10, 32); err == nil {
			limit = int32(l)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.ListModules(ctx, &proto_module.ListModulesRequest{
		Pagination: &proto_module.PaginationRequest{
			Page:  page,
			Limit: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListModules: %v\n", err)
		return
	}

	fmt.Printf("Found %d modules (Page %d of %d, Total: %d):\n",
		len(resp.Modules), resp.Pagination.Page, resp.Pagination.TotalPages, resp.Pagination.Total)

	for i, module := range resp.Modules {
		fmt.Printf("Module %d:\n", i+1)
		fmt.Printf("  ID: %s\n", module.Id)
		fmt.Printf("  Name: %s\n", module.Name)
		fmt.Printf("  Description: %s\n", module.Description)
		fmt.Printf("  Status: %s\n", module.Status)
		fmt.Printf("  Created At: %s\n", module.CreatedAt)
		if module.UpdatedAt != "" {
			fmt.Printf("  Updated At: %s\n", module.UpdatedAt)
		}
		fmt.Println()
	}
}

func (c *GRPCClient) TestUpdateModule() {
	fmt.Println("\n=== Test UpdateModule ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID to update: ")
	moduleID, _ := reader.ReadString('\n')
	moduleID = strings.TrimSpace(moduleID)

	fmt.Print("Enter new module name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter new module description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Enter new module status (active/inactive): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.UpdateModule(ctx, &proto_module.UpdateModuleRequest{
		Id:          moduleID,
		Name:        name,
		Description: description,
		Status:      status,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateModule: %v\n", err)
		return
	}

	fmt.Printf("Module updated successfully:\n")
	fmt.Printf("ID: %s\n", resp.Module.Id)
	fmt.Printf("Name: %s\n", resp.Module.Name)
	fmt.Printf("Description: %s\n", resp.Module.Description)
	fmt.Printf("Status: %s\n", resp.Module.Status)
	fmt.Printf("Created At: %s\n", resp.Module.CreatedAt)
	if resp.Module.UpdatedAt != "" {
		fmt.Printf("Updated At: %s\n", resp.Module.UpdatedAt)
	}
}

func (c *GRPCClient) TestDeleteModule() {
	fmt.Println("\n=== Test DeleteModule ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter module ID to delete: ")
	moduleID, _ := reader.ReadString('\n')
	moduleID = strings.TrimSpace(moduleID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.DeleteModule(ctx, &proto_module.DeleteModuleRequest{
		Id: moduleID,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteModule: %v\n", err)
		return
	}

	fmt.Printf("Delete result: Success = %t\n", resp.Success)
}

// Module Child Service Tests
func (c *GRPCClient) TestCreateModuleChild() {
	fmt.Println("\n=== Test CreateModuleChild ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID: ")
	moduleID, _ := reader.ReadString('\n')
	moduleID = strings.TrimSpace(moduleID)

	fmt.Print("Enter module child name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	fmt.Print("Enter method (GET/POST/PUT/DELETE): ")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(method)

	fmt.Print("Is private? (true/false): ")
	isPrivateStr, _ := reader.ReadString('\n')
	isPrivateStr = strings.TrimSpace(isPrivateStr)
	isPrivate := false
	if isPrivateStr == "true" {
		isPrivate = true
	}

	fmt.Print("Enter status (active/inactive): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.CreateModuleChild(ctx, &proto_module_child.CreateModuleChildRequest{
		ModuleId:  moduleID,
		Name:      name,
		Path:      path,
		Method:    method,
		IsPrivate: isPrivate,
		Status:    status,
	})
	if err != nil {
		fmt.Printf("Error calling CreateModuleChild: %v\n", err)
		return
	}

	fmt.Printf("Module child created successfully:\n")
	fmt.Printf("ID: %s\n", resp.ModuleChild.Id)
	fmt.Printf("Module ID: %s\n", resp.ModuleChild.ModuleId)
	fmt.Printf("Name: %s\n", resp.ModuleChild.Name)
	fmt.Printf("Path: %s\n", resp.ModuleChild.Path)
	fmt.Printf("Method: %s\n", resp.ModuleChild.Method)
	fmt.Printf("Is Private: %t\n", resp.ModuleChild.IsPrivate)
	fmt.Printf("Status: %s\n", resp.ModuleChild.Status)
	fmt.Printf("Created At: %s\n", resp.ModuleChild.CreatedAt)
}

func (c *GRPCClient) TestGetModuleChild() {
	fmt.Println("\n=== Test GetModuleChild ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter module child ID: ")
	moduleChildID, _ := reader.ReadString('\n')
	moduleChildID = strings.TrimSpace(moduleChildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.GetModuleChild(ctx, &proto_module_child.GetModuleChildRequest{
		Id: moduleChildID,
	})
	if err != nil {
		fmt.Printf("Error calling GetModuleChild: %v\n", err)
		return
	}

	fmt.Printf("Module child found:\n")
	fmt.Printf("ID: %s\n", resp.ModuleChild.Id)
	fmt.Printf("Module ID: %s\n", resp.ModuleChild.ModuleId)
	fmt.Printf("Name: %s\n", resp.ModuleChild.Name)
	fmt.Printf("Path: %s\n", resp.ModuleChild.Path)
	fmt.Printf("Method: %s\n", resp.ModuleChild.Method)
	fmt.Printf("Is Private: %t\n", resp.ModuleChild.IsPrivate)
	fmt.Printf("Status: %s\n", resp.ModuleChild.Status)
	fmt.Printf("Created At: %s\n", resp.ModuleChild.CreatedAt)
	if resp.ModuleChild.UpdatedAt != "" {
		fmt.Printf("Updated At: %s\n", resp.ModuleChild.UpdatedAt)
	}
}

func (c *GRPCClient) TestListModuleChildren() {
	fmt.Println("\n=== Test ListModuleChildren ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID (optional, leave empty for all): ")
	moduleID, _ := reader.ReadString('\n')
	moduleID = strings.TrimSpace(moduleID)

	fmt.Print("Enter page number (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = strings.TrimSpace(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.ParseInt(pageStr, 10, 32); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter page size (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = strings.TrimSpace(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.ParseInt(limitStr, 10, 32); err == nil {
			limit = int32(l)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.ListModuleChildren(ctx, &proto_module_child.ListModuleChildrenRequest{
		ModuleId: moduleID,
		Pagination: &proto_module_child.PaginationRequest{
			Page:  page,
			Limit: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListModuleChildren: %v\n", err)
		return
	}

	fmt.Printf("Found %d module children (Page %d of %d, Total: %d):\n",
		len(resp.ModuleChildren), resp.Pagination.Page, resp.Pagination.TotalPages, resp.Pagination.Total)

	for i, moduleChild := range resp.ModuleChildren {
		fmt.Printf("Module Child %d:\n", i+1)
		fmt.Printf("  ID: %s\n", moduleChild.Id)
		fmt.Printf("  Module ID: %s\n", moduleChild.ModuleId)
		fmt.Printf("  Name: %s\n", moduleChild.Name)
		fmt.Printf("  Path: %s\n", moduleChild.Path)
		fmt.Printf("  Method: %s\n", moduleChild.Method)
		fmt.Printf("  Is Private: %t\n", moduleChild.IsPrivate)
		fmt.Printf("  Status: %s\n", moduleChild.Status)
		fmt.Printf("  Created At: %s\n", moduleChild.CreatedAt)
		if moduleChild.UpdatedAt != "" {
			fmt.Printf("  Updated At: %s\n", moduleChild.UpdatedAt)
		}
		fmt.Println()
	}
}

func (c *GRPCClient) TestUpdateModuleChild() {
	fmt.Println("\n=== Test UpdateModuleChild ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module child ID to update: ")
	moduleChildID, _ := reader.ReadString('\n')
	moduleChildID = strings.TrimSpace(moduleChildID)

	fmt.Print("Enter module ID: ")
	moduleID, _ := reader.ReadString('\n')
	moduleID = strings.TrimSpace(moduleID)

	fmt.Print("Enter new module child name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter new path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	fmt.Print("Enter new method (GET/POST/PUT/DELETE): ")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(method)

	fmt.Print("Is private? (true/false): ")
	isPrivateStr, _ := reader.ReadString('\n')
	isPrivateStr = strings.TrimSpace(isPrivateStr)
	isPrivate := false
	if isPrivateStr == "true" {
		isPrivate = true
	}

	fmt.Print("Enter new status (active/inactive): ")
	status, _ := reader.ReadString('\n')
	status = strings.TrimSpace(status)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.UpdateModuleChild(ctx, &proto_module_child.UpdateModuleChildRequest{
		Id:        moduleChildID,
		ModuleId:  moduleID,
		Name:      name,
		Path:      path,
		Method:    method,
		IsPrivate: isPrivate,
		Status:    status,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateModuleChild: %v\n", err)
		return
	}

	fmt.Printf("Module child updated successfully:\n")
	fmt.Printf("ID: %s\n", resp.ModuleChild.Id)
	fmt.Printf("Module ID: %s\n", resp.ModuleChild.ModuleId)
	fmt.Printf("Name: %s\n", resp.ModuleChild.Name)
	fmt.Printf("Path: %s\n", resp.ModuleChild.Path)
	fmt.Printf("Method: %s\n", resp.ModuleChild.Method)
	fmt.Printf("Is Private: %t\n", resp.ModuleChild.IsPrivate)
	fmt.Printf("Status: %s\n", resp.ModuleChild.Status)
	fmt.Printf("Created At: %s\n", resp.ModuleChild.CreatedAt)
	if resp.ModuleChild.UpdatedAt != "" {
		fmt.Printf("Updated At: %s\n", resp.ModuleChild.UpdatedAt)
	}
}

func (c *GRPCClient) TestDeleteModuleChild() {
	fmt.Println("\n=== Test DeleteModuleChild ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter module child ID to delete: ")
	moduleChildID, _ := reader.ReadString('\n')
	moduleChildID = strings.TrimSpace(moduleChildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.DeleteModuleChild(ctx, &proto_module_child.DeleteModuleChildRequest{
		Id: moduleChildID,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteModuleChild: %v\n", err)
		return
	}

	fmt.Printf("Delete result: Success = %t\n", resp.Success)
}

func printMenu() {
	fmt.Println("\n=== gRPC Module Service Test Client ===")
	fmt.Println("1. Module Service Tests")
	fmt.Println("  1.1 Create Module")
	fmt.Println("  1.2 Get Module By ID")
	fmt.Println("  1.3 List Modules")
	fmt.Println("  1.4 Update Module")
	fmt.Println("  1.5 Delete Module")
	fmt.Println("2. Module Child Service Tests")
	fmt.Println("  2.1 Create Module Child")
	fmt.Println("  2.2 Get Module Child By ID")
	fmt.Println("  2.3 List Module Children")
	fmt.Println("  2.4 Update Module Child")
	fmt.Println("  2.5 Delete Module Child")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func main() {
	// Get server address from command line or use default
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewGRPCClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected successfully!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1.1":
			client.TestCreateModule()
		case "1.2":
			client.TestGetModule()
		case "1.3":
			client.TestListModules()
		case "1.4":
			client.TestUpdateModule()
		case "1.5":
			client.TestDeleteModule()
		case "2.1":
			client.TestCreateModuleChild()
		case "2.2":
			client.TestGetModuleChild()
		case "2.3":
			client.TestListModuleChildren()
		case "2.4":
			client.TestUpdateModuleChild()
		case "2.5":
			client.TestDeleteModuleChild()
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

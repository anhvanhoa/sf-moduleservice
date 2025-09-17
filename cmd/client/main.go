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
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yaml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

type ModuleServiceClient struct {
	moduleClient      proto_module.ModuleServiceClient
	moduleChildClient proto_module_child.ModuleChildServiceClient
	conn              *grpc.ClientConn
}

func NewModuleServiceClient(address string) (*ModuleServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &ModuleServiceClient{
		moduleClient:      proto_module.NewModuleServiceClient(conn),
		moduleChildClient: proto_module_child.NewModuleChildServiceClient(conn),
		conn:              conn,
	}, nil
}

func (c *ModuleServiceClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

// ================== Module Service Tests ==================

func (c *ModuleServiceClient) TestCreateModule() {
	fmt.Println("\n=== Test Create Module ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Enter status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

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

	fmt.Printf("Create Module result:\n")
	fmt.Printf("ID: %s\n", resp.Module.Id)
	fmt.Printf("Name: %s\n", resp.Module.Name)
	fmt.Printf("Description: %s\n", resp.Module.Description)
	fmt.Printf("Status: %s\n", resp.Module.Status)
}

func (c *ModuleServiceClient) TestGetModule() {
	fmt.Println("\n=== Test Get Module ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.GetModule(ctx, &proto_module.GetModuleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetModule: %v\n", err)
		return
	}

	fmt.Printf("Get Module result:\n")
	fmt.Printf("ID: %s\n", resp.Module.Id)
	fmt.Printf("Name: %s\n", resp.Module.Name)
	fmt.Printf("Description: %s\n", resp.Module.Description)
	fmt.Printf("Status: %s\n", resp.Module.Status)
	fmt.Printf("Created At: %s\n", resp.Module.CreatedAt)
	if resp.Module.UpdatedAt != "" {
		fmt.Printf("Updated At: %s\n", resp.Module.UpdatedAt)
	}
}

func (c *ModuleServiceClient) TestListModules() {
	fmt.Println("\n=== Test List Modules ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter page (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = cleanInput(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
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

	fmt.Printf("List Modules result:\n")
	fmt.Printf("Total: %d\n", resp.Pagination.Total)
	fmt.Printf("Page: %d\n", resp.Pagination.Page)
	fmt.Printf("Limit: %d\n", resp.Pagination.Limit)
	fmt.Printf("Total Pages: %d\n", resp.Pagination.TotalPages)
	fmt.Printf("Modules:\n")
	for i, module := range resp.Modules {
		fmt.Printf("  [%d] ID: %s, Name: %s, Description: %s, Status: %s\n", i+1, module.Id, module.Name, module.Description, module.Status)
	}
}

func (c *ModuleServiceClient) TestUpdateModule() {
	fmt.Println("\n=== Test Update Module ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter new description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Enter new status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.UpdateModule(ctx, &proto_module.UpdateModuleRequest{
		Id:          id,
		Name:        name,
		Description: description,
		Status:      status,
	})
	if err != nil {
		fmt.Printf("Error calling UpdateModule: %v\n", err)
		return
	}

	fmt.Printf("Update Module result:\n")
	fmt.Printf("ID: %s\n", resp.Module.Id)
	fmt.Printf("Name: %s\n", resp.Module.Name)
	fmt.Printf("Description: %s\n", resp.Module.Description)
	fmt.Printf("Status: %s\n", resp.Module.Status)
	fmt.Printf("Updated At: %s\n", resp.Module.UpdatedAt)
}

func (c *ModuleServiceClient) TestDeleteModule() {
	fmt.Println("\n=== Test Delete Module ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleClient.DeleteModule(ctx, &proto_module.DeleteModuleRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteModule: %v\n", err)
		return
	}

	fmt.Printf("Delete Module result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
}

// ================== Module Child Service Tests ==================

func (c *ModuleServiceClient) TestCreateModuleChild() {
	fmt.Println("\n=== Test Create Module Child ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID: ")
	moduleId, _ := reader.ReadString('\n')
	moduleId = cleanInput(moduleId)

	fmt.Print("Enter child name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter path: ")
	path, _ := reader.ReadString('\n')
	path = cleanInput(path)

	fmt.Print("Enter method (GET, POST, PUT, DELETE): ")
	method, _ := reader.ReadString('\n')
	method = cleanInput(method)

	fmt.Print("Is private? (true/false, default false): ")
	isPrivateStr, _ := reader.ReadString('\n')
	isPrivateStr = cleanInput(isPrivateStr)
	isPrivate := false
	if isPrivateStr == "true" {
		isPrivate = true
	}

	fmt.Print("Enter status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.CreateModuleChild(ctx, &proto_module_child.CreateModuleChildRequest{
		ModuleId:  moduleId,
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

	fmt.Printf("Create Module Child result:\n")
	fmt.Printf("ID: %s\n", resp.ModuleChild.Id)
	fmt.Printf("Module ID: %s\n", resp.ModuleChild.ModuleId)
	fmt.Printf("Name: %s\n", resp.ModuleChild.Name)
	fmt.Printf("Path: %s\n", resp.ModuleChild.Path)
	fmt.Printf("Method: %s\n", resp.ModuleChild.Method)
	fmt.Printf("Is Private: %t\n", resp.ModuleChild.IsPrivate)
	fmt.Printf("Status: %s\n", resp.ModuleChild.Status)
	fmt.Printf("Created At: %s\n", resp.ModuleChild.CreatedAt)
}

func (c *ModuleServiceClient) TestGetModuleChild() {
	fmt.Println("\n=== Test Get Module Child ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module child ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.GetModuleChild(ctx, &proto_module_child.GetModuleChildRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling GetModuleChild: %v\n", err)
		return
	}

	fmt.Printf("Get Module Child result:\n")
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

func (c *ModuleServiceClient) TestListModuleChildren() {
	fmt.Println("\n=== Test List Module Children ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module ID: ")
	moduleId, _ := reader.ReadString('\n')
	moduleId = cleanInput(moduleId)

	fmt.Print("Enter page (default 1): ")
	pageStr, _ := reader.ReadString('\n')
	pageStr = cleanInput(pageStr)
	page := int32(1)
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = int32(p)
		}
	}

	fmt.Print("Enter limit (default 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.ListModuleChildren(ctx, &proto_module_child.ListModuleChildrenRequest{
		ModuleId: moduleId,
		Pagination: &proto_module_child.PaginationRequest{
			Page:  page,
			Limit: limit,
		},
	})
	if err != nil {
		fmt.Printf("Error calling ListModuleChildren: %v\n", err)
		return
	}

	fmt.Printf("List Module Children result:\n")
	fmt.Printf("Total: %d\n", resp.Pagination.Total)
	fmt.Printf("Page: %d\n", resp.Pagination.Page)
	fmt.Printf("Limit: %d\n", resp.Pagination.Limit)
	fmt.Printf("Total Pages: %d\n", resp.Pagination.TotalPages)
	fmt.Printf("Module Children:\n")
	for i, child := range resp.ModuleChildren {
		fmt.Printf("  [%d] ID: %s, Name: %s, Path: %s, Method: %s, Status: %s\n", i+1, child.Id, child.Name, child.Path, child.Method, child.Status)
	}
}

func (c *ModuleServiceClient) TestUpdateModuleChild() {
	fmt.Println("\n=== Test Update Module Child ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module child ID: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Enter new name: ")
	name, _ := reader.ReadString('\n')
	name = cleanInput(name)

	fmt.Print("Enter new path: ")
	path, _ := reader.ReadString('\n')
	path = cleanInput(path)

	fmt.Print("Enter new method (GET, POST, PUT, DELETE): ")
	method, _ := reader.ReadString('\n')
	method = cleanInput(method)

	fmt.Print("Is private? (true/false): ")
	isPrivateStr, _ := reader.ReadString('\n')
	isPrivateStr = cleanInput(isPrivateStr)
	isPrivate := false
	if isPrivateStr == "true" {
		isPrivate = true
	}

	fmt.Print("Enter new status (default active): ")
	status, _ := reader.ReadString('\n')
	status = cleanInput(status)
	if status == "" {
		status = "active"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.UpdateModuleChild(ctx, &proto_module_child.UpdateModuleChildRequest{
		Id:        id,
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

	fmt.Printf("Update Module Child result:\n")
	fmt.Printf("ID: %s\n", resp.ModuleChild.Id)
	fmt.Printf("Module ID: %s\n", resp.ModuleChild.ModuleId)
	fmt.Printf("Name: %s\n", resp.ModuleChild.Name)
	fmt.Printf("Path: %s\n", resp.ModuleChild.Path)
	fmt.Printf("Method: %s\n", resp.ModuleChild.Method)
	fmt.Printf("Is Private: %t\n", resp.ModuleChild.IsPrivate)
	fmt.Printf("Status: %s\n", resp.ModuleChild.Status)
	fmt.Printf("Updated At: %s\n", resp.ModuleChild.UpdatedAt)
}

func (c *ModuleServiceClient) TestDeleteModuleChild() {
	fmt.Println("\n=== Test Delete Module Child ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter module child ID to delete: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.moduleChildClient.DeleteModuleChild(ctx, &proto_module_child.DeleteModuleChildRequest{
		Id: id,
	})
	if err != nil {
		fmt.Printf("Error calling DeleteModuleChild: %v\n", err)
		return
	}

	fmt.Printf("Delete Module Child result:\n")
	fmt.Printf("Success: %t\n", resp.Success)
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== gRPC Module Service Test Client ===")
	fmt.Println("1. Module Service")
	fmt.Println("2. Module Child Service")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
}

func printModuleMenu() {
	fmt.Println("\n=== Module Service ===")
	fmt.Println("1. Create Module")
	fmt.Println("2. Get Module")
	fmt.Println("3. List Modules")
	fmt.Println("4. Update Module")
	fmt.Println("5. Delete Module")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func printModuleChildMenu() {
	fmt.Println("\n=== Module Child Service ===")
	fmt.Println("1. Create Module Child")
	fmt.Println("2. Get Module Child")
	fmt.Println("3. List Module Children")
	fmt.Println("4. Update Module Child")
	fmt.Println("5. Delete Module Child")
	fmt.Println("0. Back to Main Menu")
	fmt.Print("Enter your choice: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Connecting to gRPC server at %s...\n", address)
	client, err := NewModuleServiceClient(address)
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
			// Module Service
			for {
				printModuleMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateModule()
				case "2":
					client.TestGetModule()
				case "3":
					client.TestListModules()
				case "4":
					client.TestUpdateModule()
				case "5":
					client.TestDeleteModule()
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
			// Module Child Service
			for {
				printModuleChildMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateModuleChild()
				case "2":
					client.TestGetModuleChild()
				case "3":
					client.TestListModuleChildren()
				case "4":
					client.TestUpdateModuleChild()
				case "5":
					client.TestDeleteModuleChild()
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

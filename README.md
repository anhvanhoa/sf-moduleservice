# Role Service

[![Go Version](https://img.shields.io/badge/Go-1.23.0-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/role-service)](https://goreportcard.com/report/github.com/yourusername/role-service)

Một thư viện Go cung cấp các chức năng xác thực và quản lý vai trò người dùng. Dự án này bao gồm cả một microservice hoàn chỉnh và một thư viện có thể sử dụng lại trong các dự án khác.

## 🚀 Tính năng

### Thư viện (pkg/roleservice)
- **Xác thực người dùng**: Đăng ký, đăng nhập, đăng xuất
- **Quản lý token**: JWT access token và refresh token
- **Xác thực tài khoản**: Gửi và xác thực mã OTP
- **Đặt lại mật khẩu**: Quên mật khẩu và đặt lại bằng mã hoặc token
- **Quản lý người dùng**: CRUD operations cho người dùng
- **Bảo mật**: Mã hóa mật khẩu với Argon2id
- **Logging**: Structured logging với Zap
- **Validation**: Input validation với struct tags

### Microservice
- **API gRPC**: Giao tiếp hiệu suất cao với Protocol Buffers
- **Cơ sở dữ liệu**: PostgreSQL với migration
- **Validation**: Input validation với protobuf-validate

## 🛠️ Công nghệ sử dụng

- **Backend**: Go 1.23.0
- **Framework**: gRPC với Protocol Buffers
- **Database**: PostgreSQL với go-pg ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: Argon2id
- **Configuration**: Viper
- **Logging**: Zap
- **Validation**: protobuf-validate
- **ID Generation**: NanoID

## 📋 Yêu cầu hệ thống

- Go 1.23.0 hoặc cao hơn
- PostgreSQL 12.0 hoặc cao hơn
- Redis (tùy chọn, cho caching)
- Docker (tùy chọn)

## 🚀 Cài đặt

### Sử dụng như thư viện

```bash
go get github.com/yourusername/role-service/pkg/roleservice
```

### Chạy microservice

#### 1. Clone repository

```bash
git clone https://github.com/yourusername/role-service.git
cd role-service
```

#### 2. Cài đặt dependencies

```bash
go mod download
```

### 3. Cài đặt công cụ cần thiết

```bash
# Cài đặt migrate CLI
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Cài đặt buf CLI
go install github.com/bufbuild/buf/cmd/buf@latest

# Cài đặt protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 4. Cấu hình môi trường

Tạo file `dev.config.yaml` dựa trên template:

```yaml
# Database Configuration
database:
  host: localhost
  port: 5432
  user: postgres
  password: 123456
  name: userservice_db
  sslmode: disable

# Redis Configuration (optional)
redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

# JWT Configuration
jwt:
  secret_key: "your-secret-key-here"
  access_token_expiry: 15m
  refresh_token_expiry: 7d

# Server Configuration
server:
  grpc_port: 40051
  http_port: 8080

# Logging
logging:
  level: debug
  format: json
```

### 5. Khởi tạo cơ sở dữ liệu

```bash
# Tạo database
make dev-create-db

# Chạy migrations
make migrate-dev-up
```

### 6. Chạy ứng dụng

```bash
# Build ứng dụng
make build

# Chạy ứng dụng
make run
```

Hoặc chạy trực tiếp:

```bash
go run cmd/main.go
```

## 📖 Sử dụng

### Sử dụng thư viện

Xem [pkg/roleservice/README.md](pkg/roleservice/README.md) để biết chi tiết về cách sử dụng thư viện.

### API Endpoints (Microservice)

Dự án sử dụng gRPC với các service sau:

#### AuthService

- `CheckToken` - Kiểm tra tính hợp lệ của token
- `Login` - Đăng nhập người dùng
- `Register` - Đăng ký tài khoản mới
- `RefreshToken` - Làm mới access token
- `Logout` - Đăng xuất
- `VerifyAccount` - Xác thực tài khoản
- `ForgotPassword` - Yêu cầu đặt lại mật khẩu
- `ResetPasswordByCode` - Đặt lại mật khẩu bằng mã
- `ResetPasswordByToken` - Đặt lại mật khẩu bằng token
- `CheckCode` - Kiểm tra mã xác thực

### Ví dụ sử dụng với gRPC client

```go
package main

import (
    "context"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "your-project/proto/exam/v1"
)

func main() {
    conn, err := grpc.Dial("localhost:40051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewAuthServiceClient(conn)

    // Đăng nhập
    loginResp, err := client.Login(context.Background(), &pb.LoginRequest{
        EmailOrPhone: "user@example.com",
        Password:     "password123",
        Os:           "web",
    })
    if err != nil {
        log.Fatalf("Login failed: %v", err)
    }

    log.Printf("Login successful: %s", loginResp.Message)
}
```

## 🏗️ Cấu trúc dự án

```
role-service/
├── pkg/               # Thư viện có thể sử dụng lại
│   └── roleservice/   # Role service library
│       ├── roleservice.go  # Main library interface
│       ├── types.go        # Request/Response types
│       ├── example/        # Usage examples
│       └── README.md       # Library documentation
├── bootstrap/          # Khởi tạo ứng dụng
│   ├── app.go         # Cấu hình ứng dụng chính
│   ├── database.go    # Kết nối database
│   ├── env.go         # Cấu hình môi trường
│   ├── redis.go       # Kết nối Redis
│   └── validator.go   # Cấu hình validation
├── cmd/               # Entry point
│   └── main.go        # Main function
├── domain/            # Domain layer
│   ├── common/        # Common utilities
│   ├── entity/        # Domain entities
│   ├── repository/    # Repository interfaces
│   ├── service/       # Domain services
│   └── usecase/       # Use cases
├── infrastructure/    # Infrastructure layer
│   ├── grpc_service/  # gRPC server implementation
│   ├── repo/          # Repository implementations
│   └── service/       # External service implementations
├── migrations/        # Database migrations
├── proto/             # Protocol Buffers definitions
│   └── exam/v1/       # API definitions
├── logs/              # Log files
├── go.mod             # Go modules
├── go.sum             # Go modules checksum
├── MakeFile           # Build scripts
├── dev.config.yaml    # Development configuration
└── README.md          # This file
```

## 🔧 Development

### Tạo migration mới

```bash
make migrate-dev-create name=migration_name
```

### Chạy migrations

```bash
# Chạy tất cả migrations
make migrate-dev-up

# Rollback migration cuối
make migrate-dev-down

# Reset database
make migrate-dev-reset

# Drop database
make migrate-dev-drop
```

### Generate Protocol Buffers

```bash
# Generate Go code từ proto files
buf generate
```

### Chạy tests

```bash
go test ./...
```

### Linting và formatting

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run
```

## 🐳 Docker

### Build Docker image

```bash
docker build -t role-service .
```

### Chạy với Docker Compose

Tạo file `docker-compose.yml`:

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "40051:40051"
    environment:
      - CONFIG_FILE=dev.config.yaml
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: userservice_db
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 123456
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

volumes:
  postgres_data:
```

Chạy:

```bash
docker-compose up -d
```

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `CONFIG_FILE` | Path to configuration file | `dev.config.yaml` |
| `GRPC_PORT` | gRPC server port | `40051` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | `123456` |
| `DB_NAME` | Database name | `userservice_db` |
| `JWT_SECRET` | JWT secret key | `your-secret-key` |

## 🤝 Đóng góp

Chúng tôi rất hoan nghênh mọi đóng góp! Vui lòng đọc [CONTRIBUTING.md](CONTRIBUTING.md) để biết thêm chi tiết.

### Quy trình đóng góp

1. Fork dự án
2. Tạo feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit thay đổi (`git commit -m 'Add some AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Mở Pull Request

## 📄 License

Dự án này được phân phối dưới giấy phép MIT. Xem file [LICENSE](LICENSE) để biết thêm chi tiết.

## 🆘 Hỗ trợ

Nếu bạn gặp vấn đề hoặc có câu hỏi:

- Tạo [Issue](https://github.com/yourusername/role-service/issues) trên GitHub
- Liên hệ qua email: your-email@example.com
- Tham gia [Discussions](https://github.com/yourusername/role-service/discussions)

## 🙏 Acknowledgments

- [gRPC](https://grpc.io/) - High-performance RPC framework
- [go-pg](https://github.com/go-pg/pg) - PostgreSQL ORM for Go
- [Zap](https://github.com/uber-go/zap) - Fast, structured logging
- [Viper](https://github.com/spf13/viper) - Configuration solution for Go applications

---

⭐ Nếu dự án này hữu ích, hãy cho chúng tôi một star trên GitHub!

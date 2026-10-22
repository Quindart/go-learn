# Bài tập ôn tập gRPC: Todo Service

## Mục tiêu

Xây dựng một ứng dụng quản lý công việc đơn giản bằng Go, gồm:

- Một gRPC server lưu dữ liệu trong bộ nhớ.
- Một CLI client dùng để gọi các RPC.
- Một file `.proto` mô tả service và message.

Sau bài tập này, bạn cần hiểu được quy trình:

```text
Định nghĩa .proto -> sinh Go code -> cài đặt server -> gọi từ client -> xử lý lỗi
```

## Làm lần lượt từ đầu đến khi hoàn thành

Phần này là lộ trình thao tác cụ thể. Chạy từng lệnh theo đúng thứ tự và chỉ chuyển sang mục tiếp theo khi checkpoint của mục hiện tại đã đạt.

### Thao tác 1: Mở terminal tại repository

```bash
cd /Users/quindart/go-learn
pwd
```

Kết quả của `pwd` phải là:

```text
/Users/quindart/go-learn
```

### Thao tác 2: Đi vào thư mục bài tập

```bash
cd grpc-review
pwd
```

Kết quả phải kết thúc bằng:

```text
/go-learn/grpc-review
```

Từ thao tác này trở đi, tất cả lệnh đều được chạy trong thư mục `grpc-review`.

### Thao tác 3: Khởi tạo Go module

```bash
go mod init example.com/grpc-review
```

Kết quả mong đợi:

```text
go: creating new go.mod: module example.com/grpc-review
```

Nếu nhận lỗi `go.mod already exists`, không cần chạy lại. Kiểm tra module path:

```bash
go list -m
```

Kết quả phải là `example.com/grpc-review`. Nếu khác, các import và `go_package` trong hướng dẫn phải được đổi cho khớp.

### Thao tác 4: Tạo thư mục mã nguồn

```bash
mkdir -p proto server client
```

Kiểm tra:

```bash
find . -maxdepth 1 -type d -print
```

Danh sách phải có `./proto`, `./server` và `./client`.

### Thao tác 5: Kiểm tra Go và Protocol Buffer Compiler

```bash
go version
protoc --version
```

Cả hai lệnh phải in ra phiên bản. Nếu chỉ `protoc` bị thiếu, cài Protocol Buffer Compiler rồi chạy lại thao tác này.

### Thao tác 6: Cài plugin sinh mã message

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Lệnh thành công thường không in gì. Kiểm tra executable:

```bash
ls "$(go env GOPATH)/bin/protoc-gen-go"
```

### Thao tác 7: Cài plugin sinh mã gRPC

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Kiểm tra executable:

```bash
ls "$(go env GOPATH)/bin/protoc-gen-go-grpc"
```

### Thao tác 8: Cho `protoc` tìm thấy hai plugin

Áp dụng cho terminal hiện tại:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Kiểm tra:

```bash
command -v protoc-gen-go
command -v protoc-gen-go-grpc
```

Cả hai lệnh phải trả về đường dẫn. Để cấu hình tồn tại sau khi đóng terminal, thêm dòng `export PATH="$PATH:$(go env GOPATH)/bin"` vào `~/.zshrc`.

### Thao tác 9: Tạo file định nghĩa gRPC

Tạo file trống:

```bash
touch proto/todo.proto
```

Mở `proto/todo.proto` bằng editor, copy toàn bộ đoạn `proto` tại [Bước 3](#bước-3-định-nghĩa-service-trong-file-proto), lưu file rồi kiểm tra:

```bash
protoc --proto_path=. --descriptor_set_out=/tmp/todo.pb proto/todo.proto
```

Không có thông báo lỗi nghĩa là cú pháp `.proto` hợp lệ.

### Thao tác 10: Sinh hai file Go tự động

```bash
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/todo.proto
```

Kiểm tra:

```bash
ls proto
```

Phải có đủ ba file:

```text
todo.proto
todo.pb.go
todo_grpc.pb.go
```

Nếu báo `program not found`, quay lại thao tác 6-8. Không tự viết hoặc chỉnh sửa hai file `.pb.go`.

### Thao tác 11: Tạo server

```bash
touch server/main.go
```

Mở `server/main.go`, copy toàn bộ đoạn Go tại [Bước 5](#bước-5-cài-đặt-grpc-server), sau đó lưu file.

Format riêng server:

```bash
gofmt -w server/main.go
```

### Thao tác 12: Tạo client

```bash
touch client/main.go
```

Mở `client/main.go`, copy toàn bộ đoạn Go tại [Bước 6](#bước-6-cài-đặt-cli-client), sau đó lưu file.

Format riêng client:

```bash
gofmt -w client/main.go
```

### Thao tác 13: Tải dependency

```bash
go mod tidy
```

Sau lệnh này phải có thêm `go.sum`. Kiểm tra:

```bash
ls go.mod go.sum
```

### Thao tác 14: Build toàn bộ project

```bash
go build ./...
```

Lệnh thành công thường không in gì. Nếu có lỗi, đọc file và số dòng được báo; sửa lỗi rồi chạy lại cho đến khi exit code bằng `0`.

### Thao tác 15: Chạy server trong terminal thứ nhất

```bash
cd /Users/quindart/go-learn/grpc-review
go run ./server
```

Giữ terminal này mở. Kết quả cần thấy:

```text
gRPC server listening on localhost:50051
```

### Thao tác 16: Tạo dữ liệu trong terminal thứ hai

Mở terminal mới:

```bash
cd /Users/quindart/go-learn/grpc-review
go run ./client create "Học gRPC"
go run ./client create "Viết gRPC client"
```

Hai todo phải có ID lần lượt là `1` và `2`, đồng thời có `completed=false`.

### Thao tác 17: Đọc dữ liệu

```bash
go run ./client get 1
go run ./client list
```

Lệnh `get` phải trả về todo số `1`; lệnh `list` phải hiển thị cả hai todo.

### Thao tác 18: Hoàn thành một todo

```bash
go run ./client complete 1
go run ./client get 1
```

Kết quả của cả hai lệnh phải có `completed=true`.

### Thao tác 19: Kiểm tra validation và mã lỗi

```bash
go run ./client create ""
go run ./client get 999
```

Lệnh đầu phải báo `InvalidArgument`; lệnh sau phải báo `NotFound`. Đây là bằng chứng server đang trả đúng mã lỗi gRPC, không chỉ trả chuỗi lỗi thông thường.

### Thao tác 20: Kiểm tra timeout khi server dừng

Quay lại terminal chạy server và nhấn `Ctrl+C`. Sau đó, trong terminal client chạy:

```bash
go run ./client list
```

Client phải báo lỗi kết nối hoặc `Unavailable` và tự kết thúc trong tối đa ba giây.

### Thao tác 21: Xác nhận đã hoàn thành

Khởi động lại server nếu cần và chạy lần cuối:

```bash
go build ./...
go run ./client list
```

Task được xem là hoàn thành khi:

- Project build thành công.
- Bốn RPC đều gọi được từ CLI.
- `InvalidArgument` và `NotFound` trả đúng trường hợp.
- Client có timeout khi server không phản hồi.
- Bạn biết file nào do mình viết và file nào do `protoc` sinh ra.

## Hướng dẫn thực hiện

### Bước 1: Khởi tạo Go module và cấu trúc thư mục

Từ thư mục gốc của repository, đi vào thư mục bài tập:

```bash
cd grpc-review
```

Khởi tạo một Go module. Hướng dẫn này dùng module path `example.com/grpc-review` để các import bên dưới thống nhất với nhau:

```bash
go mod init example.com/grpc-review
```

Tạo ba thư mục chính:

```bash
mkdir -p proto server client
```

Sau bước này, cấu trúc thư mục là:

```text
grpc-review/
├── README.md
├── go.mod
├── client/
├── proto/
└── server/
```

Checkpoint:

```bash
go env GOMOD
```

Lệnh trên phải trả về đường dẫn kết thúc bằng `grpc-review/go.mod`.

### Bước 2: Cài các plugin sinh mã Go cho `protoc`

#### 2.1. Hiểu vai trò của từng công cụ

`protoc` là trình biên dịch Protocol Buffers. Riêng với Go, `protoc` cần gọi thêm hai plugin:

- `protoc-gen-go`: sinh các kiểu Go cho message trong file `.proto`, ví dụ `Todo` và `CreateTodoRequest`.
- `protoc-gen-go-grpc`: sinh client stub, server interface và mã đăng ký gRPC service.

Hai plugin này là các chương trình độc lập. Chúng phải được cài đặt và nằm trong `PATH` để `protoc` tìm thấy.

#### 2.2. Kiểm tra công cụ nền

Kiểm tra Go và `protoc` trước khi cài plugin:

```bash
go version
protoc --version
```

Nếu cả hai lệnh đều in ra phiên bản, tiếp tục sang bước tiếp theo. Nếu `protoc` chưa tồn tại, hãy cài Protocol Buffer Compiler trước.

#### 2.3. Cài hai plugin

Chạy lần lượt:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

`@latest` yêu cầu Go tải phiên bản mới nhất và cho phép chạy `go install` ngay cả khi thư mục hiện tại chưa có `go.mod`.

Sau khi cài đặt, Go đặt các executable vào:

- Thư mục được cấu hình bởi `GOBIN`, nếu biến này có giá trị.
- `$(go env GOPATH)/bin`, nếu `GOBIN` đang để trống.

Kiểm tra cấu hình hiện tại:

```bash
go env GOBIN
go env GOPATH
```

#### 2.4. Thêm thư mục plugin vào `PATH`

Với cấu hình Go mặc định, thêm `GOPATH/bin` vào `PATH` của terminal hiện tại:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Trên macOS sử dụng Zsh, để cấu hình có hiệu lực cho những terminal mở sau này, thêm dòng sau vào `~/.zshrc`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Sau đó mở terminal mới hoặc nạp lại cấu hình:

```bash
source ~/.zshrc
```

Nếu `go env GOBIN` trả về một đường dẫn riêng, hãy thêm chính đường dẫn đó vào `PATH` thay cho `$(go env GOPATH)/bin`.

#### 2.5. Xác minh cài đặt

Kiểm tra terminal có tìm thấy hai plugin hay không:

```bash
command -v protoc-gen-go
command -v protoc-gen-go-grpc
```

Tiếp tục kiểm tra phiên bản:

```bash
protoc-gen-go --version
protoc-gen-go-grpc --version
```

Bước 2 hoàn thành khi cả bốn lệnh trên đều chạy thành công và hai lệnh `command -v` trả về đường dẫn đến executable.

#### 2.6. Lỗi thường gặp

Nếu lệnh sinh code sau này báo lỗi tương tự:

```text
protoc-gen-go: program not found or is not executable
```

hãy kiểm tra theo thứ tự:

1. Chạy lại hai lệnh `go install`.
2. Kiểm tra executable có trong `$(go env GOPATH)/bin` hoặc `GOBIN` hay không.
3. Kiểm tra đường dẫn đó đã có trong `PATH` bằng `command -v`.
4. Mở terminal mới sau khi chỉnh `~/.zshrc`.

Tài liệu tham khảo:

- [gRPC Go Quick start](https://grpc.io/docs/languages/go/quickstart/)
- [Protocol Buffers Go Generated Code Guide](https://protobuf.dev/reference/go/go-generated/)

### Bước 3: Định nghĩa service trong file `.proto`

Tạo file `proto/todo.proto` với nội dung sau:

```proto
syntax = "proto3";

package todo.v1;

option go_package = "example.com/grpc-review/proto;todo";

service TodoService {
  rpc CreateTodo(CreateTodoRequest) returns (Todo);
  rpc GetTodo(GetTodoRequest) returns (Todo);
  rpc ListTodos(ListTodosRequest) returns (ListTodosResponse);
  rpc CompleteTodo(CompleteTodoRequest) returns (Todo);
}

message Todo {
  int64 id = 1;
  string title = 2;
  bool completed = 3;
}

message CreateTodoRequest {
  string title = 1;
}

message GetTodoRequest {
  int64 id = 1;
}

message ListTodosRequest {}

message ListTodosResponse {
  repeated Todo todos = 1;
}

message CompleteTodoRequest {
  int64 id = 1;
}
```

Các phần quan trọng:

- `syntax = "proto3"` chọn cú pháp Protocol Buffers phiên bản 3.
- `package todo.v1` tạo namespace ở cấp Protocol Buffers.
- `go_package` khai báo import path và package name cho mã Go được sinh ra.
- Mỗi RPC nhận đúng một request message và trả về đúng một response message.
- `repeated Todo` tương ứng với một danh sách todo trong Go.

Checkpoint:

```bash
protoc --proto_path=. --descriptor_set_out=/tmp/todo.pb proto/todo.proto
```

Nếu file hợp lệ, lệnh kết thúc mà không in lỗi cú pháp.

### Bước 4: Sinh mã Go từ file `.proto`

Đứng tại thư mục `grpc-review` và chạy:

```bash
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  proto/todo.proto
```

Các option có ý nghĩa:

- `--go_out=.` gọi `protoc-gen-go` để sinh kiểu dữ liệu cho các message.
- `--go-grpc_out=.` gọi `protoc-gen-go-grpc` để sinh client và server API.
- `paths=source_relative` đặt file được sinh cạnh file `.proto`, thay vì tạo cây thư mục theo import path.

Kết quả mong đợi:

```text
proto/
├── todo.proto
├── todo.pb.go
└── todo_grpc.pb.go
```

Không sửa thủ công hai file `.pb.go`; hãy sửa `.proto` rồi chạy lại `protoc`.

Checkpoint:

```bash
test -f proto/todo.pb.go
test -f proto/todo_grpc.pb.go
```

### Bước 4.1: Hiểu cách client và server dùng chung generated code

Client và server không tự đoán cấu trúc dữ liệu của nhau. Cả hai cùng import package Go được sinh từ một file `proto/todo.proto` duy nhất:

```text
                         proto/todo.proto
                                 |
              +------------------+------------------+
              |                                     |
       protoc-gen-go                       protoc-gen-go-grpc
              |                                     |
              v                                     v
      proto/todo.pb.go                   proto/todo_grpc.pb.go
       message structs              client stub + server interface
              |                                     |
              +------------------+------------------+
                                 |
                  +--------------+--------------+
                  |                             |
              server/main.go                client/main.go
```

#### File `todo.pb.go` cung cấp kiểu dữ liệu chung

Từ các message trong `.proto`, plugin `protoc-gen-go` sinh các struct tương ứng:

```text
Todo                    -> pb.Todo
CreateTodoRequest       -> pb.CreateTodoRequest
GetTodoRequest          -> pb.GetTodoRequest
ListTodosRequest        -> pb.ListTodosRequest
ListTodosResponse       -> pb.ListTodosResponse
CompleteTodoRequest     -> pb.CompleteTodoRequest
```

Client dùng các kiểu này để tạo request:

```go
request := &pb.CreateTodoRequest{Title: "Học gRPC"}
```

Server nhận đúng cùng một kiểu:

```go
func (s *todoServer) CreateTodo(
	ctx context.Context,
	req *pb.CreateTodoRequest,
) (*pb.Todo, error) {
	// req có cùng schema mà client đã dùng.
}
```

Protocol Buffers dùng số thứ tự field, ví dụ `title = 1`, để mã hóa request thành bytes. Vì client và server được sinh từ cùng schema nên cả hai hiểu field số `1` là `title`.

#### File `todo_grpc.pb.go` cung cấp API cho server

Plugin `protoc-gen-go-grpc` sinh interface tương tự:

```go
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*Todo, error)
	GetTodo(context.Context, *GetTodoRequest) (*Todo, error)
	ListTodos(context.Context, *ListTodosRequest) (*ListTodosResponse, error)
	CompleteTodo(context.Context, *CompleteTodoRequest) (*Todo, error)
}
```

`server/main.go` phải:

1. Khai báo `package main`.
2. Tạo một struct server.
3. Embed `pb.UnimplementedTodoServiceServer`.
4. Viết đủ bốn method với đúng chữ ký generated interface yêu cầu.
5. Đăng ký struct bằng `pb.RegisterTodoServiceServer`.
6. Gọi `grpcServer.Serve` để bắt đầu nhận request.

Phần liên kết quan trọng:

```go
grpcServer := grpc.NewServer()
todoService := newTodoServer()
pb.RegisterTodoServiceServer(grpcServer, todoService)
grpcServer.Serve(listener)
```

Nếu chữ ký method trên server bị sai, `go build ./...` sẽ báo struct chưa implement `pb.TodoServiceServer`.

#### File `todo_grpc.pb.go` cung cấp API cho client

File generated cũng sinh `TodoServiceClient`. `client/main.go` không gọi trực tiếp struct server; nó tạo một stub gắn với connection:

```go
conn, err := grpc.NewClient(
	"localhost:50051",
	grpc.WithTransportCredentials(insecure.NewCredentials()),
)
if err != nil {
	return err
}
defer conn.Close()

client := pb.NewTodoServiceClient(conn)
```

Sau đó client gọi method trên stub:

```go
todo, err := client.CreateTodo(ctx, &pb.CreateTodoRequest{
	Title: "Học gRPC",
})
```

Stub generated thực hiện các việc còn lại:

1. Mã hóa `CreateTodoRequest` thành Protocol Buffers.
2. Gửi request đến method `/todo.v1.TodoService/CreateTodo` qua HTTP/2.
3. Nhận bytes response từ server.
4. Giải mã response thành `*pb.Todo`.
5. Chuyển gRPC status thành Go error nếu server trả lỗi.

#### Luồng đầy đủ của một request

```text
client/main.go
    |
    | client.CreateTodo(ctx, request)
    v
generated client stub
    |
    | protobuf bytes + method name
    v
HTTP/2 connection tới localhost:50051
    |
    v
generated server handler
    |
    v
todoServer.CreateTodo(...)
    |
    | trả về *pb.Todo hoặc gRPC status
    v
generated server handler -> HTTP/2 -> generated client stub
    |
    v
client nhận *pb.Todo hoặc error
```

#### Những giá trị bắt buộc phải đồng bộ

| Thành phần | Giá trị trong bài tập | Tại sao |
| --- | --- | --- |
| Go module | `example.com/grpc-review` | Là tiền tố import nội bộ |
| `go_package` | `example.com/grpc-review/proto;todo` | Sinh package có import path đúng |
| Import ở client/server | `pb "example.com/grpc-review/proto"` | Dùng chung generated package |
| Proto package | `todo.v1` | Tạo tên service trên wire |
| Server address | `localhost:50051` | Client phải kết nối đúng nơi server listen |

Kiểm tra module path hiện tại:

```bash
go list -m
```

Nếu kết quả chưa phải `example.com/grpc-review`, sửa trước khi viết client và server:

```bash
go mod edit -module=example.com/grpc-review
```

#### Quy trình bắt buộc mỗi khi sửa `.proto`

Ví dụ bạn thêm một RPC hoặc message mới, luôn làm theo thứ tự:

1. Sửa và lưu `proto/todo.proto`.
2. Chạy lại cả hai generator:

   ```bash
   protoc \
     --go_out=. \
     --go_opt=paths=source_relative \
     --go-grpc_out=. \
     --go-grpc_opt=paths=source_relative \
     proto/todo.proto
   ```

3. Kiểm tra `todo.pb.go` và `todo_grpc.pb.go` đã được cập nhật.
4. Implement method mới trong `server/main.go`.
5. Gọi method mới từ `client/main.go`.
6. Chạy:

   ```bash
   go mod tidy
   gofmt -w server/main.go client/main.go
   go build ./...
   ```

Không copy riêng file `.pb.go` từ project khác và không sửa generated code bằng tay. Nguồn sự thật duy nhất là file `.proto`.

### Bước 5: Cài đặt gRPC server

Tạo file `server/main.go`:

```go
package main

import (
	"context"
	"log"
	"net"
	"sort"
	"strings"
	"sync"

	pb "example.com/grpc-review/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type todoServer struct {
	pb.UnimplementedTodoServiceServer

	mu     sync.RWMutex
	nextID int64
	todos  map[int64]*pb.Todo
}

func newTodoServer() *todoServer {
	return &todoServer{
		todos: make(map[int64]*pb.Todo),
	}
}

func cloneTodo(todo *pb.Todo) *pb.Todo {
	return &pb.Todo{
		Id:        todo.GetId(),
		Title:     todo.GetTitle(),
		Completed: todo.GetCompleted(),
	}
}

func (s *todoServer) CreateTodo(_ context.Context, req *pb.CreateTodoRequest) (*pb.Todo, error) {
	title := strings.TrimSpace(req.GetTitle())
	if title == "" {
		return nil, status.Error(codes.InvalidArgument, "title must not be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	todo := &pb.Todo{
		Id:        s.nextID,
		Title:     title,
		Completed: false,
	}
	s.todos[todo.Id] = todo

	return cloneTodo(todo), nil
}

func (s *todoServer) GetTodo(_ context.Context, req *pb.GetTodoRequest) (*pb.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, ok := s.todos[req.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo %d not found", req.GetId())
	}

	return cloneTodo(todo), nil
}

func (s *todoServer) ListTodos(_ context.Context, _ *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*pb.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, cloneTodo(todo))
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].GetId() < todos[j].GetId()
	})

	return &pb.ListTodosResponse{Todos: todos}, nil
}

func (s *todoServer) CompleteTodo(_ context.Context, req *pb.CompleteTodoRequest) (*pb.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[req.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo %d not found", req.GetId())
	}

	todo.Completed = true
	return cloneTodo(todo), nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, newTodoServer())

	log.Println("gRPC server listening on localhost:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
```

Điểm cần hiểu:

- `UnimplementedTodoServiceServer` giúp server tương thích khi service được mở rộng.
- `RegisterTodoServiceServer` gắn phần cài đặt của bạn vào gRPC server.
- `sync.RWMutex` bảo vệ map khi nhiều RPC chạy đồng thời.
- `status.Error` và `status.Errorf` trả về mã lỗi gRPC có nghĩa rõ ràng.
- `cloneTodo` tránh trả ra con trỏ đang được lưu và thay đổi trong map.

### Bước 6: Cài đặt CLI client

Tạo file `client/main.go`:

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "example.com/grpc-review/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func usage() error {
	return errors.New("usage: client create <title> | get <id> | list | complete <id>")
}

func parseID(value string) (int64, error) {
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid todo id %q", value)
	}
	return id, nil
}

func printTodo(todo *pb.Todo) {
	fmt.Printf("id=%d title=%q completed=%t\n", todo.GetId(), todo.GetTitle(), todo.GetCompleted())
}

func run() error {
	if len(os.Args) < 2 {
		return usage()
	}

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("create gRPC client: %w", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch os.Args[1] {
	case "create":
		if len(os.Args) != 3 {
			return usage()
		}
		todo, err := client.CreateTodo(ctx, &pb.CreateTodoRequest{Title: os.Args[2]})
		if err != nil {
			return err
		}
		printTodo(todo)

	case "get":
		if len(os.Args) != 3 {
			return usage()
		}
		id, err := parseID(os.Args[2])
		if err != nil {
			return err
		}
		todo, err := client.GetTodo(ctx, &pb.GetTodoRequest{Id: id})
		if err != nil {
			return err
		}
		printTodo(todo)

	case "list":
		if len(os.Args) != 2 {
			return usage()
		}
		response, err := client.ListTodos(ctx, &pb.ListTodosRequest{})
		if err != nil {
			return err
		}
		for _, todo := range response.GetTodos() {
			printTodo(todo)
		}

	case "complete":
		if len(os.Args) != 3 {
			return usage()
		}
		id, err := parseID(os.Args[2])
		if err != nil {
			return err
		}
		todo, err := client.CompleteTodo(ctx, &pb.CompleteTodoRequest{Id: id})
		if err != nil {
			return err
		}
		printTodo(todo)

	default:
		return usage()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			fmt.Fprintf(os.Stderr, "gRPC error: code=%s message=%s\n", grpcStatus.Code(), grpcStatus.Message())
		} else {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
}
```

Client sử dụng kết nối không mã hóa vì đây là bài tập chạy local. Trong môi trường thật, hãy cấu hình TLS thay cho `insecure.NewCredentials()`.

### Bước 7: Tải dependency, format và kiểm tra build

Từ thư mục `grpc-review`, chạy:

```bash
go mod tidy
gofmt -w server/main.go client/main.go
go build ./...
```

`go mod tidy` sẽ thêm các module cần thiết vào `go.mod` và `go.sum`, bao gồm gRPC và Protocol Buffers runtime.

Checkpoint:

- `go build ./...` kết thúc với exit code `0`.
- Không có lỗi import package `example.com/grpc-review/proto`.
- Không sửa trực tiếp mã nguồn generated để xử lý lỗi build.

### Bước 8: Chạy và kiểm tra đầy đủ các RPC

Mở terminal thứ nhất, đi vào `grpc-review` và chạy server:

```bash
go run ./server
```

Kết quả mong đợi:

```text
gRPC server listening on localhost:50051
```

Giữ server chạy. Mở terminal thứ hai trong cùng thư mục để gọi client.

Tạo hai todo:

```bash
go run ./client create "Học gRPC"
go run ./client create "Viết gRPC client"
```

Kết quả tương tự:

```text
id=1 title="Học gRPC" completed=false
id=2 title="Viết gRPC client" completed=false
```

Lấy todo theo ID:

```bash
go run ./client get 1
```

Liệt kê toàn bộ todo:

```bash
go run ./client list
```

Đánh dấu todo đầu tiên đã hoàn thành:

```bash
go run ./client complete 1
go run ./client get 1
```

Kết quả cuối phải có `completed=true`.

Kiểm tra lỗi `InvalidArgument`:

```bash
go run ./client create ""
```

Kết quả cần chứa:

```text
code=InvalidArgument
```

Kiểm tra lỗi `NotFound`:

```bash
go run ./client get 999
```

Kết quả cần chứa:

```text
code=NotFound
```

Dừng server bằng `Ctrl+C`, sau đó gọi lại client:

```bash
go run ./client list
```

Client phải kết thúc sau tối đa ba giây và hiển thị lỗi kết nối hoặc `Unavailable`.

### Checklist hoàn thành bài tập

- [ ] `go.mod` dùng module path khớp với `go_package` và các import.
- [ ] Hai plugin sinh mã Go đã có trong `PATH`.
- [ ] `protoc` sinh được `todo.pb.go` và `todo_grpc.pb.go`.
- [ ] `go build ./...` chạy thành công.
- [ ] Server lắng nghe tại `localhost:50051`.
- [ ] `CreateTodo`, `GetTodo`, `ListTodos` và `CompleteTodo` hoạt động.
- [ ] Tiêu đề rỗng trả về `InvalidArgument`.
- [ ] ID không tồn tại trả về `NotFound`.
- [ ] Client có timeout ba giây.
- [ ] Bạn giải thích được vai trò của hai file Go được sinh tự động.

## Yêu cầu chức năng

Ứng dụng cung cấp bốn RPC:

```proto
service TodoService {
  rpc CreateTodo(CreateTodoRequest) returns (Todo);
  rpc GetTodo(GetTodoRequest) returns (Todo);
  rpc ListTodos(ListTodosRequest) returns (ListTodosResponse);
  rpc CompleteTodo(CompleteTodoRequest) returns (Todo);
}
```

Mỗi công việc có cấu trúc:

```proto
message Todo {
  int64 id = 1;
  string title = 2;
  bool completed = 3;
}
```

Bạn tự thiết kế các request và response message còn thiếu trong file `.proto`.

### 1. CreateTodo

- Nhận vào tiêu đề công việc.
- Tiêu đề không được để trống.
- Server tự sinh `id` tăng dần.
- Công việc mới có `completed = false`.
- Nếu tiêu đề rỗng, trả về `codes.InvalidArgument`.

### 2. GetTodo

- Tìm một công việc theo `id`.
- Nếu không tồn tại, trả về `codes.NotFound`.

### 3. ListTodos

- Trả về toàn bộ công việc hiện có.
- Nếu chưa có công việc nào, trả về danh sách rỗng.

### 4. CompleteTodo

- Tìm công việc theo `id` và chuyển `completed` thành `true`.
- Trả về công việc sau khi cập nhật.
- Nếu không tồn tại, trả về `codes.NotFound`.

## Lưu trữ dữ liệu

Không cần sử dụng database. Server có thể lưu dữ liệu trong bộ nhớ:

```go
map[int64]*pb.Todo
```

Nếu muốn chạy server với nhiều goroutine an toàn, hãy bảo vệ map bằng `sync.RWMutex`.

## CLI client

Client cần hỗ trợ các lệnh tương tự:

```bash
go run ./client create "Học gRPC"
go run ./client get 1
go run ./client list
go run ./client complete 1
```

Mỗi request phải có timeout ba giây:

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()
```

Client cần hiển thị thông báo dễ hiểu khi:

- Server chưa chạy.
- Request hết thời gian chờ.
- Server trả về `InvalidArgument` hoặc `NotFound`.

## Cấu trúc thư mục gợi ý

```text
grpc-review/
├── go.mod
├── proto/
│   └── todo.proto
├── server/
│   └── main.go
└── client/
    └── main.go
```

Bạn có thể thay đổi cấu trúc này nếu vẫn tách biệt rõ phần server, client và generated code.

## Các trường hợp cần kiểm tra

1. Tạo todo `Học gRPC` thành công.
2. Tạo todo với tiêu đề rỗng và nhận `InvalidArgument`.
3. Lấy một todo tồn tại.
4. Lấy một todo không tồn tại và nhận `NotFound`.
5. Đánh dấu một todo là đã hoàn thành.
6. Liệt kê được cả todo đã hoàn thành và chưa hoàn thành.
7. Client báo lỗi rõ ràng khi server chưa chạy.

## Tiêu chí hoàn thành

Bài tập hoàn thành khi bạn có thể:

- Khởi động server tại `localhost:50051`.
- Tạo ít nhất hai todo bằng CLI client.
- Lấy một todo theo `id`.
- Liệt kê toàn bộ todo.
- Đánh dấu một todo là đã hoàn thành.
- Xử lý đúng hai mã lỗi `InvalidArgument` và `NotFound`.
- Giải thích được vai trò của file `.proto`, generated code, server stub và client stub.

## Thử thách thêm

Sau khi hoàn thành phần cơ bản, bạn có thể thử:

- Thêm RPC `DeleteTodo`.
- Cho phép `ListTodos` lọc theo trạng thái hoàn thành.
- Thêm server-streaming RPC để server gửi từng todo về client.
- Viết unit test cho server bằng `bufconn`.
- Bật gRPC reflection và gọi service bằng `grpcurl`.

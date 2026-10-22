# Go Learn

Repository này tập hợp các bài thực hành và dự án nhỏ để học ngôn ngữ Go, từ cú pháp cơ bản đến API, cơ sở dữ liệu, đồng thời và gRPC. Hầu hết thư mục là chương trình/module độc lập; hãy vào thư mục tương ứng trước khi chạy `go run` hoặc các lệnh Go.

## Mục lục

- [Cấu trúc và nội dung](#cấu-trúc-và-nội-dung)
- [Các chủ đề có thể học](#các-chủ-đề-có-thể-học)
- [Cách chạy một bài](#cách-chạy-một-bài)

## Cấu trúc và nội dung

### Cú pháp và kiểu dữ liệu cơ bản

| Thư mục | Nội dung | Có thể học |
| --- | --- | --- |
| `go-basic-learn/` | Ví dụ chia chương trình thành các package `book`, `author`, `speaker`, `user`. | Package, import, tổ chức mã thành nhiều file và kiểu dữ liệu có cấu trúc. |
| `go-if-else/` | Ví dụ điều kiện. | Rẽ nhánh `if/else` và biểu thức điều kiện. |
| `go-switch-case/` | Ví dụ `switch`. | Chọn nhánh theo nhiều trường hợp. |
| `go-loop/` | Ví dụ vòng lặp. | Cách dùng `for` trong Go. |
| `go-arrays-handler/` | Thao tác với mảng. | Khai báo và duyệt mảng. |
| `go-string/` | Ví dụ làm việc với chuỗi. | Kiểu chuỗi và các thao tác cơ bản. |
| `go-custom-type/` | Khai báo kiểu chuỗi riêng và method. | Kiểu tùy chỉnh và method receiver. |
| `go-any-type-allowed/` | Ví dụ dùng kiểu tổng quát `any`. | Giá trị có kiểu động và cách khai báo dữ liệu linh hoạt. |
| `go-struct/` | Struct người dùng trong package `user`. | Struct, constructor, truyền giá trị/con trỏ và method. |
| `go-struct-tag/` | Ví dụ struct tag. | Gắn metadata cho trường struct, thường dùng khi mã hóa dữ liệu. |
| `go-pointer/` | Ví dụ con trỏ. | Địa chỉ bộ nhớ, dereference và cập nhật giá trị qua con trỏ. |
| `go-interface/` | Ví dụ interface và package `Todo`. | Khai báo hành vi bằng interface và hiện thực hóa hành vi. |
| `go-list/` | Ví dụ list cùng package `list`. | Tách package và mô hình hóa danh sách. |
| `go-thirty/` | Bài thực hành Go tổng hợp. | Ôn tập và kết hợp các cấu trúc ngôn ngữ cơ bản. |
| `go-practice/` | Module thực hành tổng hợp. | Đọc, chạy và thử chỉnh sửa một chương trình Go có `go.mod`. |
| `go-finbonacci/` | Tính dãy Fibonacci. | Hàm, vòng lặp và bài toán số học cơ bản. |
| `go-make-function/` | Ví dụ hàm và function value. | Truyền hàm, trả về hàm và closure. |
| `go-function-deep-dive/` | Hàm như giá trị, gồm package `functionsarevalue`. | Function value, callback và higher-order function. |

### Go module, phạm vi biến và môi trường

| Thư mục/tệp | Nội dung | Có thể học |
| --- | --- | --- |
| `ly-thuyet/go-mod-init/` | Ví dụ khởi tạo Go module. | Vai trò của `go.mod` và module path. |
| `ly-thuyet/variable-scope/` | Ví dụ phạm vi biến. | Scope của biến trong package, hàm và block. |
| `ly-thuyet/variable-type/` | Ví dụ kiểu biến. | Các kiểu dữ liệu và cách khai báo biến. |
| `ly-thuyet/tips-go-env.md` | Ghi chú về môi trường Go. | Một số thiết lập và lệnh hữu ích khi làm việc với Go. |

### Tệp, JSON và ứng dụng dòng lệnh

| Thư mục | Nội dung | Có thể học |
| --- | --- | --- |
| `go-os-reading/` | Ví dụ đọc dữ liệu từ tệp. | Đọc tệp và xử lý lỗi I/O. |
| `go-os-write/` | Ghi dữ liệu số dư vào tệp. | Ghi tệp và lưu trạng thái đơn giản. |
| `go-package/` | Chương trình chính và package `fileos`. | Tạo package riêng, import và tách phần đọc/ghi tệp. |
| `go-homework-take-note-app/` | Ứng dụng ghi chú, package `note`, dữ liệu mẫu JSON. | Struct, kiểm tra dữ liệu đầu vào và lưu dữ liệu dạng tệp. |
| `go-task-homework-financial-cal/` | Bài tập tính toán tài chính, có tệp kết quả. | Nhập liệu, hàm tính toán và xuất kết quả. |
| `go-investment-app/` | Ví dụ ứng dụng tính toán đầu tư. | Nhập dữ liệu và tính toán tài chính cơ bản. |
| `go-price-calculator/` | Tính giá sau thuế; có package `conversion`, `filemanager`, `prices` và các tệp kết quả JSON. | Chia package theo trách nhiệm, đọc giá từ tệp, chuyển đổi kiểu và xuất JSON. |
| `go-price-calculator-concurrency/` | Phiên bản tính giá có xử lý đồng thời. | Mở rộng chương trình tính giá bằng goroutine và đồng bộ hóa kết quả. |

### Đồng thời và so sánh với JavaScript

| Thư mục | Nội dung | Có thể học |
| --- | --- | --- |
| `go-concurrency/` | Ví dụ goroutine, channel và giới hạn luồng xử lý. | Mô hình đồng thời của Go và phối hợp các tác vụ. |
| `go-routine-ex/` | Ví dụ goroutine, worker và `WaitGroup`; có thể có tệp trace. | Khởi chạy goroutine, chờ hoàn tất và quan sát thực thi. |
| `interview-days/day1-wait-group/` | Bài luyện tập `WaitGroup`. | Đợi nhóm goroutine hoàn thành an toàn. |
| `interview-days/day2/` | Thư mục bài luyện tập ngày 2. | Dùng làm nơi ôn tập các chủ đề phỏng vấn Go. |
| `compare-nodejs-golang-goroutine/` | Ví dụ event loop JavaScript và goroutine Go. | So sánh mô hình xử lý bất đồng bộ giữa Node.js và Go. |

### Web API, cơ sở dữ liệu và dịch vụ

| Thư mục | Nội dung | Có thể học |
| --- | --- | --- |
| `app-event-booking/` | API đặt sự kiện với router, model, middleware xác thực, JWT, mã hóa, DB và các request mẫu `.http`. | Xây dựng REST API, CRUD, xác thực người dùng, JWT, kết nối SQLite và kiểm thử endpoint. |
| `fullstack-todo/` | Ứng dụng Todo gồm `be/` (Go, model, DB, request mẫu) và `fe/` (giao diện web). | Kết nối frontend/backend và xây dựng CRUD với lưu trữ dữ liệu. |
| `go-crm-shop-dev/` | Dịch vụ CRM/shop: `cmd/`, `api/`, `internal/`, `config/`. | Tổ chức ứng dụng Go theo lớp, xây dựng HTTP API, middleware và cấu hình. Xem thêm README trong thư mục này. |
| `user-management-api/` | Dịch vụ quản lý người dùng với `cmd/` và các package trong `internal/` như handler, service, repository, middleware, validation, config và models. | Phân lớp API, kiểm tra dữ liệu, xử lý nghiệp vụ và tách truy cập dữ liệu. |
| `test/` | Ví dụ package todo và chương trình thử nghiệm. | Làm quen với cấu trúc module và viết/đọc mã thử nghiệm. |

### gRPC

| Thư mục | Nội dung | Có thể học |
| --- | --- | --- |
| `grpc-review/` | Todo gRPC gồm `proto/`, `server/`, `client/` và README hướng dẫn chi tiết. | Định nghĩa Protocol Buffers, sinh mã Go, cài đặt gRPC server/client, RPC, status code và timeout. |

## Các chủ đề có thể học

1. **Nền tảng Go:** biến, kiểu dữ liệu, điều kiện, vòng lặp, mảng, chuỗi, struct, con trỏ và interface.
2. **Tổ chức mã:** package, module, import, hàm như giá trị và chia chương trình thành các thành phần nhỏ.
3. **Làm việc với dữ liệu:** đọc/ghi tệp, JSON, xử lý lỗi và lưu dữ liệu đơn giản.
4. **Đồng thời:** goroutine, channel, `WaitGroup` và bảo vệ dữ liệu chia sẻ.
5. **Ứng dụng thực tế:** CLI, tính toán, REST API, xác thực, database, frontend/backend và gRPC.

## Cách chạy một bài

Các thư mục có `go.mod` thường là module độc lập. Ví dụ:

```bash
cd go-concurrency
go run .
```

Với một số thư mục không có `go.mod`, có thể chạy trực tiếp tệp chính, chẳng hạn:

```bash
go run main.go
```

Hãy xem `go.mod` và README bên trong thư mục (nếu có) để biết lệnh chạy, dependency và cấu hình riêng. Một số ứng dụng cần database, server hoặc cấu hình môi trường trước khi chạy.

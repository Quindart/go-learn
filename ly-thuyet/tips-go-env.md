## `go env`

`go env` là lệnh dùng để xem các biến môi trường và cấu hình mà Go đang sử dụng.

```bash
go env
```

Một số biến thường gặp:

- `GOROOT`: thư mục cài đặt Go.
- `GOPATH`: thư mục làm việc của Go, thường chứa cache và các package tải về.
- `GOMOD`: đường dẫn đến file `go.mod` hiện tại.
- `GOPROXY`: nơi Go tải các module.
- `GOOS`: hệ điều hành mục tiêu, ví dụ `darwin`, `linux`, `windows`.
- `GOARCH`: kiến trúc CPU, ví dụ `amd64`, `arm64`.
- `GOVERSION`: phiên bản Go đang sử dụng.

Xem một biến cụ thể:

```bash
go env GOPATH
go env GOROOT
go env GOVERSION
```

Đặt giá trị cho một biến:

```bash
go env -w GOPROXY=https://proxy.golang.org,direct
```

Xóa cấu hình đã đặt:

```bash
go env -u GOPROXY
```

Kiểm tra nhanh môi trường Go:

```bash
go version
go env GOPATH GOROOT GOOS GOARCH
```

Lưu ý: `go env` chỉ hiển thị hoặc thay đổi cấu hình của Go; nó không tự tạo project hay tải dependency.

## Config trên máy hiện tại

Máy của bạn là Mac chạy chip Apple Silicon (`arm64`), macOS (`darwin`), Go phiên bản `go1.26.6`.

### Nền tảng và phiên bản

| Biến | Giá trị hiện tại | Mô tả |
| --- | --- | --- |
| `GOOS` | `darwin` | Hệ điều hành mục tiêu khi biên dịch. |
| `GOARCH` | `arm64` | Kiến trúc CPU mục tiêu. |
| `GOARM64` | `v8.0` | Mức kiến trúc ARM64 được sử dụng. |
| `GOHOSTOS` | `darwin` | Hệ điều hành của máy đang chạy Go. |
| `GOHOSTARCH` | `arm64` | Kiến trúc của máy đang chạy Go. |
| `GOVERSION` | `go1.26.6` | Phiên bản Go hiện tại. |
| `GOEXE` | rỗng | Hậu tố file thực thi; macOS không cần hậu tố. |

### Thư mục và module

| Biến | Giá trị hiện tại | Mô tả |
| --- | --- | --- |
| `GOROOT` | `/usr/local/go` | Thư mục cài đặt bộ công cụ Go. |
| `GOPATH` | `/Users/quindart/go` | Thư mục làm việc mặc định của Go. |
| `GOBIN` | rỗng | Thư mục chứa binary cài bằng `go install`; rỗng nghĩa là dùng mặc định. |
| `GOMOD` | `/dev/null` | File module đang được dùng; giá trị này cho biết bạn chạy lệnh ngoài một project có `go.mod`. |
| `GOMODCACHE` | `/Users/quindart/go/pkg/mod` | Nơi lưu các module dependency đã tải. |
| `GOWORK` | rỗng | Không sử dụng Go workspace (`go.work`) hiện tại. |
| `GOCACHE` | `/Users/quindart/Library/Caches/go-build` | Cache kết quả biên dịch để build nhanh hơn. |
| `GOTMPDIR` | rỗng | Thư mục tạm; Go sẽ tự chọn thư mục tạm của macOS. |
| `GOENV` | `/Users/quindart/Library/Application Support/go/env` | File lưu các config được đặt bằng `go env -w`. |

### Tải dependency và bảo mật module

| Biến | Giá trị hiện tại | Mô tả |
| --- | --- | --- |
| `GOPROXY` | `https://proxy.golang.org,direct` | Tải module từ proxy trước, nếu không có thì tải trực tiếp. |
| `GOSUMDB` | `sum.golang.org` | Máy chủ kiểm tra checksum của module. |
| `GOAUTH` | `netrc` | Cách Go lấy thông tin xác thực khi truy cập module riêng tư. |
| `GOPRIVATE` | rỗng | Chưa khai báo module private. |
| `GONOPROXY` | rỗng | Chưa cấu hình ngoại lệ cho proxy. |
| `GONOSUMDB` | rỗng | Chưa cấu hình module bỏ qua checksum database. |
| `GOINSECURE` | rỗng | Không cho phép module tải qua kết nối không an toàn. |
| `GOVCS` | rỗng | Dùng chính sách mặc định cho các hệ thống version control. |

### Biên dịch và C/C++

| Biến | Giá trị hiện tại | Mô tả |
| --- | --- | --- |
| `CC` | `clang` | Trình biên dịch C được dùng khi build với CGO. |
| `CXX` | `clang++` | Trình biên dịch C++. |
| `AR` | `ar` | Công cụ tạo thư viện tĩnh. |
| `GCCGO` | `gccgo` | Tên trình biên dịch GCC Go thay thế. |
| `CGO_ENABLED` | `1` | CGO đang bật, cho phép Go gọi mã C/C++. |
| `CGO_CFLAGS` | `-O2 -g` | Tùy chọn biên dịch mã C. `-O2` tối ưu, `-g` tạo thông tin debug. |
| `CGO_CPPFLAGS` | rỗng | Tùy chọn cho bộ tiền xử lý C/C++. |
| `CGO_CXXFLAGS` | `-O2 -g` | Tùy chọn biên dịch mã C++. |
| `CGO_FFLAGS` | `-O2 -g` | Tùy chọn biên dịch mã Fortran. |
| `CGO_LDFLAGS` | `-O2 -g` | Tùy chọn liên kết thư viện native. |
| `GOGCCFLAGS` | `-fPIC -arch arm64 -pthread ...` | Các cờ chi tiết Go tự truyền cho trình biên dịch native trên máy bạn. |
| `PKG_CONFIG` | `pkg-config` | Công cụ tìm cờ và thư viện native cho CGO. |

### Config chung và công cụ nội bộ

| Biến | Giá trị hiện tại | Mô tả |
| --- | --- | --- |
| `GO111MODULE` | rỗng | Go tự quyết định chế độ module; đây là lựa chọn được khuyến nghị. |
| `GOFLAGS` | rỗng | Chưa đặt cờ mặc định cho các lệnh `go`. |
| `GODEBUG` | rỗng | Chưa bật tùy chọn debug runtime nào. |
| `GOFIPS140` | `off` | Chế độ FIPS 140 đang tắt. |
| `GOEXPERIMENT` | rỗng | Không bật thử nghiệm compiler/runtime nào. |
| `GOTOOLCHAIN` | `auto` | Go tự chọn toolchain phù hợp khi module yêu cầu phiên bản khác. |
| `GOTOOLDIR` | `/usr/local/go/pkg/tool/darwin_arm64` | Thư mục chứa các công cụ build nội bộ của Go. |
| `GOCACHEPROG` | rỗng | Không dùng chương trình cache build bên ngoài. |
| `GOTELEMETRY` | `local` | Telemetry được lưu cục bộ. |
| `GOTELEMETRYDIR` | `/Users/quindart/Library/Application Support/go/telemetry` | Thư mục lưu dữ liệu telemetry. |

Nhìn chung, cấu hình hiện tại phù hợp với Mac Apple Silicon. Điểm cần nhớ là `GOMOD=/dev/null` chỉ xuất hiện vì lệnh được chạy tại thư mục home (`~`), không phải bên trong một project Go có file `go.mod`.

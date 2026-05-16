# Go Packages & Modules — Bilingual Theory | Lý thuyết Song ngữ

> **References / Tài liệu tham khảo:**
> - [go.dev/doc/code](https://go.dev/doc/code) — How to Write Go Code
> - [go.dev/ref/spec#Packages](https://go.dev/ref/spec#Packages) — Go Language Spec
> - [go.dev/ref/mod](https://go.dev/ref/mod) — Go Modules Reference
> - [go.dev/blog/using-go-modules](https://go.dev/blog/using-go-modules)

---

## TABLE OF CONTENTS | MỤC LỤC

**PART 1 — PACKAGES**
1. [Package Declaration](#1-package-declaration--khai-báo-package)
2. [Naming Conventions](#2-naming-conventions--quy-tắc-đặt-tên)
3. [Exported vs Unexported](#3-exported-vs-unexported--xuất-ra-ngoài-vs-nội-bộ)
4. [Import Paths & Aliasing](#4-import-paths--aliasing--đường-dẫn-import)
5. [Blank Identifier `_`](#5-blank-identifier---định-danh-trống)
6. [The `init()` Function](#6-the-init-function--hàm-init)
7. [Internal Packages](#7-internal-packages--package-nội-bộ)
8. [The `main` Package](#8-the-main-package--package-main)

**PART 2 — MODULES**
9. [What is a Module](#9-what-is-a-module--module-là-gì)
10. [go.mod Anatomy](#10-gomod-anatomy--cấu-trúc-gomod)
11. [go.sum File](#11-gosum-file--file-gosum)
12. [Semantic Versioning](#12-semantic-versioning--quản-lý-phiên-bản)
13. [Module Commands](#13-module-commands--các-lệnh-module)
14. [Module Proxy (GOPROXY)](#14-module-proxy--goproxy)
15. [replace Directive](#15-replace-directive)
16. [exclude & retract Directives](#16-exclude--retract-directives)
17. [Workspace Mode (go.work)](#17-workspace-mode--gowork)

---

# PART 1 — PACKAGES

---

## 1. Package Declaration | Khai báo Package

| English | Tiếng Việt |
|---------|-----------|
| Every `.go` source file **must** begin with a package declaration. | Mỗi file `.go` **bắt buộc** phải bắt đầu bằng khai báo package. |
| All files in the same directory must share the **same package name**. | Tất cả file trong cùng một thư mục phải dùng **cùng tên package**. |
| The package name does not have to match the directory name (but it is conventional). | Tên package không cần khớp với tên thư mục (nhưng theo quy ước thì nên khớp). |
| Exception: test files may use `package foo_test` for black-box testing. | Ngoại lệ: file test có thể dùng `package foo_test` để kiểm thử hộp đen (black-box). |

```go
// File: go-rest-api/backend/main.go
package main   // khai báo thuộc package "main"

// File: go-rest-api/backend/handlers.go
package main   // cùng thư mục → cùng package name
```

```go
// File: mylib/greet.go
package greet  // tên package = "greet", thư mục = "mylib"
               // hợp lệ nhưng không theo quy ước tốt
```

> **Spec ref:** [go.dev/ref/spec#Package_clause](https://go.dev/ref/spec#Package_clause)

---

## 2. Naming Conventions | Quy tắc đặt tên

| English | Tiếng Việt |
|---------|-----------|
| Package names should be **short, lowercase, single words**. | Tên package nên **ngắn, viết thường, một từ**. |
| Avoid underscores, hyphens, or mixed caps. | Tránh dùng dấu gạch dưới, gạch ngang hoặc chữ hoa lẫn lộn. |
| The name is used as a prefix when calling exported identifiers: `fmt.Println`. | Tên package dùng làm tiền tố khi gọi định danh được xuất: `fmt.Println`. |
| Avoid stuttering: don't name a method `http.HTTPServer` — prefer `http.Server`. | Tránh lặp từ: không đặt `http.HTTPServer` — hãy dùng `http.Server`. |
| Package names in the standard library are good references: `fmt`, `net`, `io`, `sync`. | Các package trong thư viện chuẩn là ví dụ tốt: `fmt`, `net`, `io`, `sync`. |

```go
// Tốt (good)
package user
package payment
package httputil

// Không tốt (bad)
package UserService   // chữ hoa
package user_service  // dấu gạch dưới
package u             // quá ngắn, không rõ nghĩa
```

> **Blog ref:** [go.dev/blog/package-names](https://go.dev/blog/package-names)

---

## 3. Exported vs Unexported | Xuất ra ngoài vs Nội bộ

| English | Tiếng Việt |
|---------|-----------|
| An identifier starting with an **uppercase letter** is **exported** (visible outside the package). | Định danh bắt đầu bằng **chữ HOA** là **exported** (có thể truy cập từ ngoài package). |
| An identifier starting with a **lowercase letter** is **unexported** (package-private). | Định danh bắt đầu bằng **chữ thường** là **unexported** (chỉ dùng trong package). |
| This applies to: functions, types, variables, constants, struct fields, and methods. | Quy tắc này áp dụng cho: hàm, kiểu, biến, hằng số, trường struct và method. |
| There is no `public`/`private` keyword — capitalization IS the visibility rule. | Không có từ khóa `public`/`private` — viết hoa/thường chính là quy tắc visibility. |

```go
package user

// Exported — truy cập được từ ngoài: user.User, user.GetByID
type User struct {
    ID    string // exported field
    Email string // exported field
    password string // unexported — chỉ dùng trong package "user"
}

func GetByID(id string) *User { ... }  // exported
func hashPassword(p string) string { ... } // unexported
```

```go
// Package khác gọi:
import "myapp/user"

u := user.GetByID("123")    // OK
_ = u.ID                    // OK
_ = u.password              // COMPILE ERROR: unexported field
user.hashPassword("x")      // COMPILE ERROR: unexported function
```

> **Spec ref:** [go.dev/ref/spec#Exported_identifiers](https://go.dev/ref/spec#Exported_identifiers)

---

## 4. Import Paths & Aliasing | Đường dẫn Import

| English | Tiếng Việt |
|---------|-----------|
| The import path is the **module path + subdirectory** of the package. | Đường dẫn import = **module path + thư mục con** chứa package. |
| Use the package's **declared name** (not the directory name) when calling it. | Khi gọi, dùng **tên package đã khai báo** (không phải tên thư mục). |
| You can alias an import to avoid name collisions or for convenience. | Bạn có thể đặt bí danh (alias) cho import để tránh xung đột tên. |
| The `_` alias runs `init()` only — no exported names are bound. | Alias `_` chỉ chạy `init()` — không bind tên nào để sử dụng. |
| The `.` alias imports names directly into the current file's scope (discouraged). | Alias `.` nhập tên trực tiếp vào scope hiện tại (không khuyến khích). |

```go
import (
    "fmt"                          // standard library — ngắn gọn
    "net/http"                     // path = net/http, tên package = http

    "github.com/gin-gonic/gin"     // third-party — từ go-rest-api của bạn
    // ↑ module path: github.com/gin-gonic/gin
    // ↑ package name: gin → dùng như: gin.Default()

    jsoniter "encoding/json"       // alias: jsoniter.Marshal(...)
    _ "myapp/migrations"           // chỉ chạy init() của package này
)
```

```go
// Từ project của bạn: go-rest-api/backend/main.go
import (
    "log"
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"     // import đường dẫn đầy đủ
)

router := gin.Default()           // dùng tên package "gin" làm prefix
```

**Import Factoring — nhóm import theo convention:**

```go
import (
    // 1. Standard library
    "fmt"
    "net/http"

    // 2. Third-party
    "github.com/gin-gonic/gin"

    // 3. Internal (cùng module)
    "myapp/internal/user"
)
```

> **Spec ref:** [go.dev/ref/spec#ImportDecl](https://go.dev/ref/spec#ImportDecl)

---

## 5. Blank Identifier `_` | Định danh trống

| English | Tiếng Việt |
|---------|-----------|
| `_` discards a value — the compiler won't complain about unused results. | `_` bỏ qua một giá trị — compiler không báo lỗi về kết quả không dùng. |
| Used in imports to trigger `init()` side-effects without using any exported name. | Dùng trong import để kích hoạt side-effect `init()` mà không dùng tên nào. |
| Used in multi-return to ignore specific return values. | Dùng trong hàm trả nhiều giá trị để bỏ qua giá trị không cần. |
| Used in for-range to ignore index or value. | Dùng trong for-range để bỏ qua index hoặc value. |

```go
// Bỏ qua giá trị trả về
val, _ := strconv.Atoi("42")  // bỏ qua error

// Bỏ qua index trong range
for _, v := range []int{1, 2, 3} {
    fmt.Println(v)
}

// Import side-effect (ví dụ: đăng ký SQL driver)
import _ "github.com/lib/pq"  // đăng ký PostgreSQL driver, không dùng tên nào

// Type assertion — kiểm tra interface được implement
var _ io.Writer = (*MyWriter)(nil) // compile-time check
```

---

## 6. The `init()` Function | Hàm `init`

| English | Tiếng Việt |
|---------|-----------|
| `init()` runs automatically **after** all variable initializations in a package. | `init()` chạy tự động **sau** khi tất cả biến trong package được khởi tạo. |
| A package can have **multiple** `init()` functions, even in the same file. | Một package có thể có **nhiều** hàm `init()`, kể cả trong cùng một file. |
| `init()` cannot be called explicitly — it is called only by the Go runtime. | `init()` không thể gọi trực tiếp — chỉ Go runtime mới gọi nó. |
| Import order determines `init()` execution order across packages. | Thứ tự import quyết định thứ tự chạy `init()` giữa các package. |
| Use `init()` for one-time setup: registering drivers, validating config. | Dùng `init()` cho các khởi tạo một lần: đăng ký driver, kiểm tra config. |

```go
package config

var AppName string

func init() {
    AppName = "MyApp"
    fmt.Println("config.init() ran")
}

func init() {
    // Hợp lệ: package này có 2 init()
    // Chạy theo thứ tự xuất hiện trong file
    validateConfig()
}
```

**Thứ tự chạy trong một package:**

```
Biến package-level → init() → main() (nếu là package main)
```

**Thứ tự chạy giữa các package:**

```
A import B, B import C
→ C init() → B init() → A init()
```

> **Spec ref:** [go.dev/ref/spec#Package_initialization](https://go.dev/ref/spec#Package_initialization)

---

## 7. Internal Packages | Package Nội bộ

| English | Tiếng Việt |
|---------|-----------|
| A package path containing `internal` can only be imported by code rooted at the **parent** of `internal`. | Package có `internal` trong đường dẫn chỉ được import bởi code nằm trong **thư mục cha** của `internal`. |
| This is enforced by the Go compiler — no configuration needed. | Điều này được compiler Go thực thi — không cần cấu hình. |
| Use `internal` to expose implementation details to sibling packages but not external consumers. | Dùng `internal` để chia sẻ chi tiết cài đặt giữa các package nội bộ nhưng ẩn với bên ngoài. |

```
myapp/
├── go.mod          (module myapp)
├── main.go
├── api/
│   └── handler.go  ← CÓ THỂ import "myapp/internal/db"
├── internal/
│   └── db/
│       └── db.go   ← package "db"
└── external/
    └── plugin.go   ← KHÔNG THỂ import "myapp/internal/db" (compiler error)
```

```go
// api/handler.go — OK vì cùng gốc "myapp/"
import "myapp/internal/db"

// external/plugin.go — COMPILE ERROR
import "myapp/internal/db"
// use of internal package myapp/internal/db not allowed
```

> **Spec ref:** [go.dev/cmd/go/#hdr-Internal_Directories](https://pkg.go.dev/cmd/go#hdr-Internal_Directories)

---

## 8. The `main` Package | Package `main`

| English | Tiếng Việt |
|---------|-----------|
| `package main` is special — it defines an **executable program**, not a library. | `package main` đặc biệt — nó định nghĩa **chương trình thực thi**, không phải thư viện. |
| An executable `package main` must have exactly one `func main()`. | `package main` dạng thực thi phải có đúng một hàm `func main()`. |
| `main()` is the entry point — execution starts here after all `init()` calls. | `main()` là điểm vào — thực thi bắt đầu từ đây sau khi tất cả `init()` đã chạy. |
| Multiple files in the same directory can all be `package main`. | Nhiều file trong cùng thư mục đều có thể là `package main`. |

```go
// go-rest-api/backend/main.go — từ project của bạn
package main

import (
    "log"
    "github.com/gin-gonic/gin"
)

func main() {                  // ← entry point
    router := gin.Default()
    // ...
    router.Run(":8080")
}
```

```
build diagram:

package main (main.go + handlers.go + models.go)
        ↓
    go build ./...
        ↓
  ./backend  (binary)
```

---

# PART 2 — MODULES

---

## 9. What is a Module | Module là gì

| English | Tiếng Việt |
|---------|-----------|
| A **module** is a collection of related Go packages versioned together. | Một **module** là tập hợp các package Go liên quan, được quản lý phiên bản cùng nhau. |
| A module is defined by a `go.mod` file at its root. | Module được định nghĩa bởi file `go.mod` ở thư mục gốc. |
| The module path is the canonical name for the module (used as import prefix). | Module path là tên chính thức của module (dùng làm tiền tố cho import). |
| Modules replaced `GOPATH` as the standard dependency management system (since Go 1.11, default since Go 1.16). | Module thay thế `GOPATH` làm hệ thống quản lý phụ thuộc tiêu chuẩn (từ Go 1.11, mặc định từ Go 1.16). |
| One repository typically contains one module, but can contain multiple. | Một repo thường chứa một module, nhưng có thể chứa nhiều module. |

```
Cấu trúc điển hình / Typical structure:

go-rest-api/backend/     ← thư mục gốc module
├── go.mod               ← định nghĩa module "go-rest-api/backend"
├── go.sum               ← checksum tất cả dependencies
├── main.go              ← package main
├── handlers.go          ← package main
└── models.go            ← package main
```

---

## 10. `go.mod` Anatomy | Cấu trúc `go.mod`

| English | Tiếng Việt |
|---------|-----------|
| `module` directive declares the module's path. | Directive `module` khai báo đường dẫn của module. |
| `go` directive declares the minimum Go language version. | Directive `go` khai báo phiên bản Go tối thiểu cần dùng. |
| `require` lists direct and indirect dependencies with their versions. | `require` liệt kê các phụ thuộc trực tiếp và gián tiếp kèm phiên bản. |
| `// indirect` comment marks a dependency not directly imported in this module's code. | Comment `// indirect` đánh dấu phụ thuộc không được import trực tiếp trong code của module này. |

```go
// go-rest-api/backend/go.mod — từ project của bạn

module go-rest-api/backend        // ← module path (tiền tố import)

go 1.23.0                         // ← phiên bản Go tối thiểu

require github.com/gin-gonic/gin v1.11.0  // ← direct dependency

require (
    github.com/bytedance/sonic v1.14.0 // indirect  ← gin cần, bạn không import trực tiếp
    github.com/go-playground/validator/v10 v10.27.0 // indirect
    // ...
)
```

**Giải thích module path:**

```
module go-rest-api/backend

→ Tất cả package trong module này có import path bắt đầu bằng "go-rest-api/backend"
→ Ví dụ: package ở thư mục "handlers/" sẽ có path "go-rest-api/backend/handlers"
→ Khi publish lên GitHub: thường dùng "github.com/username/reponame"
```

**Tất cả directives trong `go.mod`:**

| Directive | Mô tả |
|-----------|-------|
| `module` | Tên module (bắt buộc) |
| `go` | Phiên bản Go tối thiểu |
| `require` | Danh sách phụ thuộc |
| `replace` | Thay thế module bằng path khác |
| `exclude` | Loại trừ một phiên bản cụ thể |
| `retract` | Đánh dấu phiên bản bị thu hồi (dành cho tác giả module) |

> **Ref:** [go.dev/ref/mod#go-mod-file](https://go.dev/ref/mod#go-mod-file)

---

## 11. `go.sum` File | File `go.sum`

| English | Tiếng Việt |
|---------|-----------|
| `go.sum` stores **cryptographic checksums** of all module dependencies. | `go.sum` lưu **checksum mã hóa** của tất cả phụ thuộc module. |
| It ensures the same dependency files are downloaded on every build. | Đảm bảo cùng file phụ thuộc được tải xuống mỗi lần build. |
| `go.sum` should be **committed to version control**. | `go.sum` **phải được commit** vào version control. |
| Never edit `go.sum` manually — let `go mod tidy` and `go get` manage it. | Không bao giờ tự sửa `go.sum` — để `go mod tidy` và `go get` quản lý. |
| Each line: `<module> <version>[/go.mod] <hash>` | Mỗi dòng: `<module> <phiên bản>[/go.mod] <hash>` |

```
# go.sum — ví dụ một số dòng
github.com/gin-gonic/gin v1.11.0 h1:XXXXXXXX...=
github.com/gin-gonic/gin v1.11.0/go.mod h1:XXXXXXXX...=

# Dòng không có "/go.mod" → checksum của source code
# Dòng có "/go.mod" → checksum của file go.mod
```

---

## 12. Semantic Versioning | Quản lý Phiên bản

| English | Tiếng Việt |
|---------|-----------|
| Go modules use **Semantic Versioning** (semver): `vMAJOR.MINOR.PATCH`. | Go modules dùng **Semantic Versioning** (semver): `vMAJOR.MINOR.PATCH`. |
| `PATCH` bump: backward-compatible bug fix. | `PATCH` tăng: sửa lỗi, tương thích ngược. |
| `MINOR` bump: backward-compatible new features. | `MINOR` tăng: tính năng mới, tương thích ngược. |
| `MAJOR` bump: breaking API change. | `MAJOR` tăng: thay đổi API, **không** tương thích ngược. |
| **Major version v2+ requires a different import path** (the "import compatibility rule"). | **Major version v2+ yêu cầu đường dẫn import khác** (quy tắc tương thích import). |

```
v1.2.3
│ │ └─ PATCH: sửa lỗi
│ └─── MINOR: tính năng mới
└───── MAJOR: breaking change

# Import path thay đổi khi MAJOR >= 2:
import "github.com/foo/bar"     # v1.x.x
import "github.com/foo/bar/v2"  # v2.x.x  ← thêm /v2
import "github.com/foo/bar/v3"  # v3.x.x  ← thêm /v3

# Từ go.mod của bạn:
github.com/go-playground/validator/v10 v10.27.0
#                                 ↑ /v10 trong import path
```

**Pre-release & pseudo-versions:**

```
v1.2.3-alpha.1        # pre-release
v1.2.3-beta.2

# Pseudo-version (khi dùng commit hash, không có tag):
v0.0.0-20231015123456-abcdef123456
#        ↑ timestamp   ↑ commit hash
```

**Minimum Version Selection (MVS):**

| English | Tiếng Việt |
|---------|-----------|
| Go uses MVS: it picks the **minimum** version that satisfies all requirements. | Go dùng MVS: chọn phiên bản **tối thiểu** thỏa mãn tất cả yêu cầu. |
| This makes builds reproducible without a lock file (go.sum plays that role). | Điều này làm build có thể tái tạo mà không cần lock file riêng. |

> **Ref:** [go.dev/ref/mod#version-queries](https://go.dev/ref/mod#version-queries)

---

## 13. Module Commands | Các lệnh Module

| Command | What it does / Tác dụng |
|---------|------------------------|
| `go mod init <path>` | Tạo file `go.mod` với module path cho trước |
| `go mod tidy` | Thêm dependencies còn thiếu, xóa dependencies không dùng, cập nhật `go.sum` |
| `go get <pkg>@<version>` | Thêm hoặc cập nhật dependency; `@latest` cho phiên bản mới nhất |
| `go get <pkg>@none` | Xóa dependency khỏi `go.mod` |
| `go mod download` | Tải tất cả dependencies về cache (không build) |
| `go mod verify` | Kiểm tra checksum trong `go.sum` khớp với cache |
| `go mod graph` | In dependency graph dạng text |
| `go mod why <pkg>` | Giải thích tại sao module cần `<pkg>` |
| `go mod vendor` | Copy dependencies vào thư mục `vendor/` |
| `go list -m all` | Liệt kê tất cả module dependencies |
| `go list -m -versions <mod>` | Liệt kê tất cả phiên bản available của module |

```bash
# Khởi tạo module mới
go mod init github.com/yourname/myapp

# Sau khi thêm import trong code, đồng bộ go.mod
go mod tidy

# Thêm gin phiên bản mới nhất (ví dụ trong project của bạn)
go get github.com/gin-gonic/gin@latest

# Thêm gin phiên bản cụ thể
go get github.com/gin-gonic/gin@v1.11.0

# Xóa một dependency
go get github.com/some/pkg@none

# Xem module nào cần gin
go mod why github.com/gin-gonic/gin

# Kiểm tra tính toàn vẹn
go mod verify
```

> **Ref:** [go.dev/ref/mod#go-mod-tidy](https://go.dev/ref/mod#go-mod-tidy)

---

## 14. Module Proxy | GOPROXY

| English | Tiếng Việt |
|---------|-----------|
| `GOPROXY` is an env var pointing to a module proxy server. | `GOPROXY` là biến môi trường trỏ đến module proxy server. |
| Default: `GOPROXY=https://proxy.golang.org,direct` — tries proxy first, then direct. | Mặc định: `GOPROXY=https://proxy.golang.org,direct` — thử proxy trước, rồi mới direct. |
| The proxy caches modules, improving download speed and availability. | Proxy lưu cache module, cải thiện tốc độ tải và tính khả dụng. |
| `GONOSUMCHECK` / `GONOSUMDB`: skip checksum verification for private modules. | `GONOSUMCHECK` / `GONOSUMDB`: bỏ qua kiểm tra checksum cho module private. |
| `GOPRIVATE`: comma-separated list of module path prefixes that bypass proxy and sum db. | `GOPRIVATE`: danh sách tiền tố module path bỏ qua proxy và sum db. |

```bash
# Biến môi trường quan trọng
GOPROXY=https://proxy.golang.org,direct   # default
GOPROXY=off                               # không cho tải module (offline)
GOPROXY=direct                            # tải trực tiếp từ VCS

# Cho module private (ví dụ: nội bộ công ty)
GOPRIVATE=github.com/yourcompany/*
GONOSUMDB=github.com/yourcompany/*

# Proxy Việt Nam / nhanh hơn cho khu vực châu Á:
GOPROXY=https://goproxy.io,direct
GOPROXY=https://goproxy.cn,direct  # mirror của Trung Quốc
```

> **Ref:** [go.dev/ref/mod#module-proxy](https://go.dev/ref/mod#module-proxy)

---

## 15. `replace` Directive

| English | Tiếng Việt |
|---------|-----------|
| `replace` substitutes a module with a different version or local path. | `replace` thay thế một module bằng phiên bản khác hoặc đường dẫn local. |
| Common use cases: local development, forking a dependency, testing a patch. | Trường hợp dùng phổ biến: phát triển local, fork dependency, test bản vá. |
| A `replace` in a dependency's `go.mod` is **ignored** — only the top-level module's `replace` takes effect. | `replace` trong `go.mod` của dependency **bị bỏ qua** — chỉ `replace` của module cấp cao nhất có tác dụng. |

```
# Cú pháp / Syntax:
replace <module> [version] => <newmodule> [newversion]
replace <module> [version] => <localpath>
```

```go
// go.mod — ví dụ thực tế

// Dùng local fork thay vì module từ internet
replace github.com/gin-gonic/gin => ../my-gin-fork

// Dùng local path tuyệt đối
replace github.com/gin-gonic/gin => /home/user/projects/gin

// Thay thế bằng phiên bản khác
replace github.com/old/module v1.0.0 => github.com/new/module v2.0.0

// Chỉ replace phiên bản cụ thể
replace github.com/foo/bar v1.2.0 => github.com/foo/bar v1.2.1
```

> `replace` với local path rất hữu ích khi bạn đang phát triển 2 module cùng lúc trong workspace.

---

## 16. `exclude` & `retract` Directives

### `exclude`

| English | Tiếng Việt |
|---------|-----------|
| `exclude` prevents a specific module version from being used. | `exclude` ngăn một phiên bản module cụ thể được sử dụng. |
| Useful when a specific version has a known bug or security vulnerability. | Hữu ích khi một phiên bản cụ thể có lỗi hoặc lỗ hổng bảo mật. |
| The next highest compatible version is used instead. | Phiên bản tương thích cao nhất tiếp theo sẽ được dùng thay thế. |

```go
// go.mod
exclude github.com/some/module v1.2.3  // bỏ qua phiên bản này

exclude (
    github.com/some/module v1.2.3
    github.com/another/mod v0.5.0
)
```

### `retract`

| English | Tiếng Việt |
|---------|-----------|
| `retract` is used by module **authors** to mark versions that should not be used. | `retract` dành cho **tác giả** module để đánh dấu phiên bản không nên dùng. |
| Different from `exclude`: `retract` is in the module's own `go.mod`, `exclude` is in the consumer's `go.mod`. | Khác `exclude`: `retract` nằm trong `go.mod` của chính module đó, `exclude` nằm trong `go.mod` của người dùng. |
| Retracted versions still work if already in use — Go just warns the user. | Phiên bản bị retract vẫn hoạt động nếu đang dùng — Go chỉ cảnh báo người dùng. |

```go
// go.mod của tác giả module
retract v1.0.5 // Contains critical security bug

retract (
    v1.0.5    // security vulnerability
    v1.0.6    // accidentally published
)
```

---

## 17. Workspace Mode | `go.work`

| English | Tiếng Việt |
|---------|-----------|
| Workspace mode (Go 1.18+) allows working on **multiple modules simultaneously** without `replace` directives. | Workspace mode (Go 1.18+) cho phép làm việc với **nhiều module cùng lúc** mà không cần `replace`. |
| A `go.work` file at the workspace root defines which modules are included. | File `go.work` ở thư mục gốc workspace định nghĩa các module được bao gồm. |
| `go.work` should typically **not** be committed to version control (add to `.gitignore`). | `go.work` thường **không nên** commit vào version control (thêm vào `.gitignore`). |
| `go.work.sum` is the sum database file for the workspace. | `go.work.sum` là file checksum cho workspace. |

```bash
# Khởi tạo workspace
go work init ./module-a ./module-b

# Thêm module vào workspace đang có
go work use ./module-c

# Sync go.work.sum
go work sync
```

```
# go.work — ví dụ cấu trúc
go 1.23.0

use (
    ./go-rest-api/backend     # module "go-rest-api/backend"
    ./go-file/hello-go        # module "hello-go"
)
```

**So sánh `replace` vs `go.work`:**

| | `replace` | `go.work` |
|---|---|---|
| Phạm vi | Một module | Toàn workspace |
| Commit? | Thường là không | Không |
| Ảnh hưởng consumer? | Có (nếu commit go.mod) | Không |
| Use case | Production hotfix | Local development |

**Ví dụ thực tế với project của bạn:**

```
dev/go-learn/              ← workspace root
├── go.work
├── go-file/hello-go/      ← module "hello-go"
├── go-rest-api/backend/   ← module "go-rest-api/backend"
└── go-package/            ← (module mới)

# go.work
go 1.23.0
use (
    ./go-file/hello-go
    ./go-rest-api/backend
    ./go-package
)
```

> **Ref:** [go.dev/ref/mod#workspaces](https://go.dev/ref/mod#workspaces)

---

## Quick Reference | Tóm tắt nhanh

### Packages

| Concept | Rule |
|---------|------|
| Package declaration | Dòng đầu tiên của mọi `.go` file |
| Exported identifier | Bắt đầu bằng **chữ HOA** |
| Unexported identifier | Bắt đầu bằng **chữ thường** |
| `init()` | Chạy tự động, không gọi trực tiếp được |
| `internal/` | Chỉ import được từ thư mục cha |
| `package main` | Entry point của binary, cần `func main()` |

### Modules

| Command | Khi nào dùng |
|---------|-------------|
| `go mod init` | Tạo module mới |
| `go mod tidy` | Sau khi thêm/xóa imports |
| `go get pkg@version` | Thêm hoặc cập nhật dependency |
| `go mod verify` | Kiểm tra tính toàn vẹn của cache |
| `go work init` | Bắt đầu multi-module workspace |

### Key Files

| File | Mục đích |
|------|---------|
| `go.mod` | Định nghĩa module, liệt kê dependencies |
| `go.sum` | Checksum — **luôn commit** |
| `go.work` | Workspace config — **không commit** |

---

*Generated: 2026-05-16 | Source: [go.dev/ref/mod](https://go.dev/ref/mod), [go.dev/ref/spec](https://go.dev/ref/spec)*

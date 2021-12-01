### Go Module

Để giúp quản lý các package dependency trong Go, thì từ phiên bản Go `v1.11` được bổ sung thêm tính năng Go Module

Go Module là tập hợp các package. Mỗi module đều có 1 version, vì vậy các package trong module cũng có thuộc tính version. Sau đó module và tất cả package của nó sẽ tạo thành một khối độc lập, được đánh version, phát hành và phân phối cùng nhau

Thông thường một repository sẽ tương ứng với một Go Module. File `go.mod` sẽ được đặt ở top-level directory của Go Module và mỗi file `go.mod` sẽ định nghĩa một module duy nhất, có nghĩa là Go module và go.mod có quan hệ 1-1 với nhau.

Top-level directory nơi chứa file `go.mod` được gọi là root directory của module. Tất cả các Go package trong thư mục gốc và các thư mục con của nó thuộc về Go module này, và Go module này được được gọi là main module hay module chính

### Tạo Go Module

Thông thường sẽ có 3 bước để tạo Go Module

1. Câu lệnh `go mod init` để tạo file go.mod
2. `go mod tidy` để download/update các package dependencies
3. `go build` để build ứng dụng

```
$ mkdir -p goprojects/samplemodule
$ cd goprojects/samplemodule
$ touch main.go
```

Code ví dụ cho file main.go như sau

```go
package main
import "github.com/sirupsen/logrus"
func main() {
    logrus.Println("hello, go module mode")
}
```

#### Chạy “go init”

Bây giờ chạy câu lệnh `go mod init` để khởi tạo go module

```
$ go mod init github.com/tonylixu/samplemodule
go: creating new go.mod: module github.com/tonylixu/samplemodule
go: to add module requirements and sums:
 go mod tidy
```

Chúng ta có thể thấy file `go.mod` đã được tạo

```
$ cat go.mod
module github.com/tonylixu/samplemodule
go 1.17
```

File `go.mod` này giúp `samplemodule` project thành một Go Module. Root directory bây giờ là `samplemodule`. Dòng đầu tiên của file `go.mod` khai báo module path, còn dòng cuối cùng chỉ định thông tin version go được sử dụng

#### Chạy "go mod tidy"

Dựa trên kết quả đầu ra của `go mod init`, `go mod tidy` giúp download và update những package dependency được sử dụng trong project

```
$ go mod tidy
go: finding module for package github.com/sirupsen/logrus
go: downloading github.com/sirupsen/logrus v1.8.1
go: found github.com/sirupsen/logrus in github.com/sirupsen/logrus v1.8.1
go: downloading golang.org/x/sys v0.0.0-20191026070338-33540a1f6037
go: downloading github.com/stretchr/testify v1.2.2
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
```

`go mod tidy` phân tích tất cả các file trong main module hiện tại, tìm ra tất cả các third-party dependencies, xác định phiên bản của third-party dependencies và tải xuống các dependencies trực tiếp và gián tiếp của main module

Một điều cần lưu ý là `go mod tidy` sẽ đặt tất cả các package trong local module cache path, giá trị mặc định là `$GOPATH[0]/pkg/mod`. Từ go 1.15 và các phiên bản sau này có thể tùy chỉnh cache path của local module thông qua biến môi trường `GOMODCACHE`

Nội dụng của file `go.mod` sẽ trông như sau

```
$ cat go.mod
module github.com/tonylixu/samplemodule
go 1.17
require github.com/sirupsen/logrus v1.8.1
require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
```

Chúng ta có thể thấy là một file có tên là `go.sum` cũng được tạo ra

```
$ cat go.sum
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.8.1 h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=
github.com/sirupsen/logrus v1.8.1/go.mod h1:yWOB1SBYBC5VeMP7gHvWumXLIWorT60ONWic61uBYv0=
github.com/stretchr/testify v1.2.2 h1:bSDNvY7ZPG5RlJ8otE/7V6gMiyenm9RtJ7IUVIAoJ1w=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 h1:YyJpGZS1sBuBCzLAR1VEpK193GlqGZbnPFnPV/5Rsb4=
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
```

File `go.sum` lưu trữ hash value của mỗi dependency. Đây là một biện pháp bảo mật của Go Module. Khi một version cụ thể của một module ở đây được tải lại trong tương lai, go sẽ sử dụng hash value tương ứng trong file `go.sum` để so sánh với hash value của nội dung mới. Chỉ khi giá trị hash value bằng nhau thì packages mới được coi là hợp pháp

Bằng cách này, có thể đảm bảo rằng nội dụng của module mà project của chúng ta phụ thuộc vào sẽ không bị giả mạo do vô tình hay cố ý. Chúng ta nên tạo `go.mod` và `go.sum` với source code của chúng ta.

#### Chạy "go build"

Câu lệnh `go build` sẽ đọc thông tin phiên bản và dependency trong file `go.mod`, tìm phiên bản tương ứng của dependency module trong local module cache path đồng thời biên dịch và liên kết chúng lại với nhau

```
$ go build
$ ls -lh
total 4048
-rw-r--r--  1 txu  staff   165B Nov 13 11:45 go.mod
-rw-r--r--  1 txu  staff   899B Nov 13 11:45 go.sum
-rw-r--r--  1 txu  staff   111B Nov 13 11:31 main.go
-rwxr-xr-x  1 txu  staff   2.0M Nov 13 11:58 samplemodule
```

### Nguyên tắc Go Module làm việc

Đối với một dependency package thường có nhiều phiên bản, làm cách nào để Go Module xác định phiên bản nào là phù hợp nhất? Để tìm hiểu điều này, chúng ta cần biết ít nhất 2 cơ chế sau của Go Module

- Semantic Import Versioning
- Minimal Version Selection

#### Semantic Import Versioning

Semantic versioning bao gồm 3 phần : major (chính), minor (phụ) và patch (bản vá).

![Versioning](https://techmaster.vn/media/static/9479/c6b93n451co50fuc8dag)

Với semantic version specification, lệnh go có thể xác định thứ tự phát hành 2 phiên bản của cùng một module và liệu chúng có tương thích hay không. Theo semantic version specification thì 2 phiên bản có số phiên bản chính (major version) khác nhau không tương thích với nhau

Hơn nữa, khi phiên bản chính giống nhau, phiên bản phụ hầu hết tương thích ngược với phiên bản phụ có version nhỏ hơn. Phiên bản vá lỗi cũng không ảnh hưởng đến khả năng tương thích. Hơn nữa, Go Module quy định rằng nếu phiên bản mới và cũ của cùng một package, thì đường dẫn để import package phải giống nhau

Giả sử `goapp` có 2 phiên bản là `v1.1.0` và `v1.2.0` chúng có cùng version chính, vì vậy nếu các dự án của chúng ta phụ thuộc vào `goapp` thì không thành vấn đề với `v1.1.0` và `v1.2.0`, nó có thể được cập nhật như sau

```go
import "github.com/tonylixu/goapp"
```

Tuy nhiên, nếu có 2 phiên bản là `v1.1.0` và `v2.2.0` lúc này phiên bản chính là khác nhau nên chúng ta không để sử dụng cách trên để import. Thay vào đó chúng ta sử dụng phiên bản chính trong quá trình import

```go
import (
    "github.com/tonylixu/goapp"
    goappv2 "github.com/tonylixu/goapp/v2"
)
```

#### Minimal Version Selection

Theo dõi hình dưới đây, project có 2 dependencies là : `A v1.1.0` và `B v1.2.0`. Nhưng package A phụ thuộc vào package `C v1.3.0`, package B cũng phụ thuộc vào package C, nhưng là phiên bản `C v1.4.0`. Phiên bản mới nhất của package C là `C v1.7.0`. Vậy Go sẽ chọn phiên bản nào của package C để download

![Minimal Version Selection](https://techmaster.vn/media/static/9479/c6b958c51co50fuc8db0)

Hầu hết các ngôn ngữ lập trình chính thống hiện nay sẽ chọn `C v1.7.0`. Tuy nhiên, trong Go thì khác, nó không chỉ xem xét độ ổn định và bảo mật mới nhất mà còn phụ thuộc vào yêu cầu của từng module: Package A nêu rõ rằng chỉ cẩn `C v1.3.0` và package B nói rõ rằng chỉ cần `C v1.4.0`. Vì vậy, Go sẽ chọn "Phiên bản tối thiểu" đáp ứng các yêu cầu chung của dự án trong số tất cả các phiên bản dependencies của dự án
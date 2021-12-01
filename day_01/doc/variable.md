## Biến trong Golang

Biến là một định danh dùng để lưu trữ dữ liệu, thông qua biến chúng ta có thể thực hiện các thao tác với dữ liệu.

Mỗi biến có một kiểu dữ liệu riêng, dựa vào kiểu dữ liệu của biến có các thao tác khác nhau với biến.

Ví dụ:

- Biến kiểu số thì có phép cộng, trừ, nhân, chia.
- Biến kiểu chuỗi thì có độ dài của chuỗi, phép nối 2 chuỗi, ...

### Quy tắc khai báo biến

Quy tắc đặt tên biến trong golang

- Biến phải được bắt đầu bằng chữ, không được là số
- Biến không chứa dấu "`space`"
- Nếu tên biến được bắt đầu bằng một chữ cái viết thường, nó chỉ có thể được truy cập trong package hiện tại, đây được coi là các biến `unexported`.
- Nếu tên biến được bắt đầu bằng một chữ cái in hoa, nó có thể được truy cập từ các package khác, đây được coi là các biến `exported`
- Có phân biệt chữ hoa, chữ thường
- Nên viết biến theo kiểu là `camelCase`

### Zero values

Khi một biến được khai báo mà chưa được khởi tạo giá trị. Các biến này sẽ nhận giá trị Zero values tùy thuộc vào kiểu dữ liệu của biến

- string => ""
- int => 0
- bool => false
- float => 0.0
- ...

*/

### Khai báo biến một biến
```go
var number int
```

Khi khai báo biến mà không khởi tạo giá trị thì mặc định biến đó nhận giá trị **zero value** ứng với kiểu dữ liệu của biến đó

Ví dụ

```go
var number int
fmt.println(number) // 0
```

##### Khai báo biến sau đó mới khởi tạo giá trị cho biến

```go
var number int
number = 32
```

### Khai báo biến đồng thời khởi tạo giá trị cho biến

```go
var number int = 32
```

### Khai báo biến không cần chỉ định kiểu dữ liệu

```go
var number = 32
```
> Go có thể tự động suy ra kiểu của biến này từ giá trị khởi tạo đó

### Khai báo biến kiểu shorthand

```go
var number := 32
```
> Cách này không áp dụng với biến khai báo ngoài function

> Cách khai báo nhanh được sử dụng khi có ít nhất một biến mới được khai báo phía bên trái của toán tử :=

### Khai báo biến nhiều biến

```go
// Trường hợp 2 biến khác kiểu dữ liệu
var (
    number int
    name string
)
```

```go
// Trường hợp 2 biến cùng kiểu dữ liệu
var number, age int
```

```go
// Khai báo xong mới khởi tạo giá trị cho biến
var (
    number int
    name string
)
number = 23
name = "Hiên"
```

```go
// Trường hợp 2 biến cùng kiểu dữ liệu
var number, age int
number = 23
age = 40
```

```go
// Khai báo biến kiểu suy luận
var number = 23
var age = 40
```

```go
// Khai báo biến nhanh
number := 23
age := 40
```



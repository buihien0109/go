## Function trong Golang

### 1. Định nghĩa
Function là một tập hợp các đoạn code thực hiện một công việc nhất định. 

Function tiếp nhận các tham số đầu vào, sau đó thực hiện một vài tính toán và trả về kết quả đầu ra.

### 2. Khởi tạo function

#### Cú pháp

```go
func functionName(parameterName type) returnType {  
    //function body
}
```

> **parameterName**, **returnType** là các giá trị tùy chọn của function

#### Các loại function

- Function có tham số đầu vào hoặc không
- Function có trả về kết quả hoặc không

### Các ví dụ về fuction

**Function không có tham số**

```go
func sayHello() {  
    fmt.Println("Xin chào các bạn")
}
```

**Function 1 tham số**

```go
func sayHelloWithName(name string) {
	fmt.Printf("Xin chào %s\n", name)
}
```

**Function 2 tham số**

```go
func sayHelloWithInfo(name string, year int) {
	fmt.Printf("Xin chào %s. Năm nay %d tuổi\n", name, 2021-year)
}
```

**Function trả về 1 kết quả**

```go
func sum2Number(a, b int) int {
	result := a + b
	return result
}
```

**Function trả về nhiều kết quả**

```go
func calculateNumber(a, b int) (int, int) {
	add := a + b
	subtract := a - b

	return add, subtract
}

// Các giá trị trả về được đặt tên
func calculateNumber2(a, b int) (add, subtract int) {
	add = a + b
	subtract = a - b

	return
}
```

Các giá trị trả về được đặt tên

**Variadic Functions**

```go
func calculateNumbers(numbers ...int) (sum int) {
	for _, c := range numbers {
		sum += c
	}
	return sum
}
```

**Định dang trống**

"_" được hiểu là định danh trống trong Go

Dùng để loại bỏ giá trị không được sử dụng được trả về từ function

```go
func calculate(a,b int) (add, plus int){
    add = a + b
    plus = a * b
    return
}

func main() {
	_, plus := calculate(3, 4)
	fmt.Println(plus)
}
```
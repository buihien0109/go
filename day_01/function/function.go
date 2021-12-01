package function

import (
	"fmt"
	"reflect"
)

// Function không có tham số
func sayHello() {
	fmt.Println("Xin chào các bạn")
}

// Function 1 tham số
func sayHelloWithName(name string) {
	fmt.Printf("Xin chào %s\n", name)
}

// Function 2 tham số
func sayHelloWithInfo(name string, year int) {
	fmt.Printf("Xin chào %s. Năm nay %d tuổi\n", name, 2021-year)
}

// Function trả về 1 kết quả
func sum2Number(a, b int) int {
	result := a + b
	return result
}

// Function trả về nhiều kết quả
func calculateNumber(a, b int) (int, int) {
	add := a + b
	subtract := a - b

	return add, subtract
}

// Function trả về nhiều kết quả
func calculateNumber2(a, b int) (add, subtract int) {
	add = a + b
	subtract = a - b

	return
}

// Variadic Functions
func printColors(color ...string) {
	// Kiểm tra kiểu dữ liệu của color
	fmt.Printf("Color type : %T\n", color)
	fmt.Println(reflect.TypeOf(color))
	fmt.Println(reflect.ValueOf(color).Kind())

	for _, c := range color {
		fmt.Println(c)
	}
}

// Variadic Functions
func calculateNumbers(numbers ...int) (sum int) {
	for _, c := range numbers {
		sum += c
	}
	return sum
}

/*
-- First Class Function

- Function có thể gán cho 1 biến
- Truyền function vào làm tham số cho function khác
- Function return function
*/

var square = func(a int) int {
	return a * a
}

var sumNumbers = func(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return result
}

var oddNumbers = func(numbers []int) (result []int) {
	for _, number := range numbers {
		if number%2 != 0 {
			result = append(result, number)
		}
	}

	return result
}

func simple() func(a, b int) int {
    f := func(a, b int) int {
        return a + b
    }
    return f
}

func DemoFunction() {
	// sayHello()
	// sayHelloWithName("Nguyễn Văn A")
	// sayHelloWithInfo("Trần Văn B", 1990)

	// a, b := calculateNumber(10, 5)
	// fmt.Println(a, b)

	// c, _ := calculateNumber(20, 5)
	// fmt.Println(c)

	// a1, b1 := calculateNumber2(6, 2)
	// fmt.Println(a1, b1)

	// result := sum2Number(3, 4)
	// fmt.Println(result)

	// printColors("red", "green", "blue", "yellow")

	// result1 := calculateNumbers(1,2,3,4,5,6,7)
	// fmt.Println(result1)

	result2 := square(3)
	fmt.Println(result2)

	result3 := sumNumbers(oddNumbers([]int{1, 2, 3, 4, 5}))
	fmt.Println("result3: ", result3)

	s := simple()
    fmt.Println(s(60, 7))
}

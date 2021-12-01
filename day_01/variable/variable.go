package variable

import (
	"fmt"
	"reflect"
)

func DemoVariable() {
	// Khai báo biến
	var name string
	fmt.Println(name)

	// Gán dữ liệu cho biến
	name = "Bùi Hiên"
	fmt.Println(name)

	// Khai báo biến và gán giá trị cho biến
	var age int = 24
	fmt.Println(age)

	// Kiểm tra kiểu dữ liệu của biến
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))

}

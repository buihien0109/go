## First class functions là gì?

Một ngôn ngữ hỗ trợ first class functions cho phép function có thể gán cho biến, truyền vào làm tham số cho function khác và có thể được trả về từ functions khác. Go có hỗ trợ cho first class functions

Trong bài viết này, chúng ta sẽ thảo luận về cú pháp và các trường hợp sử dụng khác nhau của first class functions.

### Anonymous functions

Chúng ta bắt đầu với 1 ví dụ đơn giản bằng việc gán function cho 1 biến

```go
package main

import (
    "fmt"
)

func main() {
    a := func() {
        fmt.Println("hello world first class function")
    }
    a()
    fmt.Printf("%T", a)
}
```

Trong ví dụ trên, chúng ta gán một function cho biến `a`. Đây là cú pháp để gán một function cho 1 biến. Nếu bạn để ý kỹ, function được gán cho a không có tên. Các functions kiểu này được gọi là `anonymous functions` bởi vì chúng không có tên

Cách duy nhất để gọi function này là sử dụng biến `a`

Kết quả

```
hello world first class function
func()
```

`a()` để gọi function và in ra " **hello world first class function**" và type của a trong trường hợp này là `fun()`

Chúng ta cũng có thể gọi một anonymous function mà không cần gán nó cho một biến nào cả. Hãy xem cách này được thực hiện như thế nào trong ví dụ sau.

```go
package main

import (
    "fmt"
)

func main() {
    func() {
        fmt.Println("hello world first class function")
    }()
}
```

Trong ví dụ trên, Sau khi chúng ta định nghĩa một anonymous function được định nghĩa, sau đó sử dụng **()** ngay sau function để thực hiện gọi function

Kết quả

```
hello world first class function
```

Chúng ta cũng có thể truyền đối số cho các anonymous function giống như bất kỳ function nào khác.

```go
package main

import (
    "fmt"
)

func main() {
    func(n string) {
        fmt.Println("Welcome", n)
    }("Gophers")
}
```

Trong ví dụ trên, một đối số kiểu **string** được truyền vào hàm ẩn danh ngay sau khi nó được khai báo

Kết quả

```
Welcome Gophers
```

### Kiểu function do người dùng tự định nghĩa

Giống như cách chúng ta định nghĩa các struct, chúng ta có thể định nghĩa các kiểu function

```go
type add func(a int, b int) int
```

Đoạn code trên tạo ra một function mới có type `add`. Func này chấp nhận 2 tham số đầu vào có kiểu int và trả về giá trị cũng là kiểu int. Bây giờ chúng ta có thể định nghĩa các biến type `add`.

```go
package main

import (
    "fmt"
)

type add func(a int, b int) int

func main() {
    var a add = func(a int, b int) int {
        return a + b
    }
    s := a(5, 6)
    fmt.Println("Sum", s)
}
```

Trong chương trình trên, chúng ta định nghĩa một biến type `add` và gán cho nó một function khớp với type `add`. Sau đó gọi `a(5, 6)` và gán giá trị cho biến `s`

Kết quả

```
Sum 11
```

### Higher-order functions

Theo wiki thì Higher-order functions được định nghĩa là function thực hiện ít nhất một trong những điều sau đây:

- Nhận 1 hoặc nhiều function khác vào làm tham số
- Trả về một function khác

Hãy xem các ví dụ cho những trường hợp này

#### Nhận 1 hoặc nhiều function khác vào làm tham số

```go
package main

import (
    "fmt"
)

func simple(a func(a, b int) int) {
    fmt.Println(a(60, 7))
}

func main() {
    f := func(a, b int) int {
        return a + b
    }
    simple(f)
}
```

Ở ví dụ trên, chúng ta định nghĩa func `simple` với tham số là 1 function, nhận vào 2 tham số kiểu `int` và trả về giá trị kiểu `int`

Trong func `main` chúng ta tạo ra anonymous function `f`. Sau đó gọi `simple` và truyền `f` vào làm đối số cho nó

#### Trả về một function khác

Bây giờ chúng ta hãy viết lại chương trình trên và trả về một function từ function `simple`.

```go
package main

import (
    "fmt"
)

func simple() func(a, b int) int {
    f := func(a, b int) int {
        return a + b
    }
    return f
}

func main() {
    s := simple()
    fmt.Println(s(60, 7))
}
```

Trong ví dụ trên, func `simple` trả về một function nhận hai tham số kiểu int và trả về giá trị kiểu int.

Trong func `main` chúng ta thực hiện gọi `simple` và gán giá trị trả về của `simple` cho biến s. Sau đó gọi biến `s` và truyền vào 2 đối số là 6 và 7

### Closures

Closures là trường hợp đặc biệt của anonymous functions. Closures được hiểu là các anonymous functions truy cập các biến được định nghĩa bên ngoài body của function.

Cùng quan sát ví dụ dưới đây

```go
package main

import (
    "fmt"
)

func main() {
    a := 5
    func() {
        fmt.Println("a =", a)
    }()
}
```

Trong ví dụ trên, anonymous function truy cập vào biến `a` có bên ngoài body của nó. Do đó, anonymous function này là một closure.

Mỗi closure đều bị ràng buộc với biến bao quanh nó. Chúng ta có thể tìm hiểu điều này thông qua ví dụ sau

```go
package main

import (
    "fmt"
)

func appendStr() func(string) string {
    t := "Hello"
    c := func(b string) string {
        t = t + " " + b
        return t
    }
    return c
}

func main() {
    a := appendStr()
    b := appendStr()
    fmt.Println(a("World"))
    fmt.Println(b("Everyone"))

    fmt.Println(a("Gopher"))
    fmt.Println(b("!"))
}
```

Trong chương trình trên, func `appendStr` trả về một closure. `closure` này được ràng buộc với biến `t`.

2 biến `a` và `b` được khai báo bên trong func `main` là closure và chúng bị ràng buộc với giá trị của `t`

Đầu tiên chúng ta gọi func `a` với đối số là `World`. Bây giờ giá trị của `t` trong `a` trở thành `Hello World`.

Tiếp theo chúng ta gọi func `b` với đối số là `Everyone`, vì `b` bị ràng buộc với biến `t`. Nên lúc này giá trị ban đầu của `t` được khởi tạo lại trong b là `Hello`

Nhưng trong lần gọi func `b` tiếp theo thì giá trị ban đầu của `t` trong `b` lại là `Hello Everyone`. Đó là lý do vì sao chúng ta nhận được kết quả sau đây

Kết quả

```
Hello World
Hello Everyone
Hello World Gopher
Hello Everyone !
```
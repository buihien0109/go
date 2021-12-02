## Go HTTP basic

Các ngôn ngữ lập trình hầu hết cung cấp module, package để giúp chúng ta làm quen với http server. Trong Go cũng thế, nó cung cấp sẵn cho chúng ta package net/http bao gồm các phương thức cơ bản để giúp chúng ta bắt đầu tìm hiểu và thực hành với http server cũng như một số chức năng liên quan khác ví dụ : routing, query string, form, template, ...

Trong bài viết này chúng ta sẽ tìm hiểu cách sử dụng package net/http để tạo ra server như thế nào và thực hiện routing đơn giản

Xử lý HTTP request trong Go bao gồm 2 thành phần chính:

- **ServeMux** : về bản chất thì ServeMux là một HTTP request router (multiplexor). Trách nhiệm của nó là nhận các yêu cầu từ phía client và đưa ra hàm xử lý tương ứng
- **Handler** : Chịu trách nhiệm xử lý một request cụ thể và trả về response tương ứng cho client. Trong Go thì bất kỳ object nào cũng có thể là Handler chỉ cần nó implement interface **http.Handler**

Bây giờ chúng ta cùng đi vào một ví dụ cụ thể để tìm hiểu về HTTP server ở trong Go

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home page"))
}

func main() {
    http.HandleFunc("/", homeHandle)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", nil))
}
```

Trong func **main()** chúng ta đăng ký một router với URL Path là **"/"** và hàm xử lý tương ứng với URL Path này là **homeHandle**. Tức là khi client gửi yêu cầu đến URL " **[http://localhost:3000](http://localhost:3000)**" thì func **homeHandle** sẽ được thực thi

```jsx
func homeHandle(w http.ResponseWriter, r *http.Request) {
    // Code
}
```

Tham số truyền vào cho func bao gồm 2 tham số :

- **http.Request** chứa thông tin của Request
- **http.ResponseWriter** chứa thông tin của Response

Phần code xử lý trong hàm này cũng rất đơn giản, ở đây chúng ta trả về plain text cho client sử dụng phương thức **w.Write()**

Cuối cùng chúng ta tạo server và lắng nghe các yêu cầu gửi đến bằng func **ListenAndServe**

Trường hợp bây giờ chúng ta có nhiều router thì sao

Rất đơn giản, chúng ta có thể kiểm tra URL path để có định nghĩa ra các hàm xử lý tương ứng

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        w.Write([]byte("Home page"))
    } else if r.URL.Path == "/about" {
        w.Write([]byte("About page"))
    } else {
        w.Write([]byte("Page not found"))
    }
}

func main() {
    http.HandleFunc("/", homeHandle)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", nil))
}
```

Ở ví dụ trên chúng ta sử dụng **if/else** theo giá trị của **URL Path** để đưa ra các xử lý tương ứng. Tuy nhiên với cách này thì chúng ta phải viết tất cả các các logic xử lý trong một handle duy nhất, dẫn đến tình trạng một hàm mà làm quá nhiều công việc ⇒ khó bảo trì

Rất may là chúng ta có thể giải quyết vấn đề này bằng cách thêm nhiều **http.HandleFunc()**, mỗi hàm **http.HandleFunc()** sẽ xử lý một router tương ứng

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home page"))
}

func aboutHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About page"))
}

func main() {
    http.HandleFunc("/", homeHandle)
    http.HandleFunc("/about", aboutHandle)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", nil))
}
```

Bây giờ lại xảy ra tình trạng nữa là **http.HandleFunc()** mặc định là **Method HTTP GET**, trường hợp chúng ta muốn dùng các Method HTTP khác như **POST, PUT, DELETE**, ... thì phải làm như thế nào?

Thật buồn là package net/http không hỗ trợ chúng ta điều này. Nên nếu muốn sử dụng các Method HTTP khác, chúng ta cần phải kiểm tra Method Type được client gửi lên trong object **http.Request**

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandle(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        w.Write([]byte("Get method"))
    case "POST":
        w.Write([]byte("Post method"))
    case "PUT":
        w.Write([]byte("Put method"))
    case "DELETE":
        w.Write([]byte("Detele method"))
    }
}

func main() {
    http.HandleFunc("/", homeHandle)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", nil))
}
```

Cũng hơi loằng ngoằng nhưng cũng đành chấp nhận số phận nghiệt ngã này vậy. Ở đây chúng ta sẽ lấy Method Type bằng **r.Method**, sau đó có thể áp dụng **if/else** hoặc **switch/case** để xử lý

## ServeMux

Một điều mà các bạn cần lưu ý khi chúng ta khởi tạo server. Trong func **http.ListenAndServe(":3000", nil)**

- Tham số thứ 1 : **address** tức là port mà chúng ta sẽ truy cập
- Tham số thứ 2 : **nil** (trường hợp này sẽ sử dụng **DefaultServeMux** làm handler để xử lý request gửi đến)

Tuy nhiên Go cung cấp sẵn cho chúng ta một handler khác là **ServeMux**. Đây là một ServeMux có sẵn nên chúng ta chỉ việc lôi ra và sử dụng thôi

Cùng đi đến ví dụ sau đây

```go
package main

import (
    "fmt"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home page"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homePage)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", mux))
}
```

- Trong func main chúng ta sử dụng **http.NewServeMux()** để tạo ra một empty ServeMux
- Tiếp theo chúng ta sử dụng func **HandleFunc()** để đăng ký handle xử lý khi có request gửi đến URL Path **"/"**
- Cuối cùng chúng ta tạo ra server mới và bắt đầu lắng nghe request với hàm **http.ListenAndServe()** truyền vào port và handler tương ứng ở đây là **mux**

Ưu điểm khi sử dụng **SeverMux** so với **DefaultServeMux** :

- Với DefaultServeMux cung cấp cho chúng ta global access, tức là chúng ta có thể truy cập ở mọi lúc, mọi nơi. Trong khi **SeverMux** là local access, khi nào cần thì tạo ra đối tượng và sử dụng. Điều này phần nào đảm bảo về mặt security cho ứng dụng
- SeverMux cung cấp nhiều phương thức hơn là DefaultServeMux

## Custom Handler

Ở phần trên, chúng ta đã đề cập đến 2 handler chúng mà golang cung cấp sẵn để lắng nghe và xử lý request là SeverMux và DefaultServeMux

Ngoài ra chúng ta có thể tự Custom một Handler tùy theo nhu cầu. Để custom handler, chúng ta cần implement interface http.Handler

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

Interface **Handler** có một phương thức duy nhất là **ServeHTTP(ResponseWriter,** * **Request)** với 2 tham số truyền vào : **ResponseWriter** và * **Request**

Cùng theo dõi ví dụ sau

```jsx
package main

import (
    "fmt"
    "net/http"
)

type handle struct{}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        w.Write([]byte("Home page"))
    } else if r.URL.Path == "/about" {
        w.Write([]byte("About page"))
    } else {
        w.Write([]byte("Page not found"))
    }
}

func main() {
    handle := new(handle)
    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", handle))
}
```

- Đầu tiên chúng ta khai báo struct handle
- Để implement Interface **Handler** chúng ta sẽ khai báo phương thức **ServeHTTP()** tương ứng với kiểu handle này
- Bên trong phương thức này chỉ là các logic xử lý với các URL path khác nhau, như ở ví dụ trên chúng ta đã trình bày
- Bên trong func **main()**, tương tự như **ServeMux**, chúng ta sẽ khởi tạo 1 đối tượng thuộc kiểu **handle** và truyền nó vào trong func **http.ListenAndServe()**

## Running Mutilple Server

Trong nhiều trường hợp chúng ta cần đối mặt với trường hợp ứng dụng của chúng ta cần mở nhiều port để làm các tác vụ khác nhau như :

- Phục vụ các request HTTP
- Giao tiếp nội bộ
- ...

Vậy trong Go chúng ta sẽ làm như thế nào. Cùng theo dõi ví dụ dưới đây

**Đặt vấn đề**: Chúng ta cần tạo ra 3 server, mỗi server sẽ lắng nghe ở 1 công khác nhau

Đây là cách thông thường chúng ta nghĩ tới

```jsx
package main

import (
    "fmt"
    "net/http"
)

func createServer(port int) *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world at %d", port)
    })

    server := http.Server{
        Addr:    fmt.Sprintf(":%d", port),
        Handler: mux,
    }

    return &server
}

func main() {
    server := createServer(3000)
    fmt.Println(server.ListenAndServe())

    server1 := createServer(3001)
    fmt.Println(server1.ListenAndServe())

    server2 := createServer(3002)
    fmt.Println(server2.ListenAndServe())
}
```

- Func **createServer** giúp chúng ta tạo ra 1 instance của Server. Bên trong chúng ta sử dụng **ServeMux** là handler để xử lý request. Đồng thời đăng ký 1 router với URL Path là **"/"**
- Bên trong func **main()** chúng ta sẽ lập trình đồng bộ để tạo 3 server lắng nghe ở 3 port : **3000**, **3001**, **3002**

**Kết quả**: Chỉ có server ở port 3000 hoạt động còn 3001 và 3002 lại không hoạt động. Vì sao lại thế?

Bởi vì khi chương trình thực hiện đến câu lệnh `server.ListenAndServe()`. Lúc này nó sẽ block lại routine hiện tại (trong trường hợp này là routine main), dẫn đến các câu lệnh đằng sau câu lệnh trên sẽ không được thực hiện

**Cách giải quyết**: Sử dụng Go routine để khởi tạo server

```jsx
func main() {
    wg := new(sync.WaitGroup)

    wg.Add(3)

    go func() {
        server := createServer(3000)
        fmt.Println(server.ListenAndServe())
        wg.Done()
    }()

    go func() {
        server := createServer(3001)
        fmt.Println(server.ListenAndServe())
        wg.Done()
    }()

    go func() {
        server := createServer(3002)
        fmt.Println(server.ListenAndServe())
        wg.Done()
    }()

    wg.Wait()
}
```

Ở đây chúng ta sử dụng WaitGroup sẽ chờ một tập hợp goroutines kết thúc. Khi mà các goroutines chưa được chạy xong, thì WaitGroup sẽ block chương trình tại thời điểm đó

`wg.Add(3)` chính sẽ thêm số goroutines mà nó muốn chờ

Trong mỗi goroutine chúng ta sẽ thực hiện việc khởi tạo server và cho nó lắng nghe ở port mong muốn. Khi thực hiện việc khởi tạo server trong goroutine xong, chúng ta gọi `wg.Done()` để thông báo rằng goroutine này đã chạy xong. Mỗi lần `wg.Done()` sẽ giảm wg đi 1 và cho tới khi wg về 0 wg.Wait() mới bắt đầu cho phép chức năng chạy xuống bên dưới.
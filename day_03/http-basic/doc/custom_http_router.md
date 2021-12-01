Về cơ bản thì HTTP router là một multiplexer (bộ ghép kênh). Nó điều hướng HTTP request tới các hàm xử lý request tương ứng. Trong đó [gorilla/mux](https://github.com/gorilla/mux) là một router thường được sử dụng ở trong golang. Trong thư viện chuẩn của golang (golang standard library) cung cấp cho chúng ta [ServeMUX](https://pkg.go.dev/net/http#ServeMux) để làm điều tương tự.

Để hiểu cách thiết kế router như thế nào, chúng ta hãy xem xét những điểm sau. Một router là một tập hợp các route và một route sẽ có những đặc điểm sau đây.

1. Path
2. Method (GET/POST/PUT/DELETE)
3. Handler

Chúng ta khai báo 2 struct như sau

```go
type Route struct {
   Method  string
   Pattern string
   Handler http.Handler
}

type Router struct {
   routes []Route
}
```

Chúng ta có thể tạo một router mới và đăng ký các route. Để có thể tạo một router mới, chúng ta cần một phương thức `NewRouter`.

```go
func NewRouter() *Router {
   return &Router{}
}
```

Để có thể đăng ký các route chúng ta cần các phương thức `GET, POST, PUT, DELETE`. Tất cả các phương thức sẽ gọi `AddRoute` để append vào `routes`.

```go
func (r *Router) AddRoute(method, path string, handler http.Handler) {
   r.routes = append(r.routes, Route{Method: method, Pattern: path, Handler: handler})
}
```

Sau đó, trong các phương thức chúng ta sẽ gọi `AddRoute`

```go
func (r *Router) GET(path string, handler Handler) {
   r.AddRoute("GET", path, handler)
}

func (r *Router) POST(path string, handler Handler) {
   r.AddRoute("POST", path, handler)
}

func (r *Router) PUT(path string, handler Handler) {
   r.AddRoute("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler Handler) {
   r.AddRoute("DELETE", path, handler)
}
```

Lưu ý: Có thể thấy trong các method trên, thay vì tham số đầu vào là `http.Handler`, chúng ta sử dụng type Handler để implements handler interface. Lý do là nếu chúng ta sử dụng `http.Handler`, chúng ta sẽ không thể implements ServeHTTP cho từng handler.

Vì vậy, chúng ta có thể tự định nghĩa type handler

```go
type Handler func(r *http.Request) (statusCode int, data map[string]interface{})

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    statusCode, data := h(r)
    w.WriteHeader(statusCode)

    json.NewEncoder(w).Encode(data)
}
```

Điều này cho phép chúng ta chuẩn hóa lại cách chúng ta trả về responses cho client.

Bây giờ chúng ta có thể đăng ký các route, nhưng làm cách nào để server sử dụng router để gọi handler tương ứng?

`ListenAndServe` được sử dụng để start một server và `ListenAndServe` nhận vào hai tham số.

- addr
- handler

```go
func ListenAndServe(addr string, handler Handler) error
```

Chúng ta cần làm hai việc ở đây:

1. Cần một method để tìm kiếm các route và gọi handler tương ứng.
2. Router cần có type handler, vì vậy implement handler interface, để router có thể được truyền vào trong ListenAndServe như là một handler.

Vì vậy, chúng ta cần viết một method là `getHandler` lặp qua các route để tìm handler tương ứng.

```go
func (r *Router) getHandler(method, path string) http.Handler {
   for _, route := range r.routes {
      re := regexp.MustCompile(route.Pattern)
      if route.Method == method && re.MatchString(path){
         return route.Handler
      }
   }
   return http.NotFoundHandler()
}
```

Để implement [Handler interface](https://pkg.go.dev/net/http#Handler) , chúng ta cần định nghĩa method `ServeHTTP (ResponseWriter, * Request)` trong router, method này giúp cho router có type Handler mà chúng ta có thể truyền router vào `ListenAndServe`.

Bây giờ, chúng ta cần implement method ServeHTTP

```go
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request){
   path := req.URL.Path
   method := req.Method

   handler := r.getHandler(method, path)

   // handler middlewares viết ở đây

   handler.ServeHTTP(w, req)
}
```

![image](https://techmaster.vn/media/static/9479/c6d626451co50fuc8e00)

- Dòng màu đỏ hiển thị cách một handler được đăng ký trước khi server được chạy. `Custom Handler` là handlers mà chúng ta định nghĩa trong ứng dụng
- Dòng màu xanh lá cây hiển thị cách server tìm một handler đã được đăng ký. Chúng ta truyền router vào trong ListenAndServe và router implements handler interface. Vì vậy, phương thức ServeHTTP sẽ gọi getRoute để trả về handler tương ứng.

Lợi ích của việc chúng ta tự custom HTTP router.

- Chúng ta có thể tự tạo response formats
- Chúng ta có thể viết các middleware. Một số ví dụ về middleware mà chúng ta có thể triển khai
    - CORS
    - Viết 1 middleware để tính toán thời gian thực hiện request
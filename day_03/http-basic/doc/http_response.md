Khi làm việc với Go HTTP, chắc hẳn khi xử lý với các request chúng ta thường trả về cho client nhiều kiểu dữ liệu khác nhau cho client ví dụ như: **Text, Json, Xml, Html, File, Template, ...**

Vì vậy chúng ta cần định dạng kiểu dữ liệu trả về trong header của response với thuộc tính **Content-type** trước khi trả về dữ liệu cho client

Ví dụ:

- **Text** : Content-type : `text/plain`
- **Json** : Content-type : `application/json`
- **Xml** : Content-type : `application/xml`
- **Html** : Content-type : `text/html`
- ...

Bây giờ cùng đi vào các ví dụ dưới đây, để xem cách chúng ta trả về dữ liệu trong Go HTTP như thế nào

### Return Plain Text

```go
package main

import (
    "log"
    "net/http"
)

func returnPlainText(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

func main() {
    http.HandleFunc("/", returnPlainText)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Để trả về plain text cho client chúng ta sử dụng phương thức **Write()**. Ngoài ra cũng có thể sử dụng phương thức tương tự trong package fmt là **Fprintln()** hoặc **Fprintf()**

Mặc định **Content-type** trong header của response là **text-plain** nếu chúng ta không set

Kết quả đạt được là :

![Kết quả sử dụng REST Client](https://techmaster.vn/media/static/9479/c5d75g451co385k2k9t0)

Công cụ chúng ta thực hiện kiểm thử ở trên là extension **REST** client trong **VSCode**

### Một số công cụ để test API

Chúng ta có một số cách để test API:

- Sử dụng trực tiếp trình duyệt
- Sử dụng Postman (recommend)
- Rest client extension (sử dụng trong VSCode)
- Sử dụng Javascript để tạo request phía client
- Sử dụng Golang tạo request phía backend
- Viết unit test
- ...

Ví dụ về các viết unit test sử dụng Golang để gửi request

```go
func TestReturnPlainText(t *testing.T) {
    // Gửi GET request đến URL
    resp, err := http.Get("http://localhost:3000")
    if err != nil {
        log.Fatalln(err)
    }

    // Đọc nội dung body từ response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // Convert body sang kiểu string và so sánh với kết quả mong muốn
    sb := string(body)
    expected := "OK"
    if sb != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",sb, expected)
    }
}
```

- Đầu tiên chúng ta sử dụng **http.Get()** để gửi request đến một URL nào đó. Thông tin response trả về có kiểu là **http.Response**
- Tiếp theo chúng ta đọc nội dung body từ **resp.Body** sử dụng phương thức **ioutil.ReadAll()**
- Cuối cùng so sánh kết quả của body trong response với kết quả mà chúng ta mong muốn

Kết quả sau khi chúng ta run unit test

![Viết unit test golang](https://techmaster.vn/media/static/9479/c5d77fs51co385k2k9tg)

Bên trên chúng ta sử dụng chức năng **run test** trực tiếp trên UI của VSCode để thực hiện chạy một unit test cụ thể nào đó

Trường hợp trong ứng dụng của chúng ta có nhiều unit test, và cần chạy đồng thời cùng một lúc. Thay vì chạy từng unit test một, chúng ta có thể sử dụng câu lệnh sau để chạy đồng thời tất cả các unit có trong project

```go
go test
```

### Return JSON

**JSON** là viết tắt của **JavaScript Object Notation**, là một kiểu định dạng dữ liệu tuân theo một quy luật nhất định mà hầu hết các ngôn ngữ lập trình hiện nay đều có thể đọc được. JSON là một tiêu chuẩn mở để trao đổi dữ liệu trên web.

JSON là kiểu dữ liệu trả về rất phổ biến khi chúng ta thực hiện với các Rest API

Cùng tham khảo ví dụ dưới đây:

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type student struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func returnJson(w http.ResponseWriter, r *http.Request) {
    // Tạo mảng student
    students := []student{
        {Id: 1, Name: "Nguyễn Văn A"},
        {Id: 2, Name: "Trịnh Văn B"},
        {Id: 3, Name: "Ngô Thị C"},
    }

    studentsJson, err := json.Marshal(students)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(studentsJson)

}

func main() {
    http.HandleFunc("/", returnJson)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Trong ví dụ trên:

- Đầu tiên chúng ta đầu tiên chúng ta tạo ra mảng **students** với các phần tử trong mảng có kiểu dữ liệu là **student struct**
- Sử dụng **json.Marshal()** trong package **encoding/json** để convert students sang kiểu JSON
- Cuối cùng trả về kết quả cho Client với **Content-type** là **application/json**

Kết quả khi test bằng REST Client Extension

![Kết quả khi test](https://techmaster.vn/media/static/9479/c5d7r9c51co385k2k9u0)

Khi trả về JSON cho Client trong golang, chúng ta nên đánh tag json cho từng trường trong struct ví dụ:

```go
type student struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}
```

Tại sao cần phải đánh tag json cho các trường này?

- Trường hợp chúng ta không đánh tag json thì mặc định key của JSON Object sẽ tương ứng với key của Struct nhưng sẽ ở dạng Lowercase
- Trường hợp đánh tag json chúng ta có thể thay đổi key của JSON Object

Ví dụ:

```go
type student struct {
    Id   int    `json:"student_id"`
    Name string `json:"full_name"`
}
```

Kết quả là:

![Ví dụ về JSON Response](https://techmaster.vn/media/static/9479/c5elfok51co6ehvera30)

Lúc này các key trong JSON Object sẽ được thay đổi tương ứng với tag json mà chúng ta sử dụng trong struct

### Return XML

XML là từ viết tắt của từ **Extensible Markup Language** là ngôn ngữ đánh dấu mở rộng. XML có chức năng truyền dữ liệu và mô tả nhiều loại dữ liệu khác nhau

XML được sử dụng để mô tả dữ liệu dưới dạng text, nên hầu hết các phần mềm hay các chương trình bình thường đều có thể đọc được chúng.

Tuy nhiên trong các ứng dụng REST việc trả về dữ liệu kiểu XML rất ít khi được áp dụng

```go
package main

import (
    "encoding/xml"
    "log"
    "net/http"
)

type student struct {
    Id   int
    Name string
}

func returnXml(w http.ResponseWriter, r *http.Request) {

    students := []student{
        {Id: 1, Name: "Nguyễn Văn A"},
        {Id: 2, Name: "Trịnh Văn B"},
        {Id: 3, Name: "Ngô Thị C"},
    }

    studetsXml, err := xml.Marshal(students)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/xml")
    w.Write(studetsXml)
}

func main() {
    http.HandleFunc("/", returnXml)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

XML được sử dụng tương tự như JSON chỉ khác giá trị của thuộc tính **Content-type** là `application/xml`

Kết quả khi test

![Kết quả khi test](https://techmaster.vn/media/static/9479/c5d7s9451co385k2k9ug)

### Return HTML

Trả về mã HTML là vô cùng phổ biến với các ứng dụng sử dụng cơ chế `SSR` ( **Server Side Rendering**). Theo cơ chế này thì hầu hết các xử lý logic đều ở phía server. Trong đó, server thực hiện xử lý và tiến hành các thao tác với cơ sở dữ liệu để render ra thành mã HTML, sau đó gửi response cho client. Phía client chỉ cần hiển thị mã HTML này.

Cùng tham khảo ví dụ dưới đây

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func returnHtml(w http.ResponseWriter, r *http.Request) {
    html := `
        <!DOCTYPE html>
        <html lang="en">
            <head>
                <meta charset="UTF-8">
                <meta http-equiv="X-UA-Compatible" content="IE=edge">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>Books</title>
            </head>
            <body>
                <h1>Những cuốn sách được yêu thích</h2>
                <ul>
                    <li>Những người khốn khổ</li>
                    <li>Đắc nhân tâm</li>
                </ul>
            </body>
        </html>
    `

    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, html)
}

func main() {
    http.HandleFunc("/", returnHtml)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Ở phần Header của response chúng ta set giá trị của thuộc tính **Content-Type** là `text/html`

Kết quả khi test

![Kết quả khi test](https://techmaster.vn/media/static/9479/c5d7u1s51co385k2k9v0)

### Return File

```go
package main

import (
    "log"
    "net/http"
    "path"
)

func returnFile(w http.ResponseWriter, r *http.Request) {
    filePath := path.Join("view", "index.html")
    http.ServeFile(w, r, filePath)
}

func main() {
    http.HandleFunc("/", returnFile)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Cấu trúc thư mục phần này của chúng ta như sau:

```arduino
.
├── go.mod
├── go.sum
├── main.go
└── view
    └── index.html
```

Kết quả chúng ta được như sau

![Trả về file tĩnh](https://techmaster.vn/media/static/9479/c5ej05c51co6ehvera0g)
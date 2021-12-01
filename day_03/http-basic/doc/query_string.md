## Query string là gì

Query string là tập hợp 1 hoặc nhiều **cặp key-value** trong URL và theo sau dấu `?`, trường hợp nếu URL có nhiều **cặp key-value** thì những **cặp key-value** sẽ được ngăn cách với nhau bằng dấu `&`

Ví dụ: [https://domain.com/shoes?type=sneakers](https://domain.com/shoes?type=sneakers)

![ảnh minh họa query string](https://techmaster.vn/media/static/9479/c5sfocs51co8fjgguc50)

Quan sát hình ảnh demo chúng ta có thể thấy query string bao gồm các thành phần sau:

- **Ký hiệu ? (Question Mark)** : Chỉ định vị trí bắt đầu của query string (domain.com/shoes`?`type=sneakers)
- **Ký hiệu & (Ampersand)** : Giúp phân tách tham số nếu trong URL có nhiều parameter (domain.com/shoes?type=sneakers`&`sort=price_ascending)
- **Key** : Nó giống như title hoặc label của parameter (domain.com/shoes?`type`=sneakers)
- **Value** : Chỉ định value cho 1 key cụ thể (domain.com/shoes?type=`sneakers`)

Trong bài viết này chúng ta cùng tìm hiểu về Query string trong golang HTTP

## Ví dụ demo

**1.Khởi tạo project**

Tạo thư mục thực hành, ở đây chúng ta sẽ tạo thư mục **go-querystring**

```
mkdir go-querystring
```

Tiếp theo tạo **go module** để quản lý các dependency nếu có trong dự án golang

```
go mod init go-querystring
```

Và cuối cùng tạo file **main.go** đây là file chạy chính trong chương trình go

```
touch main.go
```

**2.Khởi tạo web server và định nghĩa router**

```go
package main

import (
    "log"
    "net/http"
)

func bookHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("book handle"))
}

func main() {
    http.HandleFunc("/", bookHandle)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
```

Trong func main chúng ta khởi tạo server sử dụng build-in package của golang là **net/http** và cho nó lắng nghe request ở cổng **3000**

```go
log.Fatal(http.ListenAndServe(":3000", nil))
```

Đồng thời định nghĩa thêm API **"/"** sử dụng method **http.HandleFunc**, khi client gửi request đến API này, server sẽ gọi func **bookHandle** để xử lý request

```go
http.HandleFunc("/", bookHandle)
```

Tạm thời trong func **bookHandle** chúng ta chưa làm gì đặc biệt cả, mà chỉ trả về cho client message " **book handle**"

```go
func bookHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("book handle"))
}
```

Bây giờ chúng ta sẽ chạy file **main.go** và thử gửi request đến API mà chúng ta đã định nghĩa ở trên để xem kết quả trả về là gì

```
curl -i localhost:3000
```

![test API](https://techmaster.vn/media/static/9479/c5sg37451co8fjgguc60)

Khi gửi request đến server sử dụng **curl** chúng ta nhận được kết quả là " **book handle**" và response này có Content-type mặc định là " **text/plain**". Ngoài sử dụng curl để test, chúng ta có thể sử dụng các công cụ khác như:

- Postman
- REST Client (extension trong VSCode)
- Sử dụng trực tiếp browser
- Viết Golang client để gửi request
- ...

Tiếp theo chúng ta định nghĩa cấu trúc dữ liệu là book struct

```go
type book struct {
    ID     int        // ID của sách
    Name   string     // Tên sách
    Author string     // Tác giả
    Year   int        // Năm xuất bản
}
```

Trong ví dụ này, chúng ta thực hiện sử dụng query string để lọc dữ liệu book theo 2 tiêu chí là `name` và `year`

Bây giờ chúng ta sẽ refactor lại func **bookHandle** để phù hợp với yêu cầu

```go
func bookHandle(w http.ResponseWriter, r *http.Request) {
    // Trả về thông tin của Query dưới dạng map[string][string]
    query := r.URL.Query()

    // Mockup 1 mảng các book
    books := []book{
        {ID: 1, Name: "book1", Author: "author1", Year: 2020},
        {ID: 2, Name: "book2", Author: "author2", Year: 2021},
        {ID: 3, Name: "book3", Author: "author3", Year: 2019},
        {ID: 4, Name: "book4", Author: "author4", Year: 2020},
    }

    // Khai báo biến booksReturn chứa dữ liệu book sau khi lọc với query string
    var booksReturn []book

    // Kiểm tra xem key name có tồn tại trong query string
    name, ok := query["name"]

    // Xử lý khi có tồn tại key name
    if ok {
        filterName := strings.Join(name, ",")

        // Gọi func filterByName để lọc ra các kết quả phù hợp
        booksReturn = filterByName(books, filterName)
    } else {

        // Nếu không có thì gán giá trị của books cho booksReturn
        booksReturn = append(booksReturn, books...)
    }

    // Làm tương tự như với key name
    year, ok := query["year"]
    if ok {
        filterYear, _ := strconv.Atoi(strings.Join(year, ""))
        booksReturn = filterByYear(booksReturn, filterYear)
    }

    // Marshal booksReturn sang kiểu JSON
    booksJson, err := json.Marshal(booksReturn)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(booksJson)
}
```

Bổ sung thêm 2 helper function phục vụ cho quá trình lọc dữ liệu

- filterByName (lọc book theo tên)
- filterByYear (lọc book theo năm xuất bản)

```go
// Lọc book theo tên
func filterByName(books []book, name string) (result []book) {
    for _, book := range books {
        if book.Name == name {
            result = append(result, book)
        }
    }
    return result
}

// Lọc book theo year
func filterByYear(books []book, year int) (result []book) {
    for _, book := range books {
        if book.Year == year {
            result = append(result, book)
        }
    }
    return result
}
```

## Testing API

### 1. Không sử dụng query string

Trường hợp này chúng ta thử gửi GET request mà không sử dụng query string

```
curl -i http://localhost:3000 | json
```

**Kết quả**

![Test 1](https://techmaster.vn/media/static/9479/c5sgnuk51co8fjgguc80)

### 2. Sử dụng query string

Tiếp theo chúng ta test thêm trường hợp sử dụng query string để lọc dữ liệu

**Ví dụ 1:**

```
curl -i http://localhost:3000?year=2020 | json
```

**Kết quả**

![Test 2](https://techmaster.vn/media/static/9479/c5sgq9k51co8fjgguc8g)

* * *

**Ví dụ 2:**

```
curl -i http://localhost:3000?year=2020&name=book1 | json
```

**Kết quả**

![Test 3](https://techmaster.vn/media/static/9479/c5sgqvs51co8fjgguc90)
package repo

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Thông số kết nối đến CSDL
var (
	host     string = "localhost"
	port     string = "3306"
	username string = "root"
	password string = "123"
	database string = "db"
)

var (
	DB     *gorm.DB   // Kết nối đến CSDL
	random *rand.Rand // Đối tượng dùng để tạo random number
)

func init() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	// Kết nối với CSDL thông qua connection string
	var err error
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
	})

	// Xử lý nếu quá trình kết nối với CSDL bị lỗi
	if err != nil {
		panic("Failed to connect database")
	}

	//Khởi động engine sinh số ngẫu nhiên
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)
}

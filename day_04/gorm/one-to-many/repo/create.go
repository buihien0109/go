package repo

import (
	"one-many/model"

	"github.com/brianvoe/gofakeit/v6"
)

func CreateData() (err error) {
	// Tạo mảng foo để lưu kết quả
	var foos []model.Foo

	// Sử dụng vòng lặp để tạo 1 số đối tượng foo
	for i := 0; i < 5; i++ {
		foo_id := NewID() // random fooID
		foo := model.Foo{
			Id:   foo_id,
			Name: gofakeit.Animal(), // Fake name
		}

		// Với mỗi đối tượng foo -> tạo 1 số đối tượng bar tương ứng
		for j := 0; j < 2+random.Intn(2); j++ {
			bar := model.Bar{
				Id:    NewID(),
				Name:  gofakeit.Animal(),
				FooId: foo_id,
			}

			foo.Bars = append(foo.Bars, bar)
		}

		foos = append(foos, foo)
	}

	// Insert vào trong CSDL
	if err := DB.Create(&foos).Error; err != nil {
		return err
	}

	return nil
}

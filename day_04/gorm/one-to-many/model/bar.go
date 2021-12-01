package model

type Foo struct {
	Id   string `gorm:"primaryKey"`
	Name string
	Bars []Bar `gorm:"foreignKey:FooId"`
}

type Bar struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	FooId string `gorm:"column:foo_id"`
	Foo   Foo
}

func (f *Foo) TableName() string {
	return "foo"
}

func (b *Bar) TableName() string {
	return "bar"
}

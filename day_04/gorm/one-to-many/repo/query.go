package repo

import "one-many/model"

func GetFooBar() (foos []model.Foo, err error) {
	if err := DB.Preload("Bars").Find(&foos).Error; err != nil {
		return nil, err
	}
	return foos, nil
}

func GetFooById(id string) (foo model.Foo, err error) {
	if err := DB.Preload("Bars").Find(&foo, "foo.id = ?", id).Error; err != nil {
		return model.Foo{}, err
	}

	return foo, nil
}

func GetBarById(id string) (bar model.Bar, err error) {
	bar = model.Bar{
		Id: id,
	}

	if err = DB.Preload("Foo").Find(&bar).Error; err != nil {
		return model.Bar{}, err
	}

	return bar, nil
}

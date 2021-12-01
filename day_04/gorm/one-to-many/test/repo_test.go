package test

import (
	"fmt"
	"one-many/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFooBar(t *testing.T) {
	foos, err := repo.GetFooBar()
	assert.Nil(t, err)
	for _, foo := range foos {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println("  " + bar.Name)
		}
	}
	assert.Positive(t, len(foos))
}

func Test_GetFooById(t *testing.T) {
	foo, err := repo.GetFooById("ox-01")
	assert.Nil(t, err)
	fmt.Println(foo.Name)
	for _, bar := range foo.Bars {
		fmt.Println("  " + bar.Name)
	}
}

func Test_GetBarById(t *testing.T) {
	bar, err := repo.GetBarById("bar1")
	assert.Nil(t, err)

	fmt.Println(bar.FooId)
	fmt.Println("  " + bar.Id)
	fmt.Println("  " + bar.Name)
}

func Test_CreateData(t *testing.T) {
	err := repo.CreateData()
	assert.Nil(t, err)
}

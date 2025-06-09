package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, error := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, error)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category {
		Id: "2",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category) // mengembalikan category jadi jika Id category diganti 1 tetap pass, jadi tidak ada error karena sama sama mengembalikan category yang sama

	result, error := categoryService.Get("2")
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id) // sebaiknya : (t, "2", result.Id) 
	assert.Equal(t, category.Name, result.Name) // sebaiknya : (t, "Laptop", result.Name)
}
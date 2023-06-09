package service

import (
	"go-jwt/entity"
	"go-jwt/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepositoryMock = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepositoryMock}

func TestProductServiceGetOneProductNotFound (t *testing.T) {
	productRepositoryMock.Mock.On("FindById", "1").Return(nil)
	
	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "Product not found", err.Error(), "Error message should be same")
}


func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id : "2",
		Name : "Kaca Mata",
	}

	productRepositoryMock.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Kaca Mata", result.Name, "Product name should be same")
	assert.Equal(t, "2", result.Id, "Product id should be same")
	assert.Equal(t, &product, result, "Product should be same")

	
}
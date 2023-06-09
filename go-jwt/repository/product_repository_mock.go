package repository

import (
	"go-jwt/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(Id string) *entity.Product {
	arguments := repository.Mock.Called(Id)
	if arguments.Get(0) == nil {
		return nil
	}
	product := arguments.Get(0).(entity.Product)
	return &product
}
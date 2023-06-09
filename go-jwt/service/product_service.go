package service

import (
	"errors"
	"go-jwt/entity"
	"go-jwt/repository"

)


type ProductService struct {

	Repository repository.ProductRepository

}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error ){
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("Product not found")
	}

	return product, nil

}
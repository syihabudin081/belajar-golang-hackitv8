package repository

import "go-jwt/entity"

type ProductRepository interface {
	FindById(Id string) *entity.Product
}
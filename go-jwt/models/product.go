package models

import (
	"github.com/asaskevich/govalidator"
"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title string `gorm:"type:varchar(100);not null" json:"title" valid:"required~ Your title of product is required"`
	Description string `gorm:"type:varchar(100);not null" json:"description" valid:"required~ Your description of product is required"`
	UserID uint
	User *User
}


func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {

	_,err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	err = nil
	return

}


func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_,err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}
	err = nil
	return
}
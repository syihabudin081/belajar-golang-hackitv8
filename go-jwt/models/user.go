package models

import (
	"go-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName	string `gorm:"not null" json:"full_name" form:"full_name" valid:"required~ Your Full name is required"`
	Email		string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~ Your email is not valid"`
	PassWord	string `gorm:"not null" json:"password" form:"password" valid:"required~ Your password is required,minstringlength(6)~ Your password must be at least 6 characters"`
	Products	[]Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	Admin		bool `gorm:"default:false" json:"admin"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_,err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.PassWord = helpers.HashPassword(u.PassWord)
	err = nil
	return
}
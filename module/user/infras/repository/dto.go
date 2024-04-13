package repository

import (
	"github.com/google/uuid"
	userdomain "my-shop/module/user/domain"
)

type UserDTO struct {
	Id          uuid.UUID `gorm:"column:id;"`
	FirstName   string    `gorm:"column:first_name;"`
	LastName    string    `gorm:"column:last_name;"`
	Address     string    `gorm:"column:address;"`
	PhoneNumber string    `gorm:"column:phone_number;"`
	Email       string    `gorm:"column:email;"`
	Password    string    `gorm:"column:password;"`
	Salt        string    `gorm:"column:salt;"`
	Role        string    `gorm:"column:role;"`
}

func (dto *UserDTO) ToEntity() (user *userdomain.User, err error) {
	return userdomain.NewUser(dto.Id, dto.FirstName, dto.LastName, dto.Address, dto.PhoneNumber, dto.Email, dto.Password, dto.Salt, userdomain.GetRole(dto.Role))
}

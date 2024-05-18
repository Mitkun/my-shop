package userrepository

import (
	"github.com/google/uuid"
	userdomain "my-shop/module/user/domain"
	"time"
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
	Status      string    `gorm:"column:status;"`
	Role        string    `gorm:"column:role;"`
}

func (dto *UserDTO) ToEntity() (user *userdomain.User, err error) {
	return userdomain.NewUser(dto.Id, dto.FirstName, dto.LastName, dto.Address, dto.PhoneNumber, dto.Email, dto.Password, dto.Salt, dto.Status, userdomain.GetRole(dto.Role))
}

type SessionDTO struct {
	Id           uuid.UUID `gorm:"column:id;"`
	UserId       uuid.UUID `gorm:"column:user_id;"`
	RefreshToken string    `gorm:"column:refresh_token;"`
	AccessExpAt  time.Time `gorm:"column:access_exp_at;"`
	RefreshExpAt time.Time `gorm:"column:refresh_exp_at;"`
}

func (dto SessionDTO) ToEntity() (*userdomain.Session, error) {
	s := userdomain.NewSession(dto.Id, dto.UserId, dto.RefreshToken, dto.AccessExpAt, dto.RefreshExpAt)

	return s, nil
}

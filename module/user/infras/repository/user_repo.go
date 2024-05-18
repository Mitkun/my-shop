package userrepository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"my-shop/common"
	userdomain "my-shop/module/user/domain"
)

const TbUserName = "users"

type userMySQLRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userMySQLRepo {
	return userMySQLRepo{db: db}
}

func (repo userMySQLRepo) FindByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	var dto UserDTO

	if err := repo.db.Table(TbUserName).Where("email = ?", email).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	return dto.ToEntity()
}

func (repo userMySQLRepo) FindByID(ctx context.Context, id uuid.UUID) (*userdomain.User, error) {
	var dto UserDTO

	if err := repo.db.Table(TbUserName).Where("id = ?", id).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}

	return dto.ToEntity()
}

func (repo userMySQLRepo) Create(ctx context.Context, data *userdomain.User) error {
	dto := UserDTO{
		Id:          data.Id(),
		FirstName:   data.FirstName(),
		LastName:    data.LastName(),
		Address:     data.Address(),
		PhoneNumber: data.PhoneNumber(),
		Email:       data.Email(),
		Password:    data.Password(),
		Salt:        data.Salt(),
		Role:        data.Role().String(),
	}

	if err := repo.db.Table(TbUserName).Create(&dto).Error; err != nil {
		return err
	}

	return nil
}

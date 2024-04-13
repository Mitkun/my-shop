package userusecase

import (
	"context"
	"errors"
	"my-shop/common"
	userdomain "my-shop/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistrationDTO) error
}

type Hashes interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
}

type useCase struct {
	repo   UserRepository
	hashes Hashes
}

func NewUseCase(repo UserRepository, hashes Hashes) UseCase {
	return &useCase{repo: repo, hashes: hashes}
}

func (uc *useCase) Register(ctx context.Context, dto EmailPasswordRegistrationDTO) error {
	user, err := uc.repo.FindByEmail(ctx, dto.Email)
	if user != nil {
		return userdomain.ErrEmailHasExisted
	}

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return err
	}

	salt, err := uc.hashes.RandomStr(30)
	if err != nil {
		return err
	}

	hashedPassword, err := uc.hashes.HashPassword(salt, dto.Password)
	if err != nil {
		return err
	}

	userEntity, err := userdomain.NewUser(
		common.GenUUID(),
		dto.FirstName,
		dto.LastName,
		dto.Address,
		dto.PhoneNumber,
		dto.Email,
		hashedPassword,
		salt,
		userdomain.RoleAdmin,
	)
	if err != nil {
		return err
	}

	if err := uc.repo.Create(ctx, userEntity); err != nil {
		return err
	}
	return nil
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
	Create(ctx context.Context, data *userdomain.User) error
}

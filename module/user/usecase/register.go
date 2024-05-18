package userusecase

import (
	"context"
	"errors"
	"github.com/viettranx/service-context/core"
	"my-shop/common"
	userdomain "my-shop/module/user/domain"
)

type registerUC struct {
	userQueryRepo UserQueryRepository
	userCmdRepo   UserCommandRepository
	hashes        Hashes
}

func NewRegisterUC(userQueryRepo UserQueryRepository, userCmdRepo UserCommandRepository, hashes Hashes) *registerUC {
	return &registerUC{userQueryRepo: userQueryRepo, userCmdRepo: userCmdRepo, hashes: hashes}
}

func (uc *registerUC) Register(ctx context.Context, dto EmailPasswordRegistrationDTO) error {
	user, err := uc.userQueryRepo.FindByEmail(ctx, dto.Email)
	if user != nil {
		return core.ErrBadRequest.WithError(userdomain.ErrEmailHasExisted.Error())
	}

	if err != nil && !errors.Is(err, common.ErrRecordNotFound) {
		return core.ErrInternalServerError.WithError("cannot register right now").WithDebug(err.Error())
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
		dto.Status,
		userdomain.RoleAdmin,
	)
	if err != nil {
		return err
	}

	if err := uc.userCmdRepo.Create(ctx, userEntity); err != nil {
		return err
	}
	return nil
}

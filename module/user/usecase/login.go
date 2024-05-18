package userusecase

import (
	"context"
	"github.com/viettranx/service-context/core"
	"my-shop/common"
	userdomain "my-shop/module/user/domain"
	"time"
)

type loginUC struct {
	userRepo      UserQueryRepository
	sessionRepo   SessionCommandRepository
	tokenProvider TokenProvider
	hashes        Hashes
}

func NewLoginUC(userRepo UserQueryRepository, sessionRepo SessionCommandRepository, tokenProvider TokenProvider, hashes Hashes) *loginUC {
	return &loginUC{userRepo: userRepo, sessionRepo: sessionRepo, tokenProvider: tokenProvider, hashes: hashes}
}

func (uc *loginUC) Login(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error) {

	user, err := uc.userRepo.FindByEmail(ctx, dto.Email)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	if ok := uc.hashes.CompareHashPassword(user.Password(), user.Salt(), dto.Password); !ok {
		return nil, core.ErrForbidden.WithError(userdomain.ErrInvalidEmailPassword.Error())
	}

	userId := user.Id()
	sessionId := common.GenUUID()

	accessToken, err := uc.tokenProvider.IssueToken(ctx, sessionId.String(), userId.String())
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	refreshToken, err := uc.hashes.RandomStr(16)
	if err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	tokenExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.TokenExpireInSeconds()))
	refreshExpAt := time.Now().UTC().Add(time.Second * time.Duration(uc.tokenProvider.RefreshExpireInSeconds()))

	session := userdomain.NewSession(sessionId, userId, refreshToken, tokenExpAt, refreshExpAt)
	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, core.ErrInternalServerError.WithError(err.Error())
	}

	return &TokenResponseDTO{
		AccessToken:       accessToken,
		AccessTokenExpIn:  uc.tokenProvider.TokenExpireInSeconds(),
		RefreshToken:      refreshToken,
		RefreshTokenExpIn: uc.tokenProvider.RefreshExpireInSeconds(),
	}, nil
}

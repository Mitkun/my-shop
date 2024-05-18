package userusecase

import (
	"context"
	"github.com/google/uuid"
	userdomain "my-shop/module/user/domain"
)

type UseCase interface {
	Register(ctx context.Context, dto EmailPasswordRegistrationDTO) error
	Login(ctx context.Context, dto EmailPasswordLoginDTO) (*TokenResponseDTO, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponseDTO, error)
}

type Hashes interface {
	RandomStr(length int) (string, error)
	HashPassword(salt, password string) (string, error)
	CompareHashPassword(hashedPassword, salt, password string) bool
}

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	TokenExpireInSeconds() int
	RefreshExpireInSeconds() int
}

type useCase struct {
	*registerUC
	*loginUC
	*refreshTokenUC
}

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCmdRepo() UserCommandRepository
	BuildHashes() Hashes
	BuildTokenProvider() TokenProvider
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCmdRepo() SessionCommandRepository
	BuildSessionRepo() SessionRepository
}

func UseCaseWithBuilder(b Builder) UseCase {
	return &useCase{
		loginUC:        NewLoginUC(b.BuildUserQueryRepo(), b.BuildSessionCmdRepo(), b.BuildTokenProvider(), b.BuildHashes()),
		registerUC:     NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCmdRepo(), b.BuildHashes()),
		refreshTokenUC: NewRefreshTokenUC(b.BuildUserQueryRepo(), b.BuildSessionRepo(), b.BuildTokenProvider(), b.BuildHashes()),
	}
}

type UserRepository interface {
	UserQueryRepository
	UserCommandRepository
}
type UserQueryRepository interface {
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*userdomain.User, error)
}

type UserCommandRepository interface {
	Create(ctx context.Context, data *userdomain.User) error
}

type SessionRepository interface {
	SessionQueryRepository
	SessionCommandRepository
}

type SessionQueryRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*userdomain.Session, error)
	FindByRefreshToken(ctx context.Context, rt string) (*userdomain.Session, error)
}

type SessionCommandRepository interface {
	Create(ctx context.Context, data *userdomain.Session) error
	Delete(ctx context.Context, id uuid.UUID) error
}

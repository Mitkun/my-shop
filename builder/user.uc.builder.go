package builder

import (
	"gorm.io/gorm"
	"my-shop/common"
	userrepository "my-shop/module/user/infras/repository"
	userusecase "my-shop/module/user/usecase"
)

type simpleBuilder struct {
	db *gorm.DB
	tp userusecase.TokenProvider
}

func NewSimpleBuilder(db *gorm.DB, tp userusecase.TokenProvider) simpleBuilder {
	return simpleBuilder{db: db, tp: tp}
}

func (s simpleBuilder) BuildUserQueryRepo() userusecase.UserQueryRepository {
	return userrepository.NewUserRepo(s.db)
}

func (s simpleBuilder) BuildUserCmdRepo() userusecase.UserCommandRepository {
	return userrepository.NewUserRepo(s.db)
}

func (simpleBuilder) BuildHashes() userusecase.Hashes {
	return &common.Hasher{}
}

func (s simpleBuilder) BuildTokenProvider() userusecase.TokenProvider {
	return s.tp
}

func (s simpleBuilder) BuildSessionQueryRepo() userusecase.SessionQueryRepository {
	return userrepository.NewSessionMySQLRepo(s.db)
}

func (s simpleBuilder) BuildSessionCmdRepo() userusecase.SessionCommandRepository {
	return userrepository.NewSessionMySQLRepo(s.db)
}

func (s simpleBuilder) BuildSessionRepo() userusecase.SessionRepository {
	return userrepository.NewSessionMySQLRepo(s.db)
}

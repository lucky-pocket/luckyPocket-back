package usecase

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
)

type Deps struct {
	UserRepository   domain.UserRepository
	NoticeRepository domain.NoticeRepository
}

type userUseCase struct{ *Deps }

func NewUserUseCase(deps *Deps) domain.UserUseCase {
	return &userUseCase{deps}
}

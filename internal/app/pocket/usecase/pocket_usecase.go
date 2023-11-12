package usecase

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/event"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
)

type Deps struct {
	UserRepository   domain.UserRepository
	PocketRepository domain.PocketRepository
	TxManager        tx.Manager
	EventManager     event.EvntManager
}

type pocketUseCase struct{ *Deps }

func NewPocketUseCase(deps *Deps) domain.PocketUseCase {
	return &pocketUseCase{deps}
}

package stubs

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
)

type PocketRepository struct {
	*mocks.PocketRepository
	*stubTransactor
}

func NewPocketRepository(t *testing.T) *PocketRepository {
	return &PocketRepository{
		stubTransactor:   &stubTransactor{},
		PocketRepository: mocks.NewPocketRepository(t),
	}
}

type UserRepository struct {
	*mocks.UserRepository
	*stubTransactor
}

func NewUserRepository(t *testing.T) *UserRepository {
	return &UserRepository{
		stubTransactor: &stubTransactor{},
		UserRepository: mocks.NewUserRepository(t),
	}
}

type NoticeRepository struct {
	*mocks.NoticeRepository
	*stubTransactor
}

func NewNoticeRepository(t *testing.T) *NoticeRepository {
	return &NoticeRepository{
		stubTransactor:   &stubTransactor{},
		NoticeRepository: mocks.NewNoticeRepository(t),
	}
}

type BlackListRepository struct {
	*mocks.BlackListRepository
	*stubTransactor
}

func NewBlackListRepository(t *testing.T) *BlackListRepository {
	return &BlackListRepository{
		stubTransactor:      &stubTransactor{},
		BlackListRepository: mocks.NewBlackListRepository(t),
	}
}

type GameLogRepository struct {
	*mocks.GameLogRepository
	*stubTransactor
}

func NewGameLogRepository(t *testing.T) *GameLogRepository {
	return &GameLogRepository{
		stubTransactor:    &stubTransactor{},
		GameLogRepository: mocks.NewGameLogRepository(t),
	}
}

type TicketRepository struct {
	*mocks.TicketRepository
	*stubTransactor
}

func NewTicketRepository(t *testing.T) *TicketRepository {
	return &TicketRepository{
		stubTransactor:   &stubTransactor{},
		TicketRepository: mocks.NewTicketRepository(t),
	}
}

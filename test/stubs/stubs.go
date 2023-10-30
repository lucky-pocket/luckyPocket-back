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

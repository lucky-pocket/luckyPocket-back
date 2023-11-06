package usecase_test

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/lucky-pocket/luckyPocket-back/test/stubs"
	"github.com/stretchr/testify/suite"
)

type PocketUseCaseTestSuite struct {
	suite.Suite
	uc                   domain.PocketUseCase
	mockPocketRepository *stubs.PocketRepository
	mockUserRepository   *stubs.UserRepository
	mockEventManager     *mocks.EvntManager
}

func TestPocketUseCaseSuite(t *testing.T) {
	suite.Run(t, new(PocketUseCaseTestSuite))
}

func (s *PocketUseCaseTestSuite) SetupSuite() {
	s.mockPocketRepository = stubs.NewPocketRepository(s.T())
	s.mockUserRepository = stubs.NewUserRepository(s.T())
	s.mockEventManager = mocks.NewEvntManager(s.T())

	s.uc = usecase.NewPocketUseCase(&usecase.Deps{
		UserRepository:   s.mockUserRepository,
		PocketRepository: s.mockPocketRepository,
		EventManager:     s.mockEventManager,
		TxManager:        tx.NewTxManager(),
	})
}

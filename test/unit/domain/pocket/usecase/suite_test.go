package usecase_test

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/pocket/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/suite"
)

type PocketUseCaseTestSuite struct {
	suite.Suite
	uc                   domain.PocketUseCase
	mockPocketRepository *mocks.PocketRepository
	mockUserRepository   *mocks.UserRepository
	mockTxManager        *mocks.Manager
}

func TestPocketUseCaseSuite(t *testing.T) {
	suite.Run(t, new(PocketUseCaseTestSuite))
}

func (s *PocketUseCaseTestSuite) SetupSuite() {
	s.mockPocketRepository = mocks.NewPocketRepository(s.T())
	s.mockUserRepository = mocks.NewUserRepository(s.T())
	s.mockTxManager = mocks.NewManager(s.T())

	s.uc = usecase.NewPocketUseCase(&usecase.Deps{
		UserRepository:   s.mockUserRepository,
		PocketRepository: s.mockPocketRepository,
		TxManager:        s.mockTxManager,
	})
}

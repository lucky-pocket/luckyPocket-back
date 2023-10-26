package usecase_test

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/user/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/stubs"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	uc                   domain.UserUseCase
	mockUserRepository   *stubs.UserRepository
	mockNoticeRepository *stubs.NoticeRepository
}

func TestPocketUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (s *UserUseCaseTestSuite) SetupSuite() {
	s.mockUserRepository = stubs.NewUserRepository(s.T())
	s.mockNoticeRepository = stubs.NewNoticeRepository(s.T())

	s.uc = usecase.NewUserUseCase(&usecase.Deps{
		UserRepository:   s.mockUserRepository,
		NoticeRepository: s.mockNoticeRepository,
	})
}

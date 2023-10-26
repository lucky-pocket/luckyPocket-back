package usecase_test

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/notice/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/stubs"
	"github.com/stretchr/testify/suite"
)

type NoticeUseCaseTestSuite struct {
	suite.Suite
	uc                   domain.NoticeUseCase
	mockNoticeRepository *stubs.NoticeRepository
}

func TestPocketUseCaseSuite(t *testing.T) {
	suite.Run(t, new(NoticeUseCaseTestSuite))
}

func (s *NoticeUseCaseTestSuite) SetupSuite() {
	s.mockNoticeRepository = stubs.NewNoticeRepository(s.T())

	s.uc = usecase.NewNoticeUseCase(&usecase.Deps{
		NoticeRepository: s.mockNoticeRepository,
	})
}

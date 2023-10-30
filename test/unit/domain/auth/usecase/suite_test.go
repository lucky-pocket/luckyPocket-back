package usecase

import (
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/onee-only/gauth-go"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/auth/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/test/stubs"
	"github.com/stretchr/testify/suite"
)

type MockGAuthClient struct {
	mock.Mock
}

func (m *MockGAuthClient) IssueToken(code string) (access, refresh string, err error) {
	args := m.Called(code)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockGAuthClient) GetUserInfo(accessToken string) (*gauth.UserInfo, error) {
	args := m.Called(accessToken)
	return args.Get(0).(*gauth.UserInfo), args.Error(1)
}

type AuthUseCaseTestSuite struct {
	suite.Suite
	uc                      domain.AuthUseCase
	mockBlackListRepository *stubs.BlackListRepository
	mockUserRepository      *stubs.UserRepository
	mockJwtIssuer           *mocks.Issuer
	mockJwtParser           *mocks.Parser
	mockGAuthClient         *MockGAuthClient
}

func TestAuthUseCaseSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (l *AuthUseCaseTestSuite) SetupSuite() {
	l.mockBlackListRepository = stubs.NewBlackListRepository(l.T())
	l.mockUserRepository = stubs.NewUserRepository(l.T())
	l.mockJwtIssuer = mocks.NewIssuer(l.T())
	l.mockJwtParser = mocks.NewParser(l.T())
	l.mockGAuthClient = new(MockGAuthClient)

	l.uc = usecase.NewAuthUseCase(&usecase.Deps{
		BlackListRepository: l.mockBlackListRepository,
		UserRepository:      l.mockUserRepository,
		JwtIssuer:           l.mockJwtIssuer,
		JwtParser:           l.mockJwtParser,
		GAuthClient:         l.mockGAuthClient,
	})
}

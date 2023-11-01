package usecase

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lucky-pocket/luckyPocket-back/internal/app/auth/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	myjwt "github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/lucky-pocket/luckyPocket-back/test/stubs"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	uc                      domain.AuthUseCase
	mockBlackListRepository *stubs.BlackListRepository
	mockUserRepository      *stubs.UserRepository
	mockJwtIssuer           *mocks.Issuer
	mockJwtParser           *mocks.Parser
	mockGAuthClient         *mocks.GAuthClient
}

func TestAuthUseCaseSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (l *AuthUseCaseTestSuite) SetupSuite() {
	l.mockBlackListRepository = stubs.NewBlackListRepository(l.T())
	l.mockUserRepository = stubs.NewUserRepository(l.T())
	l.mockJwtIssuer = mocks.NewIssuer(l.T())
	l.mockJwtParser = mocks.NewParser(l.T())
	l.mockGAuthClient = mocks.NewGAuthClient(l.T())

	l.uc = usecase.NewAuthUseCase(&usecase.Deps{
		BlackListRepository: l.mockBlackListRepository,
		UserRepository:      l.mockUserRepository,
		JwtIssuer:           l.mockJwtIssuer,
		JwtParser:           l.mockJwtParser,
		GAuthClient:         l.mockGAuthClient,
	})
}

func generateTestToken() *myjwt.Token {
	return &myjwt.Token{
		Info: auth.Info{},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
}

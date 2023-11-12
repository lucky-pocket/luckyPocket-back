package usecase

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/app/game/usecase"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/test/stubs"
	"github.com/stretchr/testify/suite"
)

type GameUseCaseTestSuite struct {
	suite.Suite
	uc                    domain.GameUseCase
	mockGameLogRepository *stubs.GameLogRepository
	mockTicketRepository  *stubs.TicketRepository
	mockUserRepository    *stubs.UserRepository
}

func TestGameUseCaseSuite(t *testing.T) {
	suite.Run(t, new(GameUseCaseTestSuite))
}

func (g *GameUseCaseTestSuite) SetupSuite() {
	g.mockGameLogRepository = stubs.NewGameLogRepository(g.T())
	g.mockTicketRepository = stubs.NewTicketRepository(g.T())
	g.mockUserRepository = stubs.NewUserRepository(g.T())

	g.uc = usecase.NewGameUseCase(&usecase.Deps{
		GameLogRepository: g.mockGameLogRepository,
		TicketRepository:  g.mockTicketRepository,
		UserRepository:    g.mockUserRepository,
		TxManager:         tx.NewTxManager(),
	})
}

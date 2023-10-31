package usecase

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/mock"
)

func (g *GameUseCaseTestSuite) TestGetTicketInfo() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.FreeInput
		on     func()
		assert func(output *output.TicketOutput, err error)
	}{
		{
			desc:  "success (free ticket)",
			input: &input.FreeInput{Free: true},
			on: func() {
				g.mockTicketRepository.On("CountByUserID", mock.Anything, mock.Anything).Return(0, nil).Once()
				g.mockTicketRepository.On("GetRefillAt", mock.Anything, mock.Anything).Return(time.Now(), nil).Once()
			},
			assert: func(output *output.TicketOutput, err error) {
				if g.Nil(err) {
					g.NotNil(output.TicketCount)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		g.Run(tc.desc, func() {
			tc.on()

			tickets, err := g.uc.GetTicketInfo(ctx)

			tc.assert(tickets, err)

			g.mockUserRepository.AssertExpectations(g.T())
			g.mockGameLogRepository.AssertExpectations(g.T())
			g.mockTicketRepository.AssertExpectations(g.T())
		})
	}
}

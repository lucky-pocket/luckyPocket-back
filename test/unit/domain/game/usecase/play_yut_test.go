package usecase

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (g *GameUseCaseTestSuite) TestYutPlay() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.FreeInput
		on     func()
		assert func(output *output.YutOutput, err error)
	}{
		{
			desc:  "success (free ticket)",
			input: &input.FreeInput{Free: true},
			on: func() {
				g.mockTicketRepository.On("CountByUserID", mock.Anything, mock.Anything).Return(0, nil).Once()
				g.mockTicketRepository.On("Increase", mock.Anything, mock.Anything).Return(nil).Once()
				g.mockGameLogRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(output *output.YutOutput, err error) {
				if g.Nil(err) {
					g.NotZero(output.CoinsEarned)
				}
			},
		},

		{
			desc:  "success (use coin)",
			input: &input.FreeInput{Free: false},
			on: func() {
				g.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(10, nil).Once()
				g.mockUserRepository.On("UpdateCoin", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
				g.mockGameLogRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(output *output.YutOutput, err error) {
				if g.Nil(err) {
					g.NotZero(output.CoinsEarned)
				}
			},
		},

		{
			desc:  "failed (no have free ticket)",
			input: &input.FreeInput{Free: true},
			on: func() {
				g.mockTicketRepository.On("CountByUserID", mock.Anything, mock.Anything).Return(1, nil).Once()
			},
			assert: func(output *output.YutOutput, err error) {
				e, ok := err.(*status.Err)
				if g.True(ok) {
					g.Equal(http.StatusForbidden, e.Code)
				}
			},
		},

		{
			desc:  "failed (no coins)",
			input: &input.FreeInput{Free: false},
			on: func() {
				g.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(1, nil).Once()
			},
			assert: func(output *output.YutOutput, err error) {
				e, ok := err.(*status.Err)
				if g.True(ok) {
					g.Equal(http.StatusForbidden, e.Code)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		g.Run(tc.desc, func() {
			tc.on()

			yuts, err := g.uc.PlayYut(ctx, tc.input)

			tc.assert(yuts, err)

			g.mockUserRepository.AssertExpectations(g.T())
			g.mockGameLogRepository.AssertExpectations(g.T())
			g.mockTicketRepository.AssertExpectations(g.T())
		})
	}
}

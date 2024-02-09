package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/mock"
)

func (g *GameUseCaseTestSuite) TestGetPlayCount() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		on     func()
		assert func(output *output.PlayCountOutput, err error)
	}{
		{
			desc: "success",
			on: func() {
				g.mockGameLogRepository.On("CountByUserID", mock.Anything, mock.Anything).Return(0, nil).Once()
			},
			assert: func(output *output.PlayCountOutput, err error) {
				if g.Nil(err) {
					g.NotNil(output.Count)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		g.Run(tc.desc, func() {
			tc.on()

			count, err := g.uc.CountPlays(ctx)

			tc.assert(count, err)

			g.mockUserRepository.AssertExpectations(g.T())
			g.mockGameLogRepository.AssertExpectations(g.T())
			g.mockTicketRepository.AssertExpectations(g.T())
		})
	}
}

package usecase_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/mock"
)

func (s *UserUseCaseTestSuite) TestCountCoins() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		on     func()
		assert func(output *output.CoinOutput, err error)
	}{
		{
			desc: "success",
			on: func() {
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(3, nil)
			},
			assert: func(output *output.CoinOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToCoinOutput(3), output)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.CountCoins(ctx)

			tc.assert(output, err)

			s.mockUserRepository.AssertExpectations(s.T())
		})
	}
}

package usecase_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/assert"
)

func (s *PocketUseCaseTestSuite) TestSendPocket() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.PocketInput
		on     func()
		assert func(err error)
	}{
		{
			desc:  "success",
			input: &input.PocketInput{},
			on: func() {
				s.mockUserRepository.On("FindByID", nil, nil).Return(nil, nil).Once()
			},
			assert: func(err error) {
				assert.Error(s.T(), err)
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			err := s.uc.SendPocket(ctx, tc.input)

			tc.assert(err)

			s.mockPocketRepository.AssertExpectations(s.T())
			s.mockUserRepository.AssertExpectations(s.T())
			s.mockTxManager.AssertExpectations(s.T())
		})
	}
}

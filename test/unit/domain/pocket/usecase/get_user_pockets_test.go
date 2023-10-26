package usecase_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/stretchr/testify/mock"
)

func (s *PocketUseCaseTestSuite) TestGetUserPockets() {
	pockets := []*domain.Pocket{
		{
			PocketID: 123,
			Receiver: &domain.User{
				UserID: 1,
			},
		},
		{
			PocketID: 1531,
			Receiver: &domain.User{
				UserID: 1 + 1,
			},
		},
	}

	testcases := []struct {
		desc   string
		input  *input.PocketQueryInput
		on     func()
		assert func(output *output.PocketListOutput, err error)
	}{
		{
			desc:  "success",
			input: &input.PocketQueryInput{},
			on: func() {
				s.mockPocketRepository.On("FindListByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(pockets, nil).Once()
			},
			assert: func(output *output.PocketListOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToPocketListOutput(pockets), output)
				}
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.GetUserPockets(context.Background(), tc.input)

			tc.assert(output, err)

			s.mockPocketRepository.AssertExpectations(s.T())
			s.mockUserRepository.AssertExpectations(s.T())
		})
	}
}

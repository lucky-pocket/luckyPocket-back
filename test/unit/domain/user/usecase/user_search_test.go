package usecase_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *UserUseCaseTestSuite) TestSearch() {
	outputs := []*domain.User{
		{
			UserID: 1,
			Name:   "김",
		},
		{
			UserID: 2,
			Name:   "김",
		},
	}

	testcases := []struct {
		desc   string
		input  *input.SearchInput
		on     func()
		assert func(output *output.SearchOutput, err error)
	}{
		{
			desc:  "success",
			input: &input.SearchInput{Query: "김"},
			on: func() {
				s.mockUserRepository.On("FindByNameContains", mock.Anything, mock.Anything).Return(outputs, nil).Once()
			},
			assert: func(output *output.SearchOutput, err error) {
				if s.Nil(err) {
					s.Len(output.Users, 2)
				}
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.Search(context.Background(), tc.input)

			tc.assert(output, err)

			s.mockUserRepository.AssertExpectations(s.T())
		})
	}
}

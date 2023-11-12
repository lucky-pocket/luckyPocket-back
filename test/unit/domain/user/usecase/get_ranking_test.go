package usecase_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/stretchr/testify/mock"
)

func (s *UserUseCaseTestSuite) TestGetRanking() {
	users := []output.RankElem{
		{
			UserInfo: output.UserInfo{
				UserID: 1,
				Name:   "hi",
			},
			Gender: constant.GenderFemale,
			Amount: 0,
		},
	}

	userStudent := constant.UserType("STUDENT")
	userTeacher := constant.UserType("TEACHER")

	testcases := []struct {
		desc   string
		input  *input.RankQueryInput
		on     func()
		assert func(output *output.RankOutput, err error)
	}{
		{
			desc:  "success (student)",
			input: &input.RankQueryInput{UserType: &userStudent},
			on: func() {
				s.mockUserRepository.On("RankStudents",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(users, nil).Once()
			},
			assert: func(output *output.RankOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToRankOutput(users), output)
				}
			},
		},
		{
			desc:  "success (non-student)",
			input: &input.RankQueryInput{UserType: &userTeacher},
			on: func() {
				s.mockUserRepository.On("RankNonStudents",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(users, nil).Once()
			},
			assert: func(output *output.RankOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToRankOutput(users), output)
				}
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.GetRanking(context.Background(), tc.input)

			tc.assert(output, err)

			s.mockUserRepository.AssertExpectations(s.T())
			s.mockNoticeRepository.AssertExpectations(s.T())
		})
	}
}

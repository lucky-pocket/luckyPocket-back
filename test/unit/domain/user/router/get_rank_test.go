package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *UserRouterTestSuite) TestGetRank() {
	testcases := []struct {
		desc       string
		query      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success (POCKET-STUDENT)",
			query: "sortType=POCKET&userType=STUDENT&grade=1&class=1&name=HI?",
			on: func() {
				s.mockUserUseCase.On("GetRanking", mock.Anything, mock.Anything).Return(&output.RankOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "success (COIN-STUDENT)",
			query: "sortType=COIN&userType=STUDENT&grade=1&class=1&name=HI?",
			on: func() {
				s.mockUserUseCase.On("GetRanking", mock.Anything, mock.Anything).Return(&output.RankOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "success (POCKET-TEACHER)",
			query: "sortType=POCKET&userType=TEACHER&name=HI?",
			on: func() {
				s.mockUserUseCase.On("GetRanking", mock.Anything, mock.Anything).Return(&output.RankOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "success (COIN-TEACHER)",
			query: "sortType=COIN&userType=TEACHER&name=HI?",
			on: func() {
				s.mockUserUseCase.On("GetRanking", mock.Anything, mock.Anything).Return(&output.RankOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:       "failed (invalid query)",
			query:      "sor=COIN",
			on:         func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			desc:       "failed (non-enum)",
			query:      "sortType=COINA",
			on:         func() {},
			statusCode: http.StatusBadRequest,
		},
		{
			desc:  "ok without userType",
			query: "sortType=COIN",
			on: func() {
				s.mockUserUseCase.On("GetRanking", mock.Anything, mock.Anything).Return(&output.RankOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/rank?"+tc.query, nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockUserUseCase.AssertExpectations(s.T())
		})
	}
}

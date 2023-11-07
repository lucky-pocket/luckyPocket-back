package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *UserRouterTestSuite) TestSearch() {
	testcases := []struct {
		desc       string
		query      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success (NO query)",
			query: "",
			on: func() {
				s.mockUserUseCase.On("Search", mock.Anything, mock.Anything).Return(&output.SearchOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "success (In query)",
			query: "query=HI?",
			on: func() {
				s.mockUserUseCase.On("Search", mock.Anything, mock.Anything).Return(&output.SearchOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users?"+tc.query, nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockUserUseCase.AssertExpectations(s.T())
		})
	}
}

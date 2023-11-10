package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *AuthRouterTestSuite) TestLogin() {
	testcases := []struct {
		desc       string
		query      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success",
			query: "hi?",
			on: func() {
				s.mockAuthUseCase.On("Login", mock.Anything, mock.Anything).Return(&output.TokenOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "failed (BAD-REQUEST)",
			query: "",
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/auth/gauth?code="+tc.query, nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockAuthUseCase.AssertExpectations(s.T())
		})
	}
}

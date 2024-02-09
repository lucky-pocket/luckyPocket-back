package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *GameRouterTestSuite) TestGetPlayCount() {
	testcases := []struct {
		desc       string
		on         func()
		statusCode int
	}{
		{
			desc: "success",
			on: func() {
				s.mockGameUseCase.On("CountPlays", mock.Anything).Return(&output.PlayCountOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/games/count", nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)

			s.mockGameUseCase.AssertExpectations(s.T())
		})
	}
}

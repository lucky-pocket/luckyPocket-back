package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (p *PocketRouterTestSuite) TestGetUserPockets() {
	testcases := []struct {
		desc       string
		query      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success",
			query: "?offset=12&limit=12",
			on: func() {
				p.mockPocketUseCase.On("GetUserPockets", mock.Anything, mock.Anything).Return(&output.PocketListOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "failed (BAD-REQUEST)",
			query: "?offset=hi?&limit=12",
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		p.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/1/pockets"+tc.query, nil)
			p.engine.ServeHTTP(w, req)

			p.Equal(tc.statusCode, w.Code, req)
			p.mockPocketUseCase.AssertExpectations(p.T())
		})
	}
}

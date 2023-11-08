package router_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/mock"
)

func (p *PocketRouterTestSuite) TestGetMyPockets() {
	testcases := []struct {
		desc       string
		query      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success",
			query: "offset=12&limit=12",
			on: func() {
				p.mockPocketUseCase.On("GetUserPockets", mock.Anything, mock.Anything).Return(&output.PocketListOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "failed (bad-request)",
			query: "offset=zz&limit=zz",
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		p.Run(tc.desc, func() {
			tc.on()

			user := auth.Info{
				UserID: 1,
				Role:   constant.RoleMember,
			}

			ctx := auth.Inject(context.Background(), user)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/me/pockets/received?"+tc.query, nil)
			req = req.WithContext(ctx)
			p.engine.ServeHTTP(w, req)

			p.Equal(tc.statusCode, w.Code, req)
			p.mockPocketUseCase.AssertExpectations(p.T())
		})
	}
}

package router_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
)

func (p *PocketRouterTestSuite) TestSetVisibility() {
	testcases := []struct {
		desc       string
		body       []byte
		on         func()
		statusCode int
	}{
		{
			desc: "success",
			body: []byte(`
			{
				"visible" : true
            }
			`),
			on: func() {
				p.mockPocketUseCase.On("SetVisibility", mock.Anything, mock.Anything).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc: "success",
			body: []byte(`
			{
				"visible" : tr
            }
			`),
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		p.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PATCH", "/users/me/pockets/1/visibility", bytes.NewBuffer(tc.body))
			p.engine.ServeHTTP(w, req)

			p.Equal(tc.statusCode, w.Code, req)
			p.mockPocketUseCase.AssertExpectations(p.T())
		})
	}
}

package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
)

func (s *NoticeRouterTestSuite) TestCheckNotice() {
	testcases := []struct {
		desc       string
		param      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success",
			param: "1",
			on: func() {
				s.mockNoticeUseCase.On("CheckNotice", mock.Anything, mock.Anything).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "invalid param",
			param: "hithere",
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PATCH", "/users/me/notices/"+tc.param, nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockNoticeUseCase.AssertExpectations(s.T())
		})
	}
}

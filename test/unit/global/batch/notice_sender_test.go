package batch_test

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/batch"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestNoticeSender(t *testing.T) {
	mockNoticeRepository := mocks.NewNoticeRepository(t)
	mockNoticePool := mocks.NewNoticePool(t)

	n := batch.NewNoticeSender(&batch.NoticeSenderDeps{
		NoticeRepository: mockNoticeRepository,
		NoticePool:       mockNoticePool,
	})

	mockNoticePool.On("TakeAll", mock.Anything).Return([]*domain.Notice{}, nil)
	mockNoticeRepository.On("CreateBulk", mock.Anything, mock.Anything).Return(nil)

	// TODO: add test for log when it is added.
	n.Do()

	mockNoticePool.AssertExpectations(t)
	mockNoticeRepository.AssertExpectations(t)
}

package batch_test

import (
	"os"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/batch"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNoticeSender(t *testing.T) {
	mockNoticeRepository := mocks.NewNoticeRepository(t)
	mockNoticePool := mocks.NewNoticePool(t)

	n := batch.NewNoticeSender(&batch.NoticeSenderDeps{
		NoticeRepository: mockNoticeRepository,
		NoticePool:       mockNoticePool,
		Logger: zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			os.Stdout,
			zap.DebugLevel,
		)),
	})

	mockNoticePool.On("Take", mock.Anything, mock.Anything).Return([]*domain.Notice{{}}, nil)
	mockNoticeRepository.On("CreateBulk", mock.Anything, mock.Anything).Return(nil)

	n.Do()

	mockNoticePool.AssertExpectations(t)
	mockNoticeRepository.AssertExpectations(t)
}

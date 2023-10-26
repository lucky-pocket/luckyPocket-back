package usecase

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
)

type Deps struct {
	NoticeRepository domain.NoticeRepository
}

type noticeUseCase struct{ *Deps }

func NewNoticeUseCase(deps *Deps) domain.NoticeUseCase {
	return &noticeUseCase{deps}
}

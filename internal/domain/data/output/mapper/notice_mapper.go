package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToNoticeListOutPut(noticeList []*domain.Notice) *output.NoticeListOutput {
	notices := &output.NoticeListOutput{}

	for _, n := range noticeList {
		notices.Notices = append(notices.Notices, output.NoticeElem{
			NoticeID:  n.NoticeID,
			Kind:      n.Type,
			PocketID:  n.Pocket.PocketID,
			Checked:   n.Checked,
			CreatedAt: n.CreatedAt,
		})
	}

	return notices
}

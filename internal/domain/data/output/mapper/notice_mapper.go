package mapper

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
)

func ToNoticeListOutPut(notices []*domain.Notice) *output.NoticeListOutput {
	out := &output.NoticeListOutput{
		Notices: make([]output.NoticeElem, 0, len(notices)),
	}

	for _, notice := range notices {
		out.Notices = append(out.Notices, output.NoticeElem{
			NoticeID:  notice.NoticeID,
			Kind:      notice.Type,
			PocketID:  notice.Pocket.PocketID,
			Checked:   notice.Checked,
			CreatedAt: notice.CreatedAt,
		})
	}
	return out
}

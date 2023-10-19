package output

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"time"
)

type NoticeElem struct {
	NoticeID  uint64              `json:"id"`
	Kind      constant.NoticeType `json:"kind"`
	PocketID  uint64              `json:"pocketId"`
	Checked   bool                `json:"checked"`
	CreatedAt time.Time           `json:"createdAt"`
}

type NoticeListOutput struct {
	Notices []NoticeElem `json:"notices"`
}

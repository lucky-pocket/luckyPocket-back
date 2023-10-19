package domain

import "github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"

type Notice struct {
	NoticeID uint64
	User     *User
	Pocket   *Pocket
	Type     constant.NoticeType
	Checked  bool
}

package constant

import "time"

const (
	JwtSigningMethod = "HS256"
	JwtAccessTTL     = 5 * time.Minute
	JwtRefreshTTL    = 10 * 24 * time.Hour
)

const (
	CostRevealSender = 2
	CostSendPocket   = 5
	CostPlayYut      = 2

	PrizeDo     = 2
	PrizeGae    = 2
	PrizeGeol   = 3
	PrizeYut    = 4
	PrizeMo     = 6
	PrizeBackDo = -2
)

const (
	PlayGameLimit = 30
	TicketLimit   = 1
)

const (
	LimitSendSame = 5
)

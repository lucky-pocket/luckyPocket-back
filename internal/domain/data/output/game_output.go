package output

import "time"

type YutResultElem struct {
	Marked    bool    `json:"marked"`
	YutPieces [3]bool `json:"yutPieces"`
}

type YutOutput struct {
	Result      YutResultElem `json:"result"`
	CoinsEarned int           `json:"coinsEarned"`
}

type TicketOutput struct {
	RefillAt    time.Time `json:"refillAt"`
	TicketCount int       `json:"ticketCount"`
}

package domain

type Pocket struct {
	PocketID   uint64
	ReceiverID *User
	SenderID   *User
	Content    string
	Coins      int
}

func (p Pocket) IsEmpty() bool {
	return p.Coins == 0
}

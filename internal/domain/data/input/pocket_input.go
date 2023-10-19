package input

type PocketInput struct {
	Receiver string
	Coins    int
	Message  string
}

type UserIDInput struct {
	UserID uint64
}

type VisibilityInput struct {
	PocketID uint64
	Visible  bool
}

type PocketIDInput struct {
	PocketID uint64
}

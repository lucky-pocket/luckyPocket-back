package input

type PocketInput struct {
	ReceiverID uint64
	Coins      int
	Message    string
	IsPublic   bool
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

type PocketQueryInput struct {
	UserID uint64
	Offset int
	Limit  int
}

package output

type PocketOutput struct {
	Content string    `json:"content"`
	Coins   int       `json:"coins"`
	Sender  *UserInfo `json:"sender"`
}

type PocketElem struct {
	PocketID   uint64  `json:"id"`
	IsEmpty    bool    `json:"isEmpty"`
	IsPublic   bool    `json:"isPublic"`
	SenderName *string `json:"sender"`
}

type PocketListOutput struct {
	Pockets []PocketElem `json:"pockets"`
}

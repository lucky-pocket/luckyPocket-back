package output

import "time"

type TokenElem struct {
	Token     string
	ExpiresAt time.Time
}

type TokenOutput struct {
	Access  TokenElem
	Refresh TokenElem
}

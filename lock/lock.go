package lock

import "time"

type Lock struct {
	Key       string
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

package xtime

import "time"

// XTimer for mocking time
type XTimer interface {
	Now() time.Time
}

type xtime struct{}

// New create xtimer
func New() XTimer {
	return &xtime{}
}

func (t *xtime) Now() time.Time {
	return time.Now()
}

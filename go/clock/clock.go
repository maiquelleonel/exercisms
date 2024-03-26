package clock

import (
	"fmt"
)

const (
	MINS_IN_DAY = 1440
)

// Define the Clock type here.
type Clock int

func New(h, m int) Clock {
	minutes := (h * 60) + m

	for minutes >= MINS_IN_DAY {
		minutes -= MINS_IN_DAY
	}

	for minutes < 0 {
		minutes += MINS_IN_DAY
	}

	return Clock(minutes)
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

func (c Clock) String() string {

	return fmt.Sprintf("%02d:%02d", int(c)/60, int(c)%60)
}

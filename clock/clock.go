package clock

var clock uint64 = 0

type Protocol interface {
	GetClock()
	CompareClock(c Clock) bool
}

type Clock struct {
	ServerId string
	Number   uint64
}

// Get clock with increase number
func (c *Clock) GetClock() {
	clock++
	c.Number = clock
}

// Compare clock
func (c *Clock) CompareClock(c1 Clock) bool {
	if (c.ServerId == c1.ServerId && c.Number > c1.Number) || c.ServerId > c1.ServerId {
		return true
	} else {
		return false
	}
}

package clock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClock_GetClock(t *testing.T) {
	clock := Clock{
		ServerId: "1234",
	}
	clock.GetClock()
	require.Equal(t, uint64(1), clock.Number)
	clock.GetClock()
	require.Equal(t, uint64(2), clock.Number)
	clock.GetClock()
	require.Equal(t, uint64(3), clock.Number)
}

func TestClock_CompareClock(t *testing.T) {
	clock1 := Clock{
		ServerId: "12345",
	}
	clock1.GetClock()
	clock2 := Clock{
		ServerId: "77777",
	}
	clock2.GetClock()

	res := clock1.CompareClock(clock2)
	require.Equal(t, false, res)
}

func TestClock_CompareClockWithSameServer(t *testing.T) {
	clock1 := Clock{
		ServerId: "12345",
	}
	clock2 := Clock{
		ServerId: "12345",
	}
	clock2.GetClock()

	res := clock2.CompareClock(clock1)
	require.Equal(t, true, res)
}

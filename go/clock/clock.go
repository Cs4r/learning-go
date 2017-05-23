package clock

import "fmt"

const testVersion = 4

// You can find more details and hints in the test file.

type Clock int64 // Complete the type definition.  Pick a suitable data type.

func New(hour, minute int) Clock {
	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock(time)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}

func (clock Clock) Add(minutes int) Clock {
	return New(0, int(clock)+minutes)
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.

//go:build test

package shared

import "time"

func UseFakeClockAt(t time.Time) *FakeClock {
	fc := &FakeClock{now: t}
	clock = fc
	return fc
}

func UseSystemClock() {
	clock = systemClock{}
}

// Fake clock

type FakeClock struct {
	now time.Time
}

func (f *FakeClock) Now() time.Time {
	return f.now
}

func (f *FakeClock) Advance(d time.Duration) {
	f.now = f.now.Add(d)
}

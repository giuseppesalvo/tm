package tm

import "time"

type Interval struct {
	Ticker  *time.Ticker
	Running bool
}

func SetInterval(callback func(), delay time.Duration) *Interval {

	interval := &Interval{
		Ticker:  time.NewTicker(time.Millisecond * delay),
		Running: true,
	}

	go func() {
		for range interval.Ticker.C {
			callback()
		}
	}()

	return interval
}

func ClearInterval(interval *Interval) {
	if interval != nil && interval.Running {
		interval.Ticker.Stop()
		interval.Running = false
	}
}

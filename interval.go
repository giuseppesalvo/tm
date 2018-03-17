package tm

import "time"

type Interval struct {
	Ticker  *time.Ticker
	Running bool
}

func setInterval(callback func(), delay time.Duration) *Interval {

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

func clearInterval(interval *Interval) {
	if interval.Running {
		interval.Ticker.Stop()
		interval.Running = false
	}
}

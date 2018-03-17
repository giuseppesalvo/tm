package tm

import "time"

type Timeout struct {
	Timer *time.Timer
	Fired bool
}

func setTimeout(callback func(), delay time.Duration) *Timeout {
	timeout := &Timeout{
		Timer: time.NewTimer(time.Millisecond * delay),
		Fired: false,
	}

	go func() {
		<- timeout.Timer.C
		callback()
		timeout.Fired = true
	}()

	return timeout
}

func clearTimeout( timeout *Timeout ) {
	if ! timeout.Fired {
		timeout.Timer.Stop()
	}
}
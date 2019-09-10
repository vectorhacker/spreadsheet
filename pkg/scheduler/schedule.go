package scheduler

import "time"

// Schedule implements a deffered function call.
func Schedule(f func()) {
	time.AfterFunc(0, f)
}

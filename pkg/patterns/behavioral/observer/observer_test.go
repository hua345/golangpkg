package observer

import (
	"testing"
	"time"
)

/**
https://github.com/tmrts/go-patterns/blob/master/behavioral/observer.md
The observer pattern allows a type instance to "publish" events to other type instances ("observers")
who wish to be updated when a particular event occurs.
*/
func TestObserver(t *testing.T) {
	// Initialize a new Notifier.
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers.
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(1 * time.Second).C
	tick := time.NewTicker(100 * time.Millisecond).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}

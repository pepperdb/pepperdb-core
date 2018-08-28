package common

import (
	"time"
)

// NonBlockingChan will drop new values when full
type NonBlockingChan struct {
	innerChan chan interface{}
}

// NewNonBlockingChan create a new non-blocking chan
func NewNonBlockingChan(cap int) *NonBlockingChan {
	return &NonBlockingChan{
		innerChan: make(chan interface{}, cap),
	}
}

// Send value into chan
func (nbCh *NonBlockingChan) Send(value interface{}) bool {
	select {
	case nbCh.innerChan <- value:
		return true
	default:
		return false
	}
}

// Recv value from chan
func (nbCh *NonBlockingChan) Recv() (interface{}, bool) {
	select {
	case value := <-nbCh.innerChan:
		return value, true
	default:
		return nil, false
	}
}

// SendWithDeadline try to send value in given duration, the value will be dropped if failed
func (nbCh *NonBlockingChan) SendWithDeadline(value interface{}, deadline time.Duration) bool {
	if deadline == 0 {
		return nbCh.Send(value)
	}

	select {
	case nbCh.innerChan <- value:
		return true
	case <-time.NewTicker(deadline).C:
		return false
	}
}

// RecvWithDeadline try to recv value in given duration, the value will be skipped if failed
func (nbCh *NonBlockingChan) RecvWithDeadline(deadline time.Duration) (interface{}, bool) {
	if deadline == 0 {
		return nbCh.Recv()
	}

	select {
	case value := <-nbCh.innerChan:
		return value, true
	case <-time.NewTicker(deadline).C:
		return nil, false
	}
}

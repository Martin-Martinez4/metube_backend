package testhelpers

import (
	"context"
	"time"
)

type MyContext struct{}

func (ctx MyContext) Deadline() (deadline time.Time, ok bool) {
	// return deadline, ok
	return
}

func (ctx MyContext) Done() <-chan struct{} {
	// ch := make(chan struct{})
	// close(ch)
	// return ch
	return nil
}

func (ctx MyContext) Err() error {
	return context.DeadlineExceeded
}

func (ctx MyContext) Value(key interface{}) interface{} {
	return nil
}

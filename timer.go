package utils

import (
	"time"
)

type Timer struct {
	Counter  int
	Canceled bool
	Ended    bool
}

/// Cancel 取消任务
func (t *Timer) Cancel() *Timer {
	t.Canceled = true
	return t
}

/// SetTimeout 延时执行任务
func SetTimeout(h func(), d time.Duration) *Timer {
	t := &Timer{}
	go func() {
		time.Sleep(d)
		if !t.Canceled {
			h()
			t.Counter++
		}
		t.Ended = true
	}()
	return t
}

/// SetInterval 循环某个周期执行任务
func SetInterval(h func(), d time.Duration) *Timer {
	t := &Timer{}
	go func() {
		for !t.Canceled {
			time.Sleep(d)
			h()
			t.Counter++
		}
		t.Ended = true
	}()
	return t
}

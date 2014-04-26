package main

import "time"

var now = time.Now

type timeImpl interface {
	Now() time.Time
}

type systemTime struct{}

func (systemTime) Now() time.Time        { return time.Now() }
func (systemTime) Sleep(d time.Duration) { time.Sleep(d) }

type fakeTime time.Time

func (t *fakeTime) Now() time.Time        { return time.Time(*t) }
func (t *fakeTime) Sleep(d time.Duration) { *t = fakeTime(time.Time(*t).Add(d)) }

func Sleeper() {

}

package service

import (
	"testing"
	"time"
)

func TestSchedulerAgainIsSoon(t *testing.T) {
	now := time.Now()
	due := ReviewScheduler{}.NextDue("again", now)
	if due.After(now.Add(time.Hour)) {
		t.Fatal("again should be soon")
	}
}

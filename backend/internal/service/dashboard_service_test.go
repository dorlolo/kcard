package service

import "testing"

func TestDashboardHasNextAction(t *testing.T) {
	s := DashboardService{}.Summary()
	if len(s.NextActions) == 0 {
		t.Fatal("missing next action")
	}
}

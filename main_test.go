package main

import "testing"

func TestPing(t *testing.T) {
	got := ping()
	if got != "test"{
		t.Errorf("ping() =%v; want test", got)
	}
}

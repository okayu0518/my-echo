package main

import "testing"

func TestPing(t *testing.T) {
	got := ping()
	if got != "pongpong" {
		t.Errorf("ping() =%v; want test", got)
	}
}

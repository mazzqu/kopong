package main

import "testing"

func TestGreeting(t *testing.T) {
	name := "lord jengkloz"
	want := "hello lord jengkloz"

	if got := greeting(name); got != want {
		t.Errorf("greeting(%q) = %q, want %q", name, got, want)
	}
}

package main

import "testing"

func TestGreeting(t *testing.T) {
	name := "alize"
	want := "hello alize"

	if got := greeting(name); got != want {
		t.Errorf("greeting(%q) = %q, want %q", name, got, want)
	}
}

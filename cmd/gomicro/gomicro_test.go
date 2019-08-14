package main

import "testing"

func TestMain(t *testing.T) {
	want := "hello"
	if got := main(); got != want {
		t.Errorf("main() = %q, want %q", got, want)
	}
}
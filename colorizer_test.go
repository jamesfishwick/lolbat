// colorizer_test.go
package main

import (
	"strings"
	"testing"
)

func TestSequential_SingleChar(t *testing.T) {
	opts := Options{Mode: "sequential", Spread: 1.0, Seed: 0}
	got := ColorizeLine("A", 0, opts)
	want := ColorAt(0).ANSI() + "A" + Reset
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSequential_SpreadTwo(t *testing.T) {
	opts := Options{Mode: "sequential", Spread: 2.0, Seed: 0}
	got := ColorizeLine("AB", 0, opts)
	// Both chars at spread=2 should use same color index 0
	wantColor := ColorAt(0).ANSI()
	if !strings.Contains(got, wantColor+"A") {
		t.Errorf("expected A to use color 0, got %q", got)
	}
	if !strings.Contains(got, wantColor+"B") {
		t.Errorf("expected B to use color 0 at spread=2, got %q", got)
	}
}

func TestSequential_SeedOffset(t *testing.T) {
	opts := Options{Mode: "sequential", Spread: 1.0, Seed: 3}
	got := ColorizeLine("A", 0, opts)
	want := ColorAt(3).ANSI() + "A" + Reset
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

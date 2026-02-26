// palette_test.go
package main

import (
	"testing"
)

func TestPaletteLength(t *testing.T) {
	if len(Palette) != 13 {
		t.Errorf("expected 13 colors, got %d", len(Palette))
	}
}

func TestColorAt_Wraps(t *testing.T) {
	c0 := ColorAt(0)
	c13 := ColorAt(13)
	if c0 != c13 {
		t.Errorf("expected palette to wrap: ColorAt(0)=%v, ColorAt(13)=%v", c0, c13)
	}
}

func TestColorAt_NegativeSeed(t *testing.T) {
	// Should not panic
	_ = ColorAt(-1)
}

func TestANSI(t *testing.T) {
	c := Color{R: 255, G: 85, B: 85}
	got := c.ANSI()
	want := "\033[38;2;255;85;85m"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

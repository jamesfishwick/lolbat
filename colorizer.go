// colorizer.go
package main

import (
	"math"
	"math/rand"
)

// Options holds the colorizer configuration parsed from CLI flags.
type Options struct {
	Mode    string
	Freq    float64
	Spread  float64
	Seed    int
	Animate bool
}

// ColorizeLine colorizes a single line of text according to the options.
// lineIndex is the line number (0-based), used for sine mode vertical offset.
func ColorizeLine(line string, lineIndex int, opts Options) string {
	switch opts.Mode {
	case "sine":
		return colorizeLineSine(line, lineIndex, opts)
	case "random":
		return colorizeLineRandom(line, lineIndex, opts)
	default:
		return colorizeLineSequential(line, opts)
	}
}

func colorizeLineSequential(line string, opts Options) string {
	out := ""
	for i, ch := range line {
		idx := opts.Seed + int(float64(i)/opts.Spread)
		out += ColorAt(idx).ANSI() + string(ch)
	}
	return out + Reset
}

func colorizeLineSine(line string, lineIndex int, opts Options) string {
	out := ""
	n := float64(len(Palette))
	for i, ch := range line {
		pos := float64(opts.Seed+i)/opts.Spread + float64(lineIndex)*opts.Freq
		// Map sine output (-1..1) to palette index (0..n-1)
		idx := int((math.Sin(pos*opts.Freq*2*math.Pi)+1)/2*(n-1)+0.5) % len(Palette)
		out += ColorAt(idx).ANSI() + string(ch)
	}
	return out + Reset
}

func colorizeLineRandom(line string, lineIndex int, opts Options) string {
	r := rand.New(rand.NewSource(int64(opts.Seed + lineIndex)))
	color := ColorAt(r.Intn(len(Palette)))
	return color.ANSI() + line + Reset
}

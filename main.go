// main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	mode := flag.String("mode", "sequential", "Coloring mode: sequential, sine, random")
	freq := flag.Float64("freq", 0.3, "Color wave frequency (sine mode)")
	spread := flag.Float64("spread", 3.0, "Characters per color step")
	seed := flag.Int("seed", 0, "Starting palette offset")
	animate := flag.Bool("a", false, "Animate colors")
	flag.Parse()

	opts := Options{
		Mode:    *mode,
		Freq:    *freq,
		Spread:  *spread,
		Seed:    *seed,
		Animate: *animate,
	}

	var lines []string
	args := flag.Args()
	if len(args) == 0 {
		lines = readLines(os.Stdin)
	} else {
		for _, path := range args {
			f, err := os.Open(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "lolbat: %v\n", err)
				os.Exit(1)
			}
			lines = append(lines, readLines(f)...)
			if err := f.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "lolbat: %v\n", err)
				os.Exit(1)
			}
		}
	}

	if opts.Animate {
		runAnimate(lines, opts)
	} else {
		printLines(lines, opts)
	}
}

func readLines(f *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func printLines(lines []string, opts Options) {
	for i, line := range lines {
		fmt.Println(ColorizeLine(line, i, opts))
	}
}

func runAnimate(lines []string, opts Options) {
	frame := 0
	for {
		// Clear screen
		fmt.Print("\033[H\033[2J")
		animOpts := opts
		animOpts.Seed = opts.Seed + frame
		printLines(lines, animOpts)
		time.Sleep(80 * time.Millisecond)
		frame++
	}
}

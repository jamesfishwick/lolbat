# lolbat Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Build a standalone Go CLI tool that colorizes stdin/file input using the Dracula ANSI color palette, with three coloring modes and the same flags as lolcat.

**Architecture:** Single Go module with three source files — `main.go` for entry point and flag parsing, `palette.go` for Dracula color definitions and indexing, `colorizer.go` for the three mode implementations. Input is read line-by-line from stdin or files, each character is wrapped in an ANSI true-color escape sequence, and output is written to stdout.

**Tech Stack:** Go 1.21+, stdlib only (no external dependencies), `go test` for testing.

---

## Palette Reference

13 colors from the official Dracula spec (AnsiBlack and AnsiWhite excluded for terminal compatibility):

```
AnsiRed           #FF5555   rgb(255, 85, 85)
AnsiBrightRed     #FF6E6E   rgb(255, 110, 110)
AnsiMagenta       #FF79C6   rgb(255, 121, 198)
AnsiBrightMagenta #FF92DF   rgb(255, 146, 223)
AnsiBlue          #BD93F9   rgb(189, 147, 249)
AnsiBrightBlue    #D6ACFF   rgb(214, 172, 255)
AnsiBrightBlack   #6272A4   rgb(98, 114, 164)
AnsiCyan          #8BE9FD   rgb(139, 233, 253)
AnsiBrightCyan    #A4FFFF   rgb(164, 255, 255)
AnsiGreen         #50FA7B   rgb(80, 250, 123)
AnsiBrightGreen   #69FF94   rgb(105, 255, 148)
AnsiYellow        #F1FA8C   rgb(241, 250, 140)
AnsiBrightYellow  #FFFFA5   rgb(255, 255, 165)
```

---

## Task 1: Initialize the Go module and repo

**Files:**
- Create: `/Users/jamesfishwick/Workspace/lolbat/go.mod`
- Create: `/Users/jamesfishwick/Workspace/lolbat/.gitignore`
- Create: `/Users/jamesfishwick/Workspace/lolbat/README.md`

**Step 1: Initialize git repo**

```bash
cd /Users/jamesfishwick/Workspace/lolbat
git init
```

**Step 2: Initialize Go module**

```bash
go mod init github.com/jamesfishwick/lolbat
```

Expected: `go.mod` created with `module github.com/jamesfishwick/lolbat` and Go version.

**Step 3: Create .gitignore**

```
lolbat
*.test
```

**Step 4: Create README.md**

```markdown
# lolbat

Colorize terminal output using the [Dracula](https://draculatheme.com) color palette.

A lolcat-compatible CLI tool with three coloring modes.

## Usage

    lolbat [options] [file ...]
    command | lolbat [options]

## Options

    --mode      sequential|sine|random (default: sequential)
    --freq      Color wave frequency, sine mode only (default: 0.3)
    --spread    Characters per color step (default: 3.0)
    --seed      Starting palette offset (default: 0)
    -a          Animate colors

## Install

    go install github.com/jamesfishwick/lolbat@latest
```

**Step 5: Commit**

```bash
git add .
git commit -m "chore: initialize go module and repo"
```

---

## Task 2: Implement palette.go

**Files:**
- Create: `/Users/jamesfishwick/Workspace/lolbat/palette.go`
- Create: `/Users/jamesfishwick/Workspace/lolbat/palette_test.go`

**Step 1: Write the failing tests**

```go
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
```

**Step 2: Run tests to verify they fail**

```bash
cd /Users/jamesfishwick/Workspace/lolbat
go test ./...
```

Expected: compilation failure — `Palette`, `ColorAt`, `Color`, `ANSI` not defined yet.

**Step 3: Implement palette.go**

```go
// palette.go
package main

import "fmt"

// Color holds an RGB triple.
type Color struct {
    R, G, B uint8
}

// ANSI returns the ANSI true-color foreground escape sequence for the color.
func (c Color) ANSI() string {
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

// Reset is the ANSI reset sequence.
const Reset = "\033[0m"

// Palette is the Dracula ANSI color palette (excluding black and white).
// Order: warm (reds/pinks) -> cool (purples/blues) -> muted -> cyans -> greens -> yellows
var Palette = []Color{
    {255, 85, 85},   // AnsiRed           #FF5555
    {255, 110, 110}, // AnsiBrightRed     #FF6E6E
    {255, 121, 198}, // AnsiMagenta       #FF79C6
    {255, 146, 223}, // AnsiBrightMagenta #FF92DF
    {189, 147, 249}, // AnsiBlue          #BD93F9
    {214, 172, 255}, // AnsiBrightBlue    #D6ACFF
    {98, 114, 164},  // AnsiBrightBlack   #6272A4
    {139, 233, 253}, // AnsiCyan          #8BE9FD
    {164, 255, 255}, // AnsiBrightCyan    #A4FFFF
    {80, 250, 123},  // AnsiGreen         #50FA7B
    {105, 255, 148}, // AnsiBrightGreen   #69FF94
    {241, 250, 140}, // AnsiYellow        #F1FA8C
    {255, 255, 165}, // AnsiBrightYellow  #FFFFA5
}

// ColorAt returns the palette color at the given index, wrapping around.
// Handles negative indices safely.
func ColorAt(i int) Color {
    n := len(Palette)
    return Palette[((i % n) + n) % n]
}
```

**Step 4: Run tests to verify they pass**

```bash
go test ./... -v
```

Expected: all 4 tests PASS.

**Step 5: Commit**

```bash
git add palette.go palette_test.go
git commit -m "feat: add Dracula palette with ANSI escape support"
```

---

## Task 3: Implement colorizer.go — sequential mode

**Files:**
- Create: `/Users/jamesfishwick/Workspace/lolbat/colorizer.go`
- Create: `/Users/jamesfishwick/Workspace/lolbat/colorizer_test.go`

**Step 1: Write failing tests**

```go
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
```

**Step 2: Run tests to verify they fail**

```bash
go test ./...
```

Expected: compilation failure — `Options`, `ColorizeLine` not defined.

**Step 3: Implement colorizer.go with sequential mode**

```go
// colorizer.go
package main

import (
    "math"
    "math/rand"
)

// Options holds the colorizer configuration parsed from CLI flags.
type Options struct {
    Mode   string
    Freq   float64
    Spread float64
    Seed   int
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
```

**Step 4: Run tests**

```bash
go test ./... -v
```

Expected: all tests PASS.

**Step 5: Commit**

```bash
git add colorizer.go colorizer_test.go
git commit -m "feat: add sequential, sine, and random colorizer modes"
```

---

## Task 4: Implement main.go — flag parsing and stdin/file reading

**Files:**
- Create: `/Users/jamesfishwick/Workspace/lolbat/main.go`

**Step 1: Implement main.go**

No unit tests for main.go — it's glue code. Manual smoke test after.

```go
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

    // Collect all lines from stdin or files
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
            f.Close()
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
```

**Step 2: Build and smoke test**

```bash
cd /Users/jamesfishwick/Workspace/lolbat
go build -o lolbat .
echo "Hello from lolbat" | ./lolbat
echo "Hello from lolbat" | ./lolbat --mode sine
echo "Hello from lolbat" | ./lolbat --mode random
```

Expected: colorized output in Dracula palette colors for each mode.

**Step 3: Commit**

```bash
git add main.go
git commit -m "feat: add main entry point with flag parsing and animate mode"
```

---

## Task 5: Create GitHub repo and push

**Step 1: Create repo on GitHub**

```bash
gh repo create lolbat --public --description "lolcat with the Dracula color palette" --source=. --remote=origin
```

**Step 2: Push**

```bash
git push -u origin main
```

**Step 3: Verify**

```bash
gh repo view lolbat --web
```

---

## Task 6: Final verification

**Step 1: Run full test suite**

```bash
go test ./... -v
```

Expected: all tests pass.

**Step 2: Build clean binary**

```bash
go build -o lolbat .
```

**Step 3: End-to-end smoke tests**

```bash
# Sequential (default)
echo "Dracula theme is awesome" | ./lolbat

# Sine mode
echo "Dracula theme is awesome" | ./lolbat --mode sine --freq 0.5

# Random mode
echo "Dracula theme is awesome" | ./lolbat --mode random

# Seed offset
echo "Dracula theme is awesome" | ./lolbat --seed 5

# Spread
echo "Dracula theme is awesome" | ./lolbat --spread 1.0

# File input
echo "test content" > /tmp/test.txt
./lolbat /tmp/test.txt

# Multi-line
printf "line one\nline two\nline three\n" | ./lolbat
```

**Step 4: Commit binary exclusion check**

```bash
git status
```

Expected: only source files tracked, `lolbat` binary ignored.

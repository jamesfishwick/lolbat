# lolbat: Dracula-themed terminal colorizer

*2026-02-26T14:11:19Z by Showboat 0.6.1*
<!-- showboat-id: 4a9f48e3-25b7-4f10-8b8f-335ada622b53 -->

lolbat colorizes terminal output using the Dracula color palette. It reads from stdin or files and supports three coloring modes: sequential, sine, and random. Flags are compatible with lolcat.

## Install

```bash
go install github.com/jamesfishwick/lolbat@latest && echo 'installed'
```

```output
installed
```

## Flags

```
-mode string    Coloring mode: sequential, sine, random (default: sequential)
-freq float     Color wave frequency, sine mode only (default: 0.3)
-spread float   Characters per color step (default: 3)
-seed int       Starting palette offset (default: 0)
-a              Animate colors (loops until Ctrl-C)
```

## Sequential mode (default)

Each character advances through the Dracula palette. `-spread` controls how many characters share a color before stepping to the next.

```bash
printf "Painful zombies quickly watch a jinxed graveyard.\n..." | lolbat -spread 1.0
```

![Sequential mode: -spread 1.0](docs/images/demo-sequential.svg)

## Sine mode

The palette index oscillates along a sine wave. `-freq` controls wave frequency; `-spread` controls horizontal compression.

```bash
printf "Painful zombies quickly watch a jinxed graveyard.\n..." | lolbat -mode sine -freq 0.5 -spread 2.0
```

![Sine mode: -mode sine -freq 0.5 -spread 2.0](docs/images/demo-sine.svg)

## Random mode

Each line gets one color, chosen by a seeded PRNG keyed on `-seed + line number`. Output is fully deterministic.

```bash
printf "Painful zombies quickly watch a jinxed graveyard.\n..." | lolbat -mode random
```

![Random mode: -mode random](docs/images/demo-random.svg)

## Mode comparison

Same pangrams run through all three modes.

```bash
printf "..." | lolbat -spread 2.0          # sequential
printf "..." | lolbat -mode sine -spread 2.0
printf "..." | lolbat -mode random
```

![Mode comparison: sequential / sine / random](docs/images/demo-modes.svg)

## Spread comparison

`-spread 1` changes color every character. `-spread 3` (default) groups ~3 chars per color. `-spread 6` produces broader bands.

```bash
printf "..." | lolbat -spread 1.0
printf "..." | lolbat -spread 3.0
printf "..." | lolbat -spread 6.0
```

![Spread comparison: -spread 1 / 3 / 6](docs/images/demo-spread.svg)

## Seed variations

`-seed` shifts the starting position in the palette. Each value produces a different color entry point.

```bash
printf "..." | lolbat -seed 0 -spread 2.0
printf "..." | lolbat -seed 3 -spread 2.0
printf "..." | lolbat -seed 7 -spread 2.0
printf "..." | lolbat -seed 10 -spread 2.0
```

![Seed variations: -seed 0 / 3 / 7 / 10](docs/images/demo-seed.svg)

## Freq comparison (sine mode)

Higher `-freq` compresses the sine wave, cycling through the palette faster.

```bash
printf "..." | lolbat -mode sine -freq 0.1 -spread 2.0
printf "..." | lolbat -mode sine -freq 0.3 -spread 2.0
printf "..." | lolbat -mode sine -freq 0.6 -spread 2.0
printf "..." | lolbat -mode sine -freq 1.0 -spread 2.0
```

![Freq comparison: -mode sine -freq 0.1 / 0.3 / 0.6 / 1.0](docs/images/demo-freq.svg)

## Animate mode

`-a` loops the output continuously, cycling colors on each pass. Interrupt with Ctrl-C.

```bash
printf "Painful zombies quickly watch a jinxed graveyard.\n..." | lolbat -a
```

![Animate mode: -a flag](docs/images/demo-animate.gif)

## Test suite

```bash
go test ./... -v
```

```output
=== RUN   TestSequential_SingleChar
--- PASS: TestSequential_SingleChar (0.00s)
=== RUN   TestSequential_SpreadTwo
--- PASS: TestSequential_SpreadTwo (0.00s)
=== RUN   TestSequential_SeedOffset
--- PASS: TestSequential_SeedOffset (0.00s)
=== RUN   TestPaletteLength
--- PASS: TestPaletteLength (0.00s)
=== RUN   TestColorAt_Wraps
--- PASS: TestColorAt_Wraps (0.00s)
=== RUN   TestColorAt_NegativeSeed
--- PASS: TestColorAt_NegativeSeed (0.00s)
=== RUN   TestANSI
--- PASS: TestANSI (0.00s)
PASS
ok   github.com/jamesfishwick/lolbat (cached)
```

## Source

- GitHub: <https://github.com/jamesfishwick/lolbat>
- Go 1.21+, stdlib only — no external dependencies.

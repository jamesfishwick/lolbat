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

```bash {image}
\![Sequential mode: -spread 1.0](demo-sequential.svg)
```

![Sequential mode: -spread 1.0](f2acbd18-2026-02-26.svg)

## Sine mode

The palette index oscillates along a sine wave. `-freq` controls wave frequency; `-spread` controls horizontal compression.

```bash
printf "Painful zombies quickly watch a jinxed graveyard.\n..." | lolbat -mode sine -freq 0.5 -spread 2.0
```

```bash {image}
\![Sine mode: -mode sine -freq 0.5 -spread 2.0](demo-sine.svg)
```

![Sine mode: -mode sine -freq 0.5 -spread 2.0](13b61472-2026-02-26.svg)

## Random mode

Each line gets one color, chosen by a seeded PRNG keyed on `-seed + line number`. Output is fully deterministic.

```bash
printf "Painful zombies quickly watch a jinxed graveyard.\n..." | lolbat -mode random
```

```bash {image}
\![Random mode: -mode random](demo-random.svg)
```

![Random mode: -mode random](79524458-2026-02-26.svg)

## Mode comparison

Same pangrams run through all three modes.

```bash
printf "..." | lolbat -spread 2.0          # sequential
printf "..." | lolbat -mode sine -spread 2.0
printf "..." | lolbat -mode random
```

```bash {image}
\![Mode comparison: sequential / sine / random](demo-modes.svg)
```

![Mode comparison: sequential / sine / random](9dcf0234-2026-02-26.svg)

## Spread comparison

`-spread 1` changes color every character. `-spread 3` (default) groups ~3 chars per color. `-spread 6` produces broader bands.

```bash
printf "..." | lolbat -spread 1.0
printf "..." | lolbat -spread 3.0
printf "..." | lolbat -spread 6.0
```

```bash {image}
\![Spread comparison: -spread 1 / 3 / 6](demo-spread.svg)
```

![Spread comparison: -spread 1 / 3 / 6](5ebdd4c0-2026-02-26.svg)

## Seed variations

`-seed` shifts the starting position in the palette. Each value produces a different color entry point.

```bash
printf "..." | lolbat -seed 0 -spread 2.0
printf "..." | lolbat -seed 3 -spread 2.0
printf "..." | lolbat -seed 7 -spread 2.0
printf "..." | lolbat -seed 10 -spread 2.0
```

```bash {image}
\![Seed variations: -seed 0 / 3 / 7 / 10](demo-seed.svg)
```

![Seed variations: -seed 0 / 3 / 7 / 10](6b00f5f5-2026-02-26.svg)

## Freq comparison (sine mode)

Higher `-freq` compresses the sine wave, cycling through the palette faster.

```bash
printf "..." | lolbat -mode sine -freq 0.1 -spread 2.0
printf "..." | lolbat -mode sine -freq 0.3 -spread 2.0
printf "..." | lolbat -mode sine -freq 0.6 -spread 2.0
printf "..." | lolbat -mode sine -freq 1.0 -spread 2.0
```

```bash {image}
\![Freq comparison: -mode sine -freq 0.1 / 0.3 / 0.6 / 1.0](demo-freq.svg)
```

![Freq comparison: -mode sine -freq 0.1 / 0.3 / 0.6 / 1.0](74c6d2dc-2026-02-26.svg)

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
ok  	github.com/jamesfishwick/lolbat	(cached)
```

## Source

- GitHub: https://github.com/jamesfishwick/lolbat
- Go 1.21+, stdlib only — no external dependencies.

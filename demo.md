# lolbat: Dracula-themed terminal colorizer

*2026-02-26T04:13:37Z by Showboat 0.6.1*
<!-- showboat-id: d07e99da-6e91-4d16-9130-16e2950b2c3d -->

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

```bash {image}
![Sequential mode: -spread 1.0](demo-sequential.svg)
```

![Sequential mode: -spread 1.0](cb277cc9-2026-02-26.svg)

## Sine mode

The palette index oscillates along a sine wave. `-freq` controls wave frequency; `-spread` controls horizontal compression.

```bash {image}
![Sine mode: -mode sine -freq 0.5 -spread 2.0](demo-sine.svg)
```

![Sine mode: -mode sine -freq 0.5 -spread 2.0](b1f3d5d1-2026-02-26.svg)

## Random mode

Each line gets one color, chosen by a seeded PRNG keyed on `-seed + line number`. Output is fully deterministic.

```bash {image}
![Random mode: -mode random](demo-random.svg)
```

![Random mode: -mode random](34570275-2026-02-26.svg)

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

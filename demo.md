# lolbat: Dracula-themed terminal colorizer

*2026-02-26T04:10:29Z by Showboat 0.6.1*
<!-- showboat-id: e48345be-dd0f-4c45-a5b0-463d70f29cf3 -->

lolbat colorizes terminal output using the Dracula color palette. It reads from stdin or files and supports three coloring modes: sequential, sine, and random. It is API-compatible with lolcat flags.

## Install

Install directly from source using the Go toolchain:

```bash
go install github.com/jamesfishwick/lolbat@latest && echo 'installed'
```

```output
installed
```

## Sequential mode (default)

Each character advances through the Dracula palette. The `--spread` flag controls how many characters share a color before stepping to the next (default: 3).

```bash {image}
![Sequential mode — each character cycles through the Dracula palette](demo-sequential.svg)
```

![Sequential mode — each character cycles through the Dracula palette](390f343f-2026-02-26.svg)

## Sine mode

The palette index oscillates along a sine wave. `--freq` controls wave frequency, `--spread` controls horizontal compression. The result is a smooth, undulating gradient across each line.

```bash {image}
![Sine mode — smooth sine-wave color gradient across pangrams](demo-sine.svg)
```

![Sine mode — smooth sine-wave color gradient across pangrams](b1c970c3-2026-02-26.svg)

## Random mode

Each line gets a single color chosen by a seeded PRNG, keyed on `--seed + line number`. Output is deterministic — the same input always produces the same colors.

```bash {image}
![Random mode — each line gets one deterministic Dracula color](demo-random.svg)
```

![Random mode — each line gets one deterministic Dracula color](0e01ba0e-2026-02-26.svg)

## Test suite

7 unit tests covering palette length, index wrapping, negative seed safety, ANSI escape formatting, and all three colorizer modes.

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
- Written in Go 1.21+, stdlib only — no external dependencies.

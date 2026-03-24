# trumprng

> *"Nobody knows more about randomness than me. Nobody. Believe me."*

![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Tests](https://img.shields.io/badge/tests-18%2F18%20passing-brightgreen)
![Quotes](https://img.shields.io/badge/entropy%20source-73%20quotes-red)
![Chi²](https://img.shields.io/badge/chi²%20test-passing-blue)

A statistically sound pseudo-random number generator seeded by Donald Trump quotes. Implements `math/rand.Source64`. 73 quotes. All tests pass. Believe me.

## How it works

```
Quote  ──→  FNV-1a hash (64-bit)  ──→  SplitMix64 PRNG  ──→  Random numbers
```

1. The quote is normalized to lowercase and hashed via **FNV-1a** into a 64-bit seed
2. The seed initializes a **SplitMix64** generator (passes BigCrush, zero allocations per call)
3. The result is fully compatible with `math/rand.Source` and `math/rand.Source64`

## Installation

```bash
go get github.com/example/trumprng
```

## Usage

```go
// Random quote as seed
rng := trumprng.New()

// Specific quote as seed
rng := trumprng.NewFromQuote("I have the best words.")

// Deterministic seeding — useful for reproducible tests
rng := trumprng.NewFromQuote("my deterministic seed text")

// Which quote was used?
fmt.Println(rng.Quote().Text)    // "I have the best words. Nobody has better words than me."
fmt.Println(rng.Quote().Year)    // 2015
fmt.Printf("Seed: 0x%X\n", rng.Seed64())

// Generate random values
rng.Float64()        // [0.0, 1.0)
rng.Intn(100)        // [0, 100)
rng.Bool()           // true or false
rng.Perm(10)         // permutation of [0..9]
rng.Shuffle(n, swap) // Fisher-Yates in-place shuffle

// Drop-in replacement for math/rand
r := rand.New(rng)
r.NormFloat64()
r.ExpFloat64()

// Shannon entropy of the underlying quote in bits
fmt.Printf("Entropy: %.2f bits\n", rng.Entropy())

// Browse the quote database
for _, q := range trumprng.AllQuotes() {
    fmt.Printf("[%d] %s\n", q.Year, q.Text)
}
fmt.Println(trumprng.QuoteCount()) // 73
```

## Statistical quality

```
Chi-squared test   (10 buckets, n=100,000):   χ² < 27.88   ✓
Serial correlation (n=10,000):                |r| < 0.02   ✓
Bool balance       (n=10,000,000):            ~50.000%     ✓
```

## Quote database

73 carefully curated quotes spanning 1987–2026, organized by theme:

- **Self-assessment** — "I know more about X than anybody" (X varies)
- **Nature & science** — wind turbines, climate change, fish
- **Diplomacy** — Canada as the 51st state, Greenland, Gaza as the Riviera of the Middle East
- **Economics** — tariffs, trade wars, "the word affordability is a Democrat scam"
- **Divine mandate** — "I was saved by God to make America great again"
- **Health & medicine** — disinfectant injections, coronavirus
- **Media & communication** — fake news, Twitter, gut feelings

## Benchmarks

```
BenchmarkUint64-8      61,370,678    19.6 ns/op    0 B/op    0 allocs/op
BenchmarkFloat64-8     56,586,916    24.3 ns/op    0 B/op    0 allocs/op
BenchmarkIntn-8        45,849,954    29.1 ns/op    0 B/op    0 allocs/op
BenchmarkHashQuote-8    4,504,376   247.9 ns/op   64 B/op    1 allocs/op
```

## Why?

Because nobody generates better random numbers than us. Nobody.

Also, this project demonstrates:
- `math/rand.Source` / `Source64` interface implementation
- SplitMix64 PRNG with zero dependencies
- FNV-1a string hashing in pure Go
- Deterministic seeding from arbitrary strings
- Shannon entropy calculation without `math` import

## License

MIT — Use it bigly.

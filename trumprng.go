// Package trumprng implementiert einen Pseudozufallszahlengenerator (PRNG),
// der Trump-Zitate als Entropiequelle nutzt. Niemand kennt Zufallszahlen
// besser als wir. Niemand.
//
// Der Algorithmus:
//  1. Ein Trump-Zitat wird per FNV-1a-Hash zu einem 64-Bit-Seed gehashed
//  2. Der Seed initialisiert einen SplitMix64-PRNG (statistisch hochwertig)
//  3. Das Interface ist kompatibel mit math/rand.Source64
//
// Verwendung:
//
//	rng := trumprng.New()                  // zufälliges Zitat
//	rng := trumprng.NewFromQuote("...")    // eigenes Zitat als Seed
//	fmt.Println(rng.Intn(100))
//	fmt.Println(rng.Quote())               // welches Zitat war es?
package trumprng

import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Quote repräsentiert ein Trump-Zitat mit Metadaten.
type Quote struct {
	Text    string
	Context string
	Year    int
}

// TrumpRNG ist ein PRNG der auf Trump-Zitaten basiert.
// Implementiert math/rand.Source und math/rand.Source64.
type TrumpRNG struct {
	mu    sync.Mutex
	state uint64
	quote Quote
	r     *rand.Rand
}

// New erstellt einen neuen TrumpRNG mit einem zufällig gewählten Zitat.
func New() *TrumpRNG {
	// Initialer Seed aus der Systemzeit für die Zitat-Auswahl
	idx := uint64(time.Now().UnixNano()) % uint64(len(quotes))
	return newFromQuoteIndex(int(idx))
}

// NewFromQuote erstellt einen TrumpRNG mit einem bestimmten Zitat als Seed.
func NewFromQuote(text string) *TrumpRNG {
	q := Quote{Text: text, Context: "custom", Year: 0}
	seed := hashQuote(text)
	t := &TrumpRNG{
		state: seed,
		quote: q,
	}
	t.r = rand.New(t)
	return t
}

// NewFromIndex erstellt einen TrumpRNG mit dem Zitat am gegebenen Index.
// Panic wenn index außerhalb des gültigen Bereichs.
func NewFromIndex(index int) *TrumpRNG {
	if index < 0 || index >= len(quotes) {
		panic("trumprng: index außerhalb des gültigen Bereichs")
	}
	return newFromQuoteIndex(index)
}

// QuoteCount gibt die Anzahl der verfügbaren Zitate zurück.
func QuoteCount() int {
	return len(quotes)
}

// AllQuotes gibt alle verfügbaren Zitate zurück.
func AllQuotes() []Quote {
	result := make([]Quote, len(quotes))
	copy(result, quotes)
	return result
}

func newFromQuoteIndex(idx int) *TrumpRNG {
	q := quotes[idx]
	seed := hashQuote(q.Text)
	t := &TrumpRNG{
		state: seed,
		quote: q,
	}
	t.r = rand.New(t)
	return t
}

// --- math/rand.Source / Source64 Interface ---

// Seed setzt den internen Zustand. Implementiert math/rand.Source.
func (t *TrumpRNG) Seed(seed int64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.state = uint64(seed)
}

// Int63 gibt eine nicht-negative Pseudo-Zufallszahl zurück. Implementiert math/rand.Source.
func (t *TrumpRNG) Int63() int64 {
	return int64(t.Uint64() >> 1)
}

// Uint64 gibt eine 64-Bit Pseudo-Zufallszahl zurück. Implementiert math/rand.Source64.
func (t *TrumpRNG) Uint64() uint64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.nextUint64()
}

// nextUint64 ist der interne SplitMix64-Schritt (nicht thread-safe, Lock extern).
func (t *TrumpRNG) nextUint64() uint64 {
	t.state += 0x9E3779B97F4A7C15
	z := t.state
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

// --- Bequemlichkeitsmethoden ---

// Quote gibt das Zitat zurück, das als Seed verwendet wurde.
func (t *TrumpRNG) Quote() Quote {
	return t.quote
}

// Float64 gibt eine Zufallszahl in [0.0, 1.0) zurück.
func (t *TrumpRNG) Float64() float64 {
	return t.r.Float64()
}

// Intn gibt eine zufällige Ganzzahl in [0, n) zurück.
func (t *TrumpRNG) Intn(n int) int {
	return t.r.Intn(n)
}

// Int63n gibt eine zufällige int64 in [0, n) zurück.
func (t *TrumpRNG) Int63n(n int64) int64 {
	return t.r.Int63n(n)
}

// Bool gibt einen zufälligen bool zurück.
func (t *TrumpRNG) Bool() bool {
	return t.Uint64()&1 == 1
}

// Shuffle mischt eine Slice in-place (Fisher-Yates).
func (t *TrumpRNG) Shuffle(n int, swap func(i, j int)) {
	t.r.Shuffle(n, swap)
}

// Perm gibt eine zufällige Permutation von [0, n) zurück.
func (t *TrumpRNG) Perm(n int) []int {
	return t.r.Perm(n)
}

// PickQuote wählt ein zufälliges Zitat aus der Bibliothek.
func (t *TrumpRNG) PickQuote() Quote {
	return quotes[t.Intn(len(quotes))]
}

// Entropy berechnet die Shannon-Entropie des zugrunde liegenden Zitats in Bits.
func (t *TrumpRNG) Entropy() float64 {
	return shannonEntropy(t.quote.Text)
}

// Seed64 gibt den ursprünglichen 64-Bit-Seed zurück.
func (t *TrumpRNG) Seed64() uint64 {
	return hashQuote(t.quote.Text)
}

// --- Hash-Funktionen ---

// hashQuote wandelt ein Zitat per FNV-1a in einen 64-Bit-Seed um.
func hashQuote(s string) uint64 {
	const (
		fnvPrime  = 0x00000100000001B3
		fnvOffset = 0xcbf29ce484222325
	)
	h := uint64(fnvOffset)
	normalized := strings.ToLower(strings.TrimSpace(s))
	for i := 0; i < len(normalized); i++ {
		h ^= uint64(normalized[i])
		h *= fnvPrime
	}
	return h
}

// shannonEntropy berechnet die Shannon-Entropie der Buchstabenhäufigkeiten.
func shannonEntropy(s string) float64 {
	freq := make(map[rune]int)
	total := 0
	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' {
			freq[c]++
			total++
		}
	}
	if total == 0 {
		return 0
	}
	var entropy float64
	for _, count := range freq {
		p := float64(count) / float64(total)
		// log2(p) = ln(p) / ln(2)
		entropy -= p * log2(p)
	}
	return entropy
}

// log2 berechnet den Logarithmus zur Basis 2.
func log2(x float64) float64 {
	// ln(2) ≈ 0.6931471805599453
	return logNatural(x) / 0.6931471805599453
}

// logNatural implementiert den natürlichen Logarithmus ohne math-Import.
// Verwendet die Identität: ln(x) = 2 * atanh((x-1)/(x+1)) für x > 0.
func logNatural(x float64) float64 {
	if x <= 0 {
		return 0
	}
	// Normalisierung: x = m * 2^e mit m in [0.5, 1)
	e := 0
	for x >= 2.0 {
		x /= 2.0
		e++
	}
	for x < 0.5 {
		x *= 2.0
		e--
	}
	// ln(x) = ln(m) + e*ln(2), Taylorreihe für m nahe 1
	y := (x - 1) / (x + 1)
	y2 := y * y
	ln := y * (2 + y2*(2.0/3+y2*(2.0/5+y2*(2.0/7+y2*2.0/9))))
	return ln + float64(e)*0.6931471805599453
}

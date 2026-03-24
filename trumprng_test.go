package trumprng

import (
	"math"
	"testing"
)

// --- Unit Tests ---

func TestNew(t *testing.T) {
	rng := New()
	if rng == nil {
		t.Fatal("New() hat nil zurückgegeben")
	}
	q := rng.Quote()
	if q.Text == "" {
		t.Error("Quote.Text ist leer")
	}
	if q.Year == 0 {
		t.Error("Quote.Year ist 0")
	}
}

func TestNewFromQuote(t *testing.T) {
	quote := "I have the best words. Nobody has better words than me."
	rng := NewFromQuote(quote)
	if rng.Quote().Text != quote {
		t.Errorf("erwartet %q, bekommen %q", quote, rng.Quote().Text)
	}
}

func TestDeterministisch(t *testing.T) {
	// Gleiche Zitat → gleiche Sequenz
	q := "Nobody knows randomness better than me."
	rng1 := NewFromQuote(q)
	rng2 := NewFromQuote(q)

	for i := range 20 {
		a, b := rng1.Uint64(), rng2.Uint64()
		if a != b {
			t.Errorf("Schritt %d: %d != %d — nicht deterministisch!", i, a, b)
		}
	}
}

func TestUnterschiedlicheZitate(t *testing.T) {
	// Verschiedene Zitate → verschiedene Sequenzen
	rng1 := NewFromQuote("I have the best words.")
	rng2 := NewFromQuote("The windmills are killing all the birds.")

	gleich := 0
	n := 100
	for range n {
		if rng1.Uint64() == rng2.Uint64() {
			gleich++
		}
	}
	// Höchstens 1% Kollision erwartet bei guter Streuung
	if gleich > n/100 {
		t.Errorf("Zu viele Kollisionen: %d/%d", gleich, n)
	}
}

func TestFloat64Bereich(t *testing.T) {
	rng := NewFromQuote("believe me")
	for i := range 10_000 {
		f := rng.Float64()
		if f < 0 || f >= 1.0 {
			t.Fatalf("Float64() außerhalb [0,1) bei Schritt %d: %f", i, f)
		}
	}
}

func TestIntnBereich(t *testing.T) {
	rng := New()
	for _, n := range []int{1, 2, 10, 100, 1000} {
		for i := range 1000 {
			v := rng.Intn(n)
			if v < 0 || v >= n {
				t.Fatalf("Intn(%d) = %d bei Schritt %d: außerhalb [0,%d)", n, v, i, n)
			}
		}
	}
}

func TestBool(t *testing.T) {
	rng := New()
	trueCount := 0
	n := 10_000
	for range n {
		if rng.Bool() {
			trueCount++
		}
	}
	// Erwarte etwa 50% — ±5% Toleranz
	ratio := float64(trueCount) / float64(n)
	if ratio < 0.45 || ratio > 0.55 {
		t.Errorf("Bool() unausgeglichen: %.2f%% true (erwartet ~50%%)", ratio*100)
	}
}

func TestShuffle(t *testing.T) {
	rng := New()
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	original := make([]int, len(s))
	copy(original, s)

	rng.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })

	// Prüfe dass alle Elemente noch vorhanden sind
	sumOrig, sumShuf := 0, 0
	for i := range s {
		sumOrig += original[i]
		sumShuf += s[i]
	}
	if sumOrig != sumShuf {
		t.Error("Shuffle hat Elemente verloren oder verändert")
	}
}

func TestPerm(t *testing.T) {
	rng := New()
	n := 20
	p := rng.Perm(n)
	if len(p) != n {
		t.Fatalf("Perm(%d) hat %d Elemente", n, len(p))
	}
	seen := make(map[int]bool)
	for _, v := range p {
		if v < 0 || v >= n {
			t.Fatalf("Perm enthält ungültigen Wert %d", v)
		}
		if seen[v] {
			t.Fatalf("Perm enthält Duplikat: %d", v)
		}
		seen[v] = true
	}
}

func TestHashKonsistenz(t *testing.T) {
	// FNV-1a muss für gleiche Strings immer denselben Hash liefern
	text := "I will be the greatest jobs president that God ever created."
	h1 := hashQuote(text)
	h2 := hashQuote(text)
	if h1 != h2 {
		t.Errorf("hashQuote() ist nicht konsistent: %d != %d", h1, h2)
	}
}

func TestHashCaseInsensitive(t *testing.T) {
	// Groß-/Kleinschreibung soll denselben Seed ergeben
	h1 := hashQuote("BELIEVE ME")
	h2 := hashQuote("believe me")
	h3 := hashQuote("Believe Me")
	if h1 != h2 || h2 != h3 {
		t.Error("hashQuote() sollte case-insensitive sein")
	}
}

func TestEntropie(t *testing.T) {
	rng := NewFromQuote("I have the best words. Nobody has better words than me.")
	e := rng.Entropy()
	// Shannon-Entropie für englischen Text typisch 3.5–4.5 Bits
	if e < 2.0 || e > 5.0 {
		t.Errorf("unplausible Entropie: %.4f", e)
	}
}

func TestQuoteCount(t *testing.T) {
	if QuoteCount() < 10 {
		t.Errorf("Zu wenige Zitate: %d", QuoteCount())
	}
}

func TestAllQuotes(t *testing.T) {
	qs := AllQuotes()
	if len(qs) != QuoteCount() {
		t.Errorf("AllQuotes() hat %d, QuoteCount() sagt %d", len(qs), QuoteCount())
	}
	// Prüfe dass es eine Kopie ist (keine Mutation des Originals)
	qs[0].Text = "mutiert"
	if AllQuotes()[0].Text == "mutiert" {
		t.Error("AllQuotes() gibt Referenz statt Kopie zurück")
	}
}

func TestNewFromIndexPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("NewFromIndex(-1) hätte paniken sollen")
		}
	}()
	NewFromIndex(-1)
}

func TestRandSourceKompatibilitaet(t *testing.T) {
	// TrumpRNG muss als math/rand.Source funktionieren
	rng := New()
	for range 1000 {
		v := rng.Int63()
		if v < 0 {
			t.Fatal("Int63() hat negativen Wert zurückgegeben")
		}
	}
}

// --- Statistische Tests ---

func TestChiQuadrat(t *testing.T) {
	// Chi-Quadrat-Test auf Gleichverteilung von Intn(10)
	rng := NewFromQuote("Statistics are beautiful when they work for you.")
	buckets := make([]int, 10)
	n := 100_000

	for range n {
		buckets[rng.Intn(10)]++
	}

	expected := float64(n) / 10.0
	chi2 := 0.0
	for _, count := range buckets {
		diff := float64(count) - expected
		chi2 += (diff * diff) / expected
	}

	// Kritischer Wert für chi2 mit df=9, p=0.001: ~27.88
	if chi2 > 27.88 {
		t.Errorf("Chi²-Test fehlgeschlagen: χ²=%.2f > 27.88 (Gleichverteilung unwahrscheinlich)", chi2)
	}
}

func TestSerielleKorrelation(t *testing.T) {
	// Prüfe serielle Korrelation zwischen aufeinanderfolgenden Werten
	rng := NewFromQuote("There is no correlation. None.")
	n := 10_000
	values := make([]float64, n)
	for i := range n {
		values[i] = rng.Float64()
	}

	// Pearson-Korrelation zwischen x[i] und x[i+1]
	var sumX, sumY, sumXY, sumX2, sumY2 float64
	m := float64(n - 1)
	for i := 0; i < n-1; i++ {
		x, y := values[i], values[i+1]
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}
	numerator := m*sumXY - sumX*sumY
	denom := math.Sqrt((m*sumX2 - sumX*sumX) * (m*sumY2 - sumY*sumY))
	r := numerator / denom

	// |r| < 0.02 erwartet für guten PRNG
	if math.Abs(r) > 0.02 {
		t.Errorf("serielle Korrelation zu hoch: r=%.4f (erwartet |r| < 0.02)", r)
	}
}

// --- Benchmarks ---

func BenchmarkUint64(b *testing.B) {
	rng := New()
	b.ResetTimer()
	for range b.N {
		_ = rng.Uint64()
	}
}

func BenchmarkFloat64(b *testing.B) {
	rng := New()
	b.ResetTimer()
	for range b.N {
		_ = rng.Float64()
	}
}

func BenchmarkIntn(b *testing.B) {
	rng := New()
	b.ResetTimer()
	for range b.N {
		_ = rng.Intn(1000)
	}
}

func BenchmarkHashQuote(b *testing.B) {
	q := quotes[0].Text
	b.ResetTimer()
	for range b.N {
		_ = hashQuote(q)
	}
}

func BenchmarkNew(b *testing.B) {
	for range b.N {
		_ = New()
	}
}

// Dieses Beispiel zeigt die Verwendung von trumprng.
package trumprng_test

import (
	"fmt"

	"github.com/example/trumprng"
)

func ExampleNew() {
	// Neuen RNG mit zufälligem Zitat initialisieren
	rng := trumprng.NewFromQuote("I have the best words. Nobody has better words than me.")

	fmt.Printf("Zitat: %q\n", rng.Quote().Text)
	fmt.Printf("Jahr:  %d\n", rng.Quote().Year)
	fmt.Printf("Seed:  0x%X\n", rng.Seed64())

	// Output wird variieren — daher kein festes Output-Kommentar
}

func ExampleTrumpRNG_Intn() {
	rng := trumprng.NewFromQuote("believe me")

	// 10 Würfelwürfe (1–6)
	for range 10 {
		fmt.Printf("%d ", rng.Intn(6)+1)
	}
	fmt.Println()
}

func ExampleTrumpRNG_Shuffle() {
	rng := trumprng.NewFromQuote("Trade wars are good, and easy to win.")

	kandidaten := []string{
		"Trump", "Biden", "Obama", "Clinton", "Bush",
	}
	rng.Shuffle(len(kandidaten), func(i, j int) {
		kandidaten[i], kandidaten[j] = kandidaten[j], kandidaten[i]
	})
	fmt.Println(kandidaten)
}

func ExampleTrumpRNG_Bool() {
	rng := trumprng.NewFromQuote("We will have so much winning.")

	// Münzwurf-Simulation
	kopf := 0
	n := 1_000_000
	for range n {
		if rng.Bool() {
			kopf++
		}
	}
	fmt.Printf("Kopf: %.1f%%\n", float64(kopf)/float64(n)*100)
}

func ExampleAllQuotes() {
	qs := trumprng.AllQuotes()
	fmt.Printf("Verfügbare Zitate: %d\n", len(qs))

	// Ältestes Zitat
	oldest := qs[0]
	for _, q := range qs {
		if q.Year > 0 && (oldest.Year == 0 || q.Year < oldest.Year) {
			oldest = q
		}
	}
	fmt.Printf("Ältestes Zitat (%d): %q\n", oldest.Year, oldest.Text[:50])
}

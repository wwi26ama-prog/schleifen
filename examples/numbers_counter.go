package examples

// Diese Datei enthält Beispiele für einfache Schleifen,
// die über Zahlenbereiche iterieren.

import "fmt"

// LoopWithCounter erwartet eine Zahl `n` und gibt alle
// Zahlen von `0` bis `n-1` auf der Konsole aus.
// Für `n < 0` passiert nichts.
func LoopWithCounter(n int) {
	for i := 0; i < n; i += 1 {
		fmt.Println(i)
	}
	// Einfache Zähler-Schleife:
	// - Eine Zählvariable `i` wird mit `0` initialisiert.
	// - Die Schleife läuft, solange `i < n`.
	// - Nach jedem Schritt wird `i` um `1` erhöht.
}

// EvenNumbersCounter1 erwartet eine Zahl `n` und gibt alle
// geraden Zahlen von `0` bis `n-1` auf der Konsole aus.
// Für `n < 0` passiert nichts.
func EvenNumbersCounter1(n int) {
	for i := 0; i < n; i += 1 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	// Zähler-Schleife mit zusätzlicher Bedingung:
	// - In jedem Schritt wird geprüft, ob sie gerade ist.
	// - Davon abgesehen ist die Schleife identisch zu einer normalen Zähler-Schleife.
}

// EvenNumbersCounter2 erwartet eine Zahl `n` und gibt alle
// geraden Zahlen von `0` bis `n-1` auf der Konsole aus.
// Für `n < 0` passiert nichts.
func EvenNumbersCounter2(n int) {
	for i := 0; i < n; i += 2 {
		fmt.Println(i)
	}
	// Zähler-Schleife mit Schrittweite 2:
	// - `i` wird in jedem Schritt direkt um `2` erhöht.
	// - Dadurch ist die Bedingung `i%2 == 0` nicht mehr erforderlich.
}

// SumNumbersCounter erwartet eine Zahl `n` und gibt die Summe aller
// positiven ganzen Zahlen von `0` bis `n-1` als Ergebnis zurück.
// Für `n < 0` wird `0` zurückgegeben.
func SumNumbersCounter(n int) int {
	sum := 0
	for i := 0; i < n; i += 1 {
		sum += i
	}
	return sum
	// Zähler-Schleife zur Berechnung der Summe:
	// - Vor der Schleife wird eine Hilfsvariable für das Ergebnis definiert.
	// - Die Schleife funktioniert wie die bekannten Zähler-Schleifen.
	// - Statt die Werte auf der Konsole auszugeben, werden sie zur Summe addiert.
	// - Am Ende wird die Gesamtsumme zurückgegeben.
}

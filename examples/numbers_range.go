package examples

import "fmt"

// Rangeloop erwartet eine Zahl `n` und gibt alle
// Zahlen von `0` bis `n-1` auf der Konsole aus.
// Für `n < 0` passiert nichts.
func RangeLoop(n int) {
	for i := range n {
		fmt.Println(i)
	}
	// Range-basierte Schleife:
	// - Verhalten ist identisch zu einer Zähler-Schleife von `0` bis `n-1`.
	// - `i := range n` ersetzt die manuelle Verwaltung der Zählvariable.
}

// EvenNumbersRange1 erwartet eine Zahl `n` und gibt alle
// geraden Zahlen von `0` bis `n-1` auf der Konsole aus.
// Für `n < 0` passiert nichts.
func EvenNumbersRange1(n int) {
	for i := range n {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	// Range-basierte Schleife mit zusätzlicher Bedingung:
	// - In jedem Schritt wird geprüft, ob die Zahl gerade ist.
	// - Davon abgesehen ist die Schleife identisch zu einer normalen Range-Schleife.
}

// EvenNumbersRange2 erwartet eine Zahl `n` und gibt alle
// geraden Zahlen von `0` bis `n-1` auf der Konsole aus.
// Für `n < 0` passiert nichts.
func EvenNumbersRange2(n int) {
	for i := range n / 2 {
		fmt.Println(i * 2)
	}
	// Range-basierte Schleife:
	// - Statt zu überprüfen, ob die Zahl gerade ist, läuft die Schleife
	//   nur bis `n/2` und multipliziert jedes `i` mit `2`.
}

// SumNumbersRange erwartet eine Zahl `n` und gibt die Summe aller
// positiven ganzen Zahlen von `0` bis `n-1` als Ergebnis zurück.
// Für `n < 0` wird `0` zurückgegeben.
func SumNumbersRange(n int) int {
	sum := 0
	for i := range n {
		sum += i
	}
	return sum
	// Range-basierte Schleife zur Berechnung der Summe:
	// - Vor der Schleife wird eine Hilfsvariable für das Ergebnis definiert.
	// - Die Schleife funktioniert wie die bekannten Range-Schleifen.
	// - Statt die Werte auf der Konsole auszugeben, werden sie zur Summe addiert.
	// - Am Ende wird die Gesamtsumme zurückgegeben.
}

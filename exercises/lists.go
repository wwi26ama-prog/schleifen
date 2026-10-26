package exercises

// CountOdd erwartet eine Liste von Zahlen und liefert
// die Anzahl der ungeraden Zahlen.
func CountOdd(numbers []int) int {
	// Hinweis:
	// Gehen Sie vor wie in `CountEven` aus den Beispielen.
	// Nur die Bedingung im Schleifenrumpf muss angepasst werden.

	count := 0
	for _, number := range numbers {
		if number%2 != 0 {
			count += 1
			// Anmerkung:
			// Ungerade Zahlen mit negativem Vorzeichen:
			// - Der Rest bei `%` hat in Go dasselbe Vorzeichen wie die geteilte Zahl.
			//   Deshalb ist `-3 % 2` nicht `1`, sondern `-1`.
			// - Eine Bedingung wie `number%2 == 1` erkennt negative ungerade
			//   Zahlen deshalb nicht.
			// - `number%2 != 0` funktioniert dagegen für alle Zahlen.
		}
	}
	return count
}

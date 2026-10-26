package examples

// Diese Datei enthält Beispiele für Schleifen über Listen (Slices).
// Die Schleifen werten Eigenschaften der Elemente aus, etwa ob eine
// Zahl gerade ist oder wie lang ein String ist.
// Je nach Aufgabe wird die Form der Schleife gewählt, die am besten passt.

import "fmt"

// PrintList erwartet eine Liste von Zahlen und gibt jedes
// Element in einer eigenen Zeile auf der Konsole aus.
func PrintList(numbers []int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
	// Range-Schleife über eine Liste:
	// - `range numbers` liefert in jedem Schritt zwei Werte:
	//   den Index und das Element an dieser Stelle.
	// - Der Index wird hier nicht benötigt und deshalb mit `_` verworfen.
	// - `number` enthält direkt das Element, ein Zugriff über
	//   `numbers[i]` ist nicht nötig.
	// - Für eine leere Liste wird der Schleifenrumpf nie ausgeführt.
}

// PrintListWithIndex erwartet eine Liste von Zahlen und gibt jedes
// Element zusammen mit seiner Position (Index) auf der Konsole aus.
func PrintListWithIndex(numbers []int) {
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	// Range-Schleife mit Index und Element:
	// - Diesmal werden beide Werte von `range` verwendet.
	// - Der Index beginnt wie bei der Zähler-Schleife bei `0`.
}

// CountEven erwartet eine Liste von Zahlen und liefert
// die Anzahl der geraden Zahlen.
func CountEven(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if number%2 == 0 {
			count += 1
		}
	}
	return count
	// Zählen mit Bedingung:
	// - Die Zählvariable wird nur erhöht, wenn die Bedingung zutrifft.
	// - Die Zahl selbst wird nicht aufsummiert, nur die Treffer.
}

// Sum erwartet eine Liste von Zahlen und liefert deren Summe.
// Für eine leere Liste wird `0` zurückgegeben.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
	// Summe über eine Liste:
	// - Das Vorgehen entspricht der Summe über einen Zahlenbereich.
	// - Statt des Zählers wird jeweils das Element addiert.
}

// TotalLength erwartet eine Liste von Strings und liefert
// die Summe ihrer Längen.
func TotalLength(words []string) int {
	total := 0
	for _, word := range words {
		total += len(word)
	}
	return total
	// Summe über eine Eigenschaft der Elemente:
	// - Aufsummiert werden nicht die Elemente selbst, sondern
	//   ein Wert, der aus jedem Element berechnet wird: `len(word)`.
	// - Achtung: `len` zählt Bytes, nicht Zeichen. Für Wörter ohne
	//   Umlaute und Sonderzeichen ist das dasselbe.
}

// Contains erwartet eine Liste von Zahlen und eine gesuchte Zahl
// und liefert `true`, falls die Zahl in der Liste vorkommt.
func Contains(numbers []int, wanted int) bool {
	for _, number := range numbers {
		if number == wanted {
			return true
		}
	}
	return false
	// Suche mit vorzeitigem Ende:
	// - Sobald die Zahl gefunden ist, steht das Ergebnis fest.
	//   `return` beendet dann die Schleife und die ganze Funktion.
	// - `false` steht erst fest, wenn alle Elemente geprüft wurden.
	//   Deshalb steht dieses `return` hinter der Schleife.
}

// Max erwartet eine Liste von Zahlen und liefert die größte Zahl.
// Für eine leere Liste wird `0` zurückgegeben.
func Max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	largest := numbers[0]
	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}
	return largest
	// Größtes Element bestimmen:
	// - Die Hilfsvariable startet mit dem ersten Element. Mit `0` als
	//   Startwert wäre das Ergebnis für reine Minuszahlen falsch.
	// - Das erste Element gibt es nur in einer nicht-leeren Liste,
	//   deshalb wird der Fall der leeren Liste vorher behandelt.
	// - Jedes größere Element ersetzt den bisherigen Wert.
}

// CountIncreases erwartet eine Liste von Zahlen und liefert, wie oft
// eine Zahl größer ist als die Zahl direkt davor.
func CountIncreases(numbers []int) int {
	count := 0
	for i := 1; i < len(numbers); i += 1 {
		if numbers[i] > numbers[i-1] {
			count += 1
		}
	}
	return count
	// Zähler-Schleife über eine Liste:
	// - Verglichen wird jedes Element mit seinem Vorgänger `numbers[i-1]`.
	// - Das erste Element hat keinen Vorgänger, deshalb startet
	//   der Zähler bei `1` statt bei `0`.
	// - Hier passt die Zähler-Schleife besser als `range`, weil der
	//   Index und nicht nur das aktuelle Element benötigt wird.
	// - Für Listen mit weniger als zwei Elementen läuft die Schleife nicht.
}

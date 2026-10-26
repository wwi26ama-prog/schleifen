package examples

// Diese Datei enthält Beispiele für Schleifen, die eine Liste von
// Tagestemperaturen in Grad Celsius auswerten, z.B. die Messwerte
// einer Woche. Das erste Element ist die Temperatur von Tag 1.

import "fmt"

// PrintTemperatures erwartet eine Liste von Tagestemperaturen und gibt
// sie zusammen mit der Nummer des Tages auf der Konsole aus.
func PrintTemperatures(temperatures []int) {
	for i, celsius := range temperatures {
		fmt.Printf("day %d: %d°C\n", i+1, celsius)
	}
	// Range-Schleife mit Index:
	// - Der Index beginnt bei `0`, Tage werden aber ab `1` gezählt.
	// - Deshalb wird für die Ausgabe `i+1` verwendet. Der Index
	//   selbst bleibt unverändert.
}

// CountFrostDays erwartet eine Liste von Tagestemperaturen und
// liefert die Anzahl der Tage mit Frost, also unter `0` Grad.
func CountFrostDays(temperatures []int) int {
	count := 0
	for _, celsius := range temperatures {
		if celsius < 0 {
			count += 1
		}
	}
	return count
	// Zählen mit Bedingung:
	// - Gezählt werden nur die Elemente, für die die Bedingung zutrifft.
	// - Die Temperaturen selbst werden nicht addiert, nur die Treffer.
}

// AverageTemperature erwartet eine Liste von Tagestemperaturen und
// liefert deren Durchschnitt. Für eine leere Liste wird `0` zurückgegeben.
func AverageTemperature(temperatures []int) float64 {
	if len(temperatures) == 0 {
		return 0
	}
	sum := 0
	for _, celsius := range temperatures {
		sum += celsius
	}
	return float64(sum) / float64(len(temperatures))
	// Durchschnitt aus Summe und Anzahl:
	// - Die Schleife berechnet nur die Summe.
	// - Die Anzahl muss nicht gezählt werden, sie liefert `len`.
	// - Vor der Division werden beide Werte mit `float64(...)`
	//   umgewandelt, sonst wäre das Ergebnis eine Ganzzahl-Division
	//   und Nachkommastellen gingen verloren.
	// - Ohne die Prüfung auf eine leere Liste würde durch `0` geteilt.
}

// WarmestDay erwartet eine Liste von Tagestemperaturen und liefert
// den Index des wärmsten Tages. Bei gleich warmen Tagen zählt der
// erste davon. Für eine leere Liste wird `-1` zurückgegeben.
func WarmestDay(temperatures []int) int {
	if len(temperatures) == 0 {
		return -1
	}
	warmest := 0
	for i, celsius := range temperatures {
		if celsius > temperatures[warmest] {
			warmest = i
		}
	}
	return warmest
	// Position des größten Elements bestimmen:
	// - Statt des größten Wertes merkt sich die Hilfsvariable dessen Index.
	// - Der Wert ist über den Index jederzeit erreichbar: `temperatures[warmest]`.
	// - `>` statt `>=` sorgt dafür, dass bei gleichen Temperaturen
	//   der erste Tag erhalten bleibt.
	// - `-1` ist kein gültiger Index und zeigt deshalb an,
	//   dass es keinen wärmsten Tag gibt.
}

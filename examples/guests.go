package examples

// Diese Datei enthält Beispiele für Schleifen über eine Gästeliste,
// also eine Liste von Namen. Neben dem Durchsuchen der Liste wird
// gezeigt, wie aus einer Liste eine neue entsteht und wie zwei
// Listen miteinander verglichen werden.

import (
	"fmt"
	"strings"
)

// PrintInvitations erwartet eine Gästeliste und gibt für jeden Gast
// eine Einladung auf der Konsole aus.
func PrintInvitations(guests []string) {
	for _, name := range guests {
		fmt.Println("Dear " + name + ", you are invited!")
	}
	// Range-Schleife über Strings:
	// - Funktioniert genauso wie über Zahlen, `name` ist hier ein `string`.
	// - Für jedes Element wird dieselbe Aktion ausgeführt.
}

// IsInvited erwartet eine Gästeliste und einen Namen und liefert
// `true`, falls der Name auf der Liste steht.
// Groß- und Kleinschreibung wird unterschieden.
func IsInvited(guests []string, name string) bool {
	for _, guest := range guests {
		if guest == name {
			return true
		}
	}
	return false
	// Suche mit vorzeitigem Ende:
	// - Sobald der Name gefunden ist, beendet `return` die Suche.
	// - Erst nach einem vollständigen Durchlauf steht fest,
	//   dass der Name nicht auf der Liste steht.
	// - `==` vergleicht Strings exakt: "anna" ist nicht "Anna".
}

// CleanNames erwartet eine Gästeliste und liefert eine neue Liste,
// in der bei jedem Namen Leerzeichen am Anfang und am Ende entfernt sind.
// Die ursprüngliche Liste bleibt unverändert.
func CleanNames(guests []string) []string {
	cleaned := []string{}
	for _, name := range guests {
		cleaned = append(cleaned, strings.TrimSpace(name))
	}
	return cleaned
	// Neue Liste aus einer bestehenden Liste erzeugen:
	// - Vor der Schleife wird eine leere Ergebnisliste angelegt.
	// - `append` hängt ein Element an und liefert die verlängerte Liste,
	//   die wieder in `cleaned` gespeichert werden muss.
	// - Die eigentliche Arbeit pro Element erledigt eine Funktion aus der
	//   Standardbibliothek: `strings.TrimSpace`. Das Paket `strings`
	//   enthält viele weitere nützliche Funktionen für Strings.
}

// CountCommonGuests erwartet zwei Gästelisten und liefert, wie viele
// Namen auf beiden Listen stehen. Innerhalb einer Liste kommt
// jeder Name höchstens einmal vor.
func CountCommonGuests(guests1 []string, guests2 []string) int {
	count := 0
	for _, name1 := range guests1 {
		for _, name2 := range guests2 {
			if name1 == name2 {
				count += 1
			}
		}
	}
	return count
	// Verschachtelte Schleifen:
	// - Für jeden Namen der ersten Liste läuft die innere Schleife
	//   einmal komplett über die zweite Liste.
	// - So wird jeder Name der ersten Liste mit jedem Namen
	//   der zweiten Liste verglichen.
	// - Bei Listen mit 3 und 4 Namen sind das 3 * 4 = 12 Vergleiche.
	// - Kämen Namen mehrfach vor, würden sie auch mehrfach gezählt.
}

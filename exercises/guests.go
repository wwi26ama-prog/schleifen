package exercises

import "strings"

// LowercaseNames erwartet eine Gästeliste und liefert eine neue Liste,
// in der alle Namen kleingeschrieben sind.
// Die ursprüngliche Liste bleibt unverändert.
func LowercaseNames(guests []string) []string {
	// Hinweis:
	// Gehen Sie vor wie in `CleanNames` aus den Beispielen:
	// - Legen Sie vor der Schleife eine leere Ergebnisliste an.
	// - Hängen Sie in der Schleife jeden umgewandelten Namen mit `append` an.
	// - Für die Umwandlung müssen Sie keine eigene Schleife über die
	//   Buchstaben schreiben. Suchen Sie im Paket `strings` der
	//   Standardbibliothek nach einer passenden Funktion.
	//   Denken Sie daran, das Paket mit `import` einzubinden.

	lowercased := []string{}
	for _, name := range guests {
		lowercased = append(lowercased, strings.ToLower(name))
	}
	return lowercased
}

// HasDuplicates erwartet eine Gästeliste und liefert `true`, falls
// mindestens ein Name mehr als einmal auf der Liste steht.
func HasDuplicates(guests []string) bool {
	// Hinweis:
	// Vergleichen Sie wie in `CountCommonGuests` jeden Namen mit jedem
	// anderen Namen, diesmal innerhalb derselben Liste.
	// - Ein Name darf dabei nicht mit sich selbst verglichen werden.
	//   Dafür brauchen Sie die Indizes der beiden Elemente.
	// - Es genügt, jeden Namen mit den Namen dahinter zu vergleichen:
	//   Die innere Schleife kann deshalb bei `i+1` beginnen.
	// - Sobald ein doppelter Name gefunden ist, steht das Ergebnis fest.

	for i := 0; i < len(guests); i += 1 {
		for j := i + 1; j < len(guests); j += 1 {
			if guests[i] == guests[j] {
				return true
			}
		}
	}
	return false
}

// CountDuplicates erwartet eine Gästeliste und liefert, wie viele
// Einträge überflüssig sind, weil derselbe Name schon weiter vorne
// auf der Liste steht. Für `["Anna", "Ben", "Anna", "Anna"]`
// ist das Ergebnis also `2`.
func CountDuplicates(guests []string) int {
	// Hinweis:
	// Prüfen Sie für jeden Eintrag, ob derselbe Name schon
	// weiter vorne auf der Liste steht.
	// - Die innere Schleife läuft dafür über die Einträge vor dem
	//   aktuellen Eintrag, also von `0` bis `i-1`.
	// - Zählen Sie jeden Eintrag höchstens einmal, auch wenn der Name
	//   weiter vorne mehrfach steht. Mit `break` können Sie die innere
	//   Schleife vorzeitig beenden.

	count := 0
	for i := 0; i < len(guests); i += 1 {
		for j := 0; j < i; j += 1 {
			if guests[i] == guests[j] {
				count += 1
				break
			}
		}
	}
	return count
}

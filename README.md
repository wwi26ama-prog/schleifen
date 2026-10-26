# Schleifen

Diese Einheit zeigt Schleifen mit `for`, zuerst anhand von Zahlenbereichen,
dann anhand von Listen (Slices) und schließlich in Anwendungsbeispielen.

- [`examples/numbers_counter.go`](examples/numbers_counter.go): Zähler-Schleifen
  (alle Zahlen, gerade Zahlen, Summe).
- [`examples/numbers_range.go`](examples/numbers_range.go): Schleifen mit `range` über Zahlen.
- [`examples/lists.go`](examples/lists.go): Schleifen über Listen, die Eigenschaften
  der Elemente auswerten (gerade Zahlen zählen, Summe, Gesamtlänge von Strings,
  Suche, Maximum, Vergleich mit dem Vorgänger).
- [`examples/temperatures.go`](examples/temperatures.go): dieselben Muster angewendet
  auf Tagestemperaturen (Frosttage, Durchschnitt, wärmster Tag).
- [`examples/guests.go`](examples/guests.go): Schleifen über eine Gästeliste
  (Suche, neue Liste mit `append` und `strings.TrimSpace`, verschachtelte
  Schleifen zum Vergleich zweier Listen).

Dazu kommen Übungsaufgaben. Die erste ist eine kurze Aufwärmübung,
die übrigen orientieren sich an den Anwendungsbeispielen:

- [`exercises/lists.go`](exercises/lists.go): ungerade Zahlen zählen.
- [`exercises/guests.go`](exercises/guests.go): Namen kleinschreiben
  (mit Hilfe der Standardbibliothek), doppelte Namen finden und zählen.
- [`exercises/temperatures.go`](exercises/temperatures.go): Tage zählen, die kälter
  als der Vortag sind, und die Spanne zwischen höchster und niedrigster Temperatur.

Die Erklärungen stehen als Kommentare im Quelltext, die Ausgaben
in den zugehörigen Tests. Ausführen mit `go test ./...` in diesem Verzeichnis.

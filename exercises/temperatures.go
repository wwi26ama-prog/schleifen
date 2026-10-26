package exercises

// CountColderDays erwartet eine Liste von Tagestemperaturen in Grad
// Celsius und liefert, an wie vielen Tagen es kälter war als am Vortag.
func CountColderDays(temperatures []int) int {
	// Hinweis:
	// Vergleichen Sie jede Temperatur mit der Temperatur davor,
	// ähnlich wie in `CountIncreases` aus den Beispielen.
	// - Für den Vortag benötigen Sie den Index, z.B. `temperatures[i-1]`.
	// - Der erste Tag hat keinen Vortag. Wo muss die Schleife also beginnen?

	count := 0
	for i := 1; i < len(temperatures); i += 1 {
		if temperatures[i] < temperatures[i-1] {
			count += 1
		}
	}
	return count
}

// TemperatureSpread erwartet eine Liste von Tagestemperaturen in Grad
// Celsius und liefert den Unterschied zwischen der höchsten und der
// niedrigsten Temperatur. Für eine leere Liste wird `0` zurückgegeben.
func TemperatureSpread(temperatures []int) int {
	// Hinweis:
	// Bestimmen Sie in einer einzigen Schleife sowohl die höchste
	// als auch die niedrigste Temperatur.
	// - Verwenden Sie dafür zwei Hilfsvariablen, die beide mit
	//   der ersten Temperatur starten (siehe `Max` in den Beispielen).
	// - Behandeln Sie den Fall der leeren Liste vor der Schleife.
	// - Das Ergebnis ist die Differenz der beiden Hilfsvariablen.

	if len(temperatures) == 0 {
		return 0
	}
	highest := temperatures[0]
	lowest := temperatures[0]
	for _, celsius := range temperatures {
		if celsius > highest {
			highest = celsius
		}
		if celsius < lowest {
			lowest = celsius
		}
	}
	return highest - lowest
}

package exercises

import "fmt"

func ExampleLowercaseNames() {
	guests := []string{"Anna", "BEN", "clara"}
	fmt.Printf("%q\n", LowercaseNames(guests))
	fmt.Printf("%q\n", guests)
	fmt.Printf("%q\n", LowercaseNames([]string{}))

	// Output:
	// ["anna" "ben" "clara"]
	// ["Anna" "BEN" "clara"]
	// []
}

func ExampleHasDuplicates() {
	fmt.Println(HasDuplicates([]string{"Anna", "Ben", "Clara"}))
	fmt.Println(HasDuplicates([]string{"Anna", "Ben", "Anna"}))
	fmt.Println(HasDuplicates([]string{"Anna", "anna"}))
	fmt.Println(HasDuplicates([]string{}))

	// Output:
	// false
	// true
	// false
	// false
}

func ExampleCountDuplicates() {
	fmt.Println(CountDuplicates([]string{"Anna", "Ben", "Clara"}))
	fmt.Println(CountDuplicates([]string{"Anna", "Ben", "Anna", "Anna"}))
	fmt.Println(CountDuplicates([]string{"Anna", "Ben", "Ben", "Anna"}))
	fmt.Println(CountDuplicates([]string{}))

	// Output:
	// 0
	// 2
	// 2
	// 0
}

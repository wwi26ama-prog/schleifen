package examples

import "fmt"

func ExamplePrintInvitations() {
	PrintInvitations([]string{"Anna", "Ben"})

	// Output:
	// Dear Anna, you are invited!
	// Dear Ben, you are invited!
}

func ExampleIsInvited() {
	guests := []string{"Anna", "Ben", "Clara"}
	fmt.Println(IsInvited(guests, "Ben"))
	fmt.Println(IsInvited(guests, "ben"))
	fmt.Println(IsInvited(guests, "David"))

	// Output:
	// true
	// false
	// false
}

func ExampleCleanNames() {
	guests := []string{" Anna", "Ben  ", "  Clara "}
	fmt.Printf("%q\n", CleanNames(guests))
	fmt.Printf("%q\n", guests)

	// Output:
	// ["Anna" "Ben" "Clara"]
	// [" Anna" "Ben  " "  Clara "]
}

func ExampleCountCommonGuests() {
	fmt.Println(CountCommonGuests(
		[]string{"Anna", "Ben", "Clara"},
		[]string{"Clara", "David", "Anna", "Emil"},
	))

	// Output:
	// 2
}

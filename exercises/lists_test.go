package exercises

import "fmt"

func ExampleCountOdd() {
	fmt.Println(CountOdd([]int{4, 8, 15, 16, 23, 42}))
	fmt.Println(CountOdd([]int{}))

	// Output:
	// 2
	// 0
}

func ExampleCountOdd_negative() {
	fmt.Println(CountOdd([]int{-3, -2, -1, 0, 1}))

	// Output:
	// 3
}

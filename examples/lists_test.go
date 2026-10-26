package examples

import "fmt"

func ExamplePrintList() {
	PrintList([]int{4, 8, 15})

	// Output:
	// 4
	// 8
	// 15
}

func ExamplePrintListWithIndex() {
	PrintListWithIndex([]int{4, 8, 15})

	// Output:
	// 0 4
	// 1 8
	// 2 15
}

func ExampleCountEven() {
	fmt.Println(CountEven([]int{4, 8, 15, 16, 23, 42}))
	fmt.Println(CountEven([]int{}))

	// Output:
	// 4
	// 0
}

func ExampleSum() {
	fmt.Println(Sum([]int{4, 8, 15}))
	fmt.Println(Sum([]int{}))

	// Output:
	// 27
	// 0
}

func ExampleTotalLength() {
	fmt.Println(TotalLength([]string{"go", "for", "range"}))
	fmt.Println(TotalLength([]string{}))

	// Output:
	// 10
	// 0
}

func ExampleContains() {
	fmt.Println(Contains([]int{4, 8, 15}, 8))
	fmt.Println(Contains([]int{4, 8, 15}, 16))

	// Output:
	// true
	// false
}

func ExampleMax() {
	fmt.Println(Max([]int{4, 15, 8}))
	fmt.Println(Max([]int{-4, -15, -8}))
	fmt.Println(Max([]int{}))

	// Output:
	// 15
	// -4
	// 0
}

func ExampleCountIncreases() {
	fmt.Println(CountIncreases([]int{1, 3, 2, 5, 5, 7}))
	fmt.Println(CountIncreases([]int{1}))

	// Output:
	// 3
	// 0
}

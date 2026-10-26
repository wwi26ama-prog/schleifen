package examples

import "fmt"

func ExampleRangeLoop() {
	RangeLoop(3)

	// Output:
	// 0
	// 1
	// 2
}

func ExampleEvenNumbersRange1() {
	EvenNumbersRange1(4)

	// Output:
	// 0
	// 2
}

func ExampleEvenNumbersRange2() {
	EvenNumbersRange2(4)

	// Output:
	// 0
	// 2
}

func ExampleSumNumbersRange() {
	fmt.Println(SumNumbersRange(4))

	// Output:
	// 6
}

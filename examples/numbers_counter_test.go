package examples

import "fmt"

func ExampleLoopWithCounter() {
	LoopWithCounter(3)

	// Output:
	// 0
	// 1
	// 2
}

func ExampleEvenNumbersCounter1() {
	EvenNumbersCounter1(4)

	// Output:
	// 0
	// 2
}

func ExampleEvenNumbersCounter2() {
	EvenNumbersCounter2(4)

	// Output:
	// 0
	// 2
}

func ExampleSumNumbersCounter() {
	fmt.Println(SumNumbersCounter(4))

	// Output:
	// 6
}

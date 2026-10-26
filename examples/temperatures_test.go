package examples

import "fmt"

func ExamplePrintTemperatures() {
	PrintTemperatures([]int{-2, 3, 1})

	// Output:
	// day 1: -2°C
	// day 2: 3°C
	// day 3: 1°C
}

func ExampleCountFrostDays() {
	fmt.Println(CountFrostDays([]int{-2, 3, 0, -5, 1}))
	fmt.Println(CountFrostDays([]int{}))

	// Output:
	// 2
	// 0
}

func ExampleAverageTemperature() {
	fmt.Println(AverageTemperature([]int{-2, 3, 0, -5, 1}))
	fmt.Println(AverageTemperature([]int{20, 21}))
	fmt.Println(AverageTemperature([]int{}))

	// Output:
	// -0.6
	// 20.5
	// 0
}

func ExampleWarmestDay() {
	fmt.Println(WarmestDay([]int{-2, 3, 0, 3, 1}))
	fmt.Println(WarmestDay([]int{-5, -2}))
	fmt.Println(WarmestDay([]int{}))

	// Output:
	// 1
	// 1
	// -1
}

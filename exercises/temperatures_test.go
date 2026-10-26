package exercises

import "fmt"

func ExampleCountColderDays() {
	fmt.Println(CountColderDays([]int{5, 3, 4, 4, -1, 2}))
	fmt.Println(CountColderDays([]int{5}))
	fmt.Println(CountColderDays([]int{}))

	// Output:
	// 2
	// 0
	// 0
}

func ExampleTemperatureSpread() {
	fmt.Println(TemperatureSpread([]int{5, 3, 4, -1, 2}))
	fmt.Println(TemperatureSpread([]int{-3, -8}))
	fmt.Println(TemperatureSpread([]int{7}))
	fmt.Println(TemperatureSpread([]int{}))

	// Output:
	// 6
	// 5
	// 0
	// 0
}

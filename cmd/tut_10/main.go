// generics
package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type engine[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var intSlice = []int{1, 2, 3, 4}
	var floatSlice = []float32{1, 3, 4, 5, 3}
	fmt.Println(sumint(intSlice))
	fmt.Println(sumfloat(floatSlice))
	fmt.Println(sumSlice[int](intSlice)) // we can remove [int], but in some places go cant infer the type of the slice that we pass

	var myEngine engine[gasEngine] = engine[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 50,
			mpg:     30,
		},
	}

	fmt.Println(myEngine)
}

func sumint(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum
}

// if we want to find the sum of float32
func sumfloat(mySlice []float32) float32 {
	var sum float32 = 0
	for _, v := range mySlice {
		sum += v
	}
	return sum
}

// to solve this issue we make use of generics
// like templates in cpp
func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// we also have the 'any' type for generics/
// but we cant use 'any' everywhere for eg in the above example as we cant add bool
// we can use any in this case >
func checkEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

package main

import "fmt"

// Number is an int64 or float64
type Number interface {
	int64 | float64
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// Sum takes a comparable K and V int64 or float64
func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Map applies a function to every element of a slice
func Map[A, B any](slice []A, f func(A) B) []B {
	result := make([]B, len(slice))
	for i, e := range slice {
		result[i] = f(e)
	}
	return result
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		Sum(ints),
		Sum(floats))

	sliceInts := []int{1, 2, 3}
	sliceStrings := []string{"1", "2", "3"}

	fmt.Println(Map(sliceInts, func(a int) int {
		return a * 2
	}))

	fmt.Println(Map(sliceStrings, func(a string) string {
		return a + "!"
	}))
}

package main

import "fmt"

func main() {
	s := []string{"apple", "banana", "mango"}
	i := []int{111, 222, 333}

	genericPrintFunction[string](s)
	genericPrintFunction[int](i)
}

func genericPrintFunction[T any](t []T) {
	for i, v := range t {
		fmt.Printf("Index %v contains value %v\n", i, v)
	}
}

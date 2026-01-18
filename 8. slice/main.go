package main

import "fmt"

func main() {
	var slice = make([]int, 5, 7)
	fmt.Printf("len(slice): %v\n", len(slice))
	fmt.Printf("cap(slice): %v\n", cap(slice))
	fmt.Printf("address: %p\n", slice)
	fmt.Printf("address: %p\n", &slice)

	fmt.Printf("slice: %v\n", slice)
	slice = append(slice, 6, 7)
	fmt.Printf("len(slice): %v\n", len(slice))
	fmt.Printf("cap(slice): %v\n", cap(slice))
	fmt.Printf("address: %p\n", slice)
	fmt.Printf("address: %p\n", &slice)
	fmt.Printf("slice: %v\n", slice)

	slice = append(slice, 8)
	fmt.Printf("len(slice): %v\n", len(slice))
	fmt.Printf("cap(slice): %v\n", cap(slice))
	fmt.Printf("address: %p\n", slice)
	fmt.Printf("address: %p\n", &slice)
	fmt.Printf("slice: %v\n", slice)

	// new cap = old cap + (oldcap + 256 * 3) / 4
}

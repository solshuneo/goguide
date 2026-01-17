package main

import "fmt"

func main() {
	// array: finite element & same datatype
	// 1 2 3 4 5
	// "AAA", "BBB", "CCC"
	// true true true true false false false

	// var name type = expression
	// var a [5]int = [5]int{1, 2, 3, 4, 5} //brace
	// // fmt.Println(a)
	// // fmt.Println("a[0]", a[0])
	// // fmt.Println("a[1]", a[1])
	// // fmt.Println("a[2]", a[2])
	// // fmt.Println("a[3]", a[3])
	// // fmt.Println("a[4]", a[4])

	// a[1] = 10
	// fmt.Println(a)
	// fmt.Println("a[1]", a[1])
	// decleration
	// var emptyArray = [7]int{}
	// fmt.Println(emptyArray)
	// var intArray = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(intArray)
	// var lackArray = [5]int{1, 2, 3}
	// fmt.Println(lackArray)
	var fixedArray = [5]int{1: 2, 4: 1} // 0, 2, 3 (zero value)

	fmt.Println(fixedArray)
	var lengthArray = len(fixedArray)
	fmt.Println(lengthArray)
	// // length of array
	// fmt.Println(len(emptyArray))
	// fmt.Println(len(intArray))
	// fmt.Println(len(lackArray))
	// fmt.Println(len(fixedArray))

}

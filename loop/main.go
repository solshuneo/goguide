package main

import "fmt"

func main() {
	// standard for loop
	// var a = [5]int{1, 2, 3, 4}
	// var length = 5
	// for i := 0; i < length; i += 1 {
	// 	fmt.Println(a[i])
	// }
	// for while
	// var x = 10
	// for x != 0 {
	// 	fmt.Println(x)

	// 	x -= 1
	// }
	// for x := 10; x != 0; x -= 1 {
	// 	fmt.Println(x)

	// }
	// infinite loop
	for {
		fmt.Println("infinite loop")
		break
	}
	// // for range
	// var a = [5]int{1, 2, 3, 4, 5}
	// for index, value := range a {
	// 	fmt.Printf("Vị trí %d là giá trị %d\n", index, value)
	// }
}

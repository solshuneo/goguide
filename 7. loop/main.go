package main

import "fmt"

func main() {
	// var a = [4]int{10, 20, 30, 40}
	// var length = 4
	// var i = 0 //index
	// for i < length {
	// 	fmt.Println(a[i])
	// 	i += 1
	// }

	// for i := 0; i < length; i += 1 {
	// 	fmt.Println(a[i])
	// }

	// }
	// infinite loop
	// var i = 1
	// var limit = 10
	// for {
	// 	fmt.Println(i, "infinite loop")
	// 	if i < limit {
	// 		i += 1
	// 		continue
	// 	}
	// 	break
	// }
	// for range
	var a = [5]int{1, 2, 3, 4, 5}
	for index, value := range a {
		fmt.Printf("Vị trí %d là giá trị %d\n", index, value)
	}
	for i := 1; i < 5; i += 1 {
		var index = i
		var value = a[i]
		fmt.Printf("Vị trí %d là giá trị %d\n", index, value)
	}
}

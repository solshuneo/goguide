package main

import "fmt"

func main() {
	fmt.Printf("fibo(5): %v\n", fibo(5))
	// ham khong ten, anonymous function
	var fibo func(int) int
	fibo = func(n int) int {
		return fibo(n-1) + fibo(n-2)
	}
}

func fibo(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

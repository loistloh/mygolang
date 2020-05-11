package main

import "fmt"

func main() {
	sample1()
	fmt.Println("=====================")
	sample2()
	fmt.Println("=====================")
}

func sample1() {
	var i int = 30
	fmt.Printf("%d 的阶乘是 %d\n", i, factorial(uint64(i)))
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func sample2() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Println()
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

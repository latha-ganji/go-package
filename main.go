package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(add(10, 20))
}

func add(a, b int) int {
	return a + b
}

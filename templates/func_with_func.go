package main

import "fmt"

func main() {
	fmt.Println("call func with a func")
	funcWithFunchParameter(someFunc)
}

// argument func must have signature so the passed func can be validated on compile time
func funcWithFunchParameter(someFunc func(int) int) {
	fmt.Println("func that use func that is passed")
	someFunc(0)
}

//
func someFunc(int) int {
	fmt.Println("the func it self")
	return 0
}

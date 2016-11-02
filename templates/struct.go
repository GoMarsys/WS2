package main

import "fmt"

type ExampleStruct struct {
	ExportedValue   int
	unexportedValue int
}

// GetterForUnexportedValue can be used from outside of the package
func (es ExampleStruct) GetterForUnexportedValue() int {
	return es.unexportedValue
}

func main() {

	// Creating a new object/struct
	value := ExampleStruct{
		ExportedValue: 1,
		unexportedValue: 2
	}

	// this is a Exported value that can be received from anywhere
	fmt.Println(value.ExportedValue)

	// this is a method that return value under the unexported value
	fmt.Println(value.GetterForUnexportedValue())

	// this will raise error from outside of the package space
	fmt.Println(value.unexportedValue)

}

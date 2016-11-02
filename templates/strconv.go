package main

import (
	"fmt"
	"strconv"
)

func main() {
	StrToInt()
	IntToStr()

	StrToFloat()
	FloatToStr()
}

func IntToStr() {
	i := strconv.Itoa(42)
	fmt.Printf("%T: %v\n", i, i)
}

func StrToInt() {

	i, e := strconv.Atoi("42")

	if e != nil {
		panic("OMG!")
	}

	fmt.Printf("%T: %v\n", i, i)

}

func StrToFloat() {

	i, e := strconv.ParseFloat("42.5", 64)

	if e != nil {
		panic("OMG!")
	}

	fmt.Printf("%T: %v\n", i, i)

}

func FloatToStr() {
	i := strconv.FormatFloat(42.42, 'E', -1, 64)
	fmt.Printf("%T: %v\n", i, i)
}

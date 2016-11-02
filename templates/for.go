package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 10 {
			break
		}

		fmt.Println(i)

		i++
	}
}

package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello World", uuid.NewString())

	a := 1
	if a > 0 {
		fmt.Println("A is greater than 0")
	} else {
		fmt.Println("A is less than 0")
	}

	for i := 0; i < 10; i++ {
		fmt.Println("i is", i)
	}

	i := []int{1, 2, 3, 4, 5}
	for index, value := range i {
		fmt.Println("The index and value are: ", index, value)
	}
}
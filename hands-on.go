package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type secretAgents struct {
	person
	salary string
}

func main() {
	x1 := make([]int, 50)
	fmt.Println(x1)
	x2 := make([]int, 0, 50)
	fmt.Println(x2)

	am := map[string]int{
		"tod":   42,
		"Henry": 51,
		"Ben":   39,
	}

	fmt.Println(am)

	for key, value := range am {
		fmt.Println(key, value)
	}

	delete(am, "tod")

	for key, value := range am {
		fmt.Println(key, value)
	}

	p1 := person{
		name: "Syad",
		age:  23,
	}

	fmt.Println(p1)

	p2 := secretAgents{
		person: person{
			name: "Uhuy",
			age:  25,
		},
		salary: "1000000",
	}

	fmt.Println(p2)

	p3 := struct {
		name string
		age  int
	}{
		name: "James",
		age:  22,
	}

	fmt.Println(p3)
}

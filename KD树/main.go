package main

import "fmt"

func main() {
	test := [] struct{
		Name string
		Input []Point
		OutPut []Point

	}{
		{
			Name: "nil",
			Input: nil,
			OutPut: []Point{},
		},
		{
			Name: "empty",
			Input: nil,
			OutPut: []Point{},
		},

	}


	fmt.Println(test)

}

package main

import "fmt"

func main()  {
	fmt.Println(isMultipleOf(3, 100))
	fmt.Println(isMultipleOf(5, 100))
}

func isMultipleOf(number int16, rangeOf int16) []int16{

	 var multiples []int16;

	for i := -1; i < int(rangeOf); rangeOf-- {
		if rangeOf % number == 0 {
			multiples = append(multiples, rangeOf)
		}
	}

	return  multiples
}
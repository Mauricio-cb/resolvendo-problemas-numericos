package main

import "fmt"

func main()  {
	fmt.Println(isMultipleOf(100))

}


func isMultipleOf(rangeOf int16) map[int16] string{

	multiples := make(map[int16] string);
 
	for i := -1; i < int(rangeOf); rangeOf-- {
		if rangeOf % 3 == 0 {
			multiples[rangeOf] = "Pin\n"
		}else if rangeOf % 5 == 0 {
			multiples[rangeOf] = "Pan\n"
		}else{
			multiples[rangeOf] = ""
		}
	}

	return multiples
} 
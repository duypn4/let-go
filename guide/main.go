package main

import (
	"fmt"
	"strconv"

	"guide.com/guide/person"
)

func main() {
	// number := 1

	// fmt.Println(&number)
	// fmt.Println(number)

	// display(&number)
	// fmt.Println("original number", number)
	person := person.New("Quyet")
	person.SetName("Duy")
	fmt.Println(person)

	numberString := "1"
	number, err := strconv.ParseInt(numberString, 10, 64)

	if err != nil {
		panic("invalid number")
	}

	fmt.Println(number, 1, 2, 3)
}

// func display(number *int) {
// 	*number++
// 	fmt.Println("function value", *number)
// }

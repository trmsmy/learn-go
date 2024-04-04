//create basic main function

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Starting")

	var str string = "resume"
	fmt.Println(str)

	var str2 string = "Résumé"
	fmt.Println(str2)

	//iterate the str2  -- iterates over array representation
	for i := 0; i < len(str2); i++ {
		var char = str2[i]
		fmt.Printf("%v %v %c \n", i, char, char)
	}

	//iterate over string2 -- iterates over string representation
	//you won't see 2nd place in the iteration output
	for i, v := range str2 {
		fmt.Printf("%v %v %c \n", i, v, v)
	}

	var indexed = str2[2]
	fmt.Printf("%c %v %T\n", indexed, indexed, indexed)

	fmt.Println("Length of string representation: ", len(str2))

	//-------------- rune to rule the world

	var resume = []rune(str2)

	//iterate through resume
	for i, v := range resume {
		fmt.Printf("%v %v %c \n", i, v, v)
	}

	fmt.Println("Length of rune representation: ", len(resume))

	//--------------------------------

	//string builder example
	var builder = strings.Builder{}
	builder.WriteString(str2)
	fmt.Println(builder.String())
	fmt.Println("Length of string representation: ", len(builder.String()))
	fmt.Println("Length of string representation: ", len(str2))

}

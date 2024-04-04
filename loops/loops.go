package main

import "fmt"

func main() {
	var myMap = make(map[string]string)
	myMap["key1"] = "value1"
	myMap["key2"] = "value2"
	myMap["key3"] = "value3"
	myMap["key4"] = "value4"

	fmt.Println("Hello, world!", myMap)

	//iterate the map and print the keys and values
	for key, value := range myMap {
		fmt.Printf("Key: %v, Value: %v \n", key, value)
	}

	//create an array of string
	var myArray = []string{}
	myArray = append(myArray, "Hello", "World")
	myArray = append(myArray, "!")
	myArray = append(myArray, "!")
	fmt.Println(myArray)

	//iterate the array and print the values
	for _, value := range myArray {
		fmt.Printf("Value: %v  \n", value)
	}

	fmt.Println(len(myArray))
	fmt.Println(cap(myArray))

}

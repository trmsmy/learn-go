package main

import "fmt"

func main() {
	fmt.Println("datastructure ..")

	var myarr [3]int32 = [3]int32{1, 2, 3}
	var myarr1 [3]int32 = [...]int32{1, 2, 3}

	myarr2 := [...]int32{1, 2, 3}

	fmt.Println(myarr)
	fmt.Println(myarr1)
	fmt.Println(myarr2)

	var splice []int32 = []int32{}
	newsplice := append(splice, 2)

	fmt.Printf("Splice: %v", newsplice)

}

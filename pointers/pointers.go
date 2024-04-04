//create basic main function

package main

import "fmt"

func main() {
	fmt.Println("Starting")

	var n int32 = 23
	var np *int32 = &n

	var emptyp *int32 = new(int32)

	fmt.Printf("n ==> value=%v and pointer=%v \n", n, np)

	fmt.Printf("emptyp ==> value=%v and pointer=%v \n", *emptyp, emptyp)

	*emptyp = *np
	fmt.Printf("emptyp ==> value=%v and pointer=%v \n", *emptyp, emptyp)

	*emptyp = 2
	fmt.Printf("emptyp ==> value=%v and pointer=%v \n", *emptyp, emptyp)

	//emptyp pointer should be same as np pointer
	emptyp = &n
	fmt.Printf("emptyp ==> value=%v and pointer=%v \n", *emptyp, emptyp)

	//----------------------------------------------------------------

	pointers_test()
}

func pointers_test() {

	var str = "hello world"
	fmt.Printf("str value=%v and its pointer=%p \n", str, &str)

	append_num(str)
	fmt.Println("str  value after append by value:", str)
	append_num_using_pointer(&str)
	fmt.Println("str  value after append by pointer:", str)

}

func append_num(str string) string {
	fmt.Println("pointer in func:", &str)

	str = str + "123"
	return str
}

func append_num_using_pointer(str *string) string {
	*str += "123"
	fmt.Println("pointer in func:", str)
	return *str
}

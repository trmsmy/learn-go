//create main function

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var results = []string{}
var wg = sync.WaitGroup{}
var m = sync.Mutex{}

func main() {

	t0 := time.Now()
	fmt.Println("About to save 10 user data...")
	var users = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, userId := range users {
		wg.Add(1)
		go insertToDatabase(userId)
	}
	wg.Wait()
	fmt.Printf("Time taken to save 10 user data: %v\n", time.Since(t0))
	fmt.Printf("results ==> %v, length=%v", results, len(results))

}

func insertToDatabase(userId int) string {
	defer wg.Done()
	//fmt.Println("About to save user data...", userId)
	var result = insert(userId)
	fmt.Printf("Saved user data...id=%v  value=%v \n", userId, result)
	return result
}

func insert(id int) string {
	time.Sleep(1 * time.Second)
	var result = "id" + strconv.Itoa(id)
	m.Lock()
	results = append(results, result)
	m.Unlock()
	return result
}

func readAll() []string {
	return []string{}
}

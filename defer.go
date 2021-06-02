package main

import "fmt"

//explainDefer explains working of defer statement, defer schedules a function call to be run after the function completes.
//Deferred functions are executed in LIFO order.
func explainDefer() {
	defer fmt.Println("second_0")
	defer fmt.Println("second_1")
	defer fmt.Println("second_2")
	defer fmt.Println("second_3")
	fmt.Println("first_0")
	fmt.Println("first_1")
	fmt.Println("first_2")
	fmt.Println("first_3")
}

func main() {
	explainDefer()
}

// OUTPUT:-
// first_0
// first_1
// first_2
// first_3
// second_3
// second_2
// second_1
// second_0

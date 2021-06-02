package main

import "fmt"

func hands_on_1() {
	x := 42
	y := "James_bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}


var x int
var y string
var z bool

func hands_on_2() {
	fmt.Println(x,y,z)
}


// these variables are assigned as package level scope
var a int = 42
var b string = "james_bond"
var c bool = true

func hands_on_3(){
	s := fmt.Sprintf("%v %v %v", a, b, c)
	fmt.Println(s)

}

type hotdog int
var abc hotdog

func hands_on_4(){

	fmt.Println(abc)
	fmt.Printf("%T",abc)	
}

var d hotdog 
var e int

func hands_on_5(){

	e = int(d)
	fmt.Println(e)

}

func main() {

	hands_on_1()

	//getting default values for null variables
	hands_on_2()

	//getting out put in string format
	hands_on_3()

	//desclaring underlying type variable
	//In this a variable is decraled with type int or string, and then another variable can be assign with same data type.
	hands_on_4()

	//type casting 
	hands_on_5()
}

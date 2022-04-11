package main

import "fmt"

func main()  {
	func(name string) {//Closure adalah fungsi dalam fungsi, IIFE method
		fmt.Println(name)
	}("Hacktiv8")


		test := greet("hello", 20)
		test("Hello")

}


func greet(names string, age int) func(name string) {

	return func(name string) {
		fmt.Println(name, age)
	}

}
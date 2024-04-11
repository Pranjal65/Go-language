package main

import "fmt"

// we can use this variables and functions in other file directly with same package name
// if there are multiple file withs same package name they are considered as part of same package and we cann access variables from any file

// to access variables or functions all the files with same pakage name should be running     #go run main.go package.go

var qr = []int{1,2,3,4,5,6};
// variables which are outside any function are accessible in any other file otherwise it will not be accessed by other files due to its local scope
func sayHellofunc(s string){
	fmt.Println("Hellooo ",s);
}
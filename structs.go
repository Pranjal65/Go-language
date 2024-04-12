package main

import "fmt"

// structure vertex: type and  struct keywords are used to define struct in go
type Vertex struct {
	X int
	Y int
}

func main() {
	// create  object of each struct or variable for struct and the acess struct elements by (.).
     var v Vertex;
	 v.X=1;
	 v.Y=2;
	 // we can pass struct as a paramater to the function
	 // function(struct_var struct_name)
	 
	 // Pointers to structs
	point:= &v;
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(point.X)
	fmt.Println(point.Y)
}

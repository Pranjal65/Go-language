package main

import (
	"fmt"
)

func main(){
	fmt.Println("********Loops**********")

	// while loop :In Go language for keyword is used for while loop
	x:=0
	for x<5{
		fmt.Println("this is whhile loop");
		x++;
	}

	//for loop

	for i:=0 ; i<5; i++{
        fmt.Println("this is for loop");
	}

	//Range keyword

	arr:= []string {"asd","kjh","tye","fghjk"}

	// here index and value are the local copy of variable change in index or values does not reflect in original array
	for index,value := range arr{
		fmt.Printf("the position of %v is %v \n",value,index);
	}

	// if we dont want any of the element like value or index simply replace it with _
	for _,value := range arr{
		fmt.Printf(" %v is array element \n",value);
	}

	for index,_ := range arr{
		fmt.Println(index);
	}

}
package main

import (
	"fmt"
	"sort"
	"strings"
	"cmp"
)

func standard_Library(){
	fmt.Println("************* Standard Libraries ******************");
    // user inputs
	fmt.Println("enter a value: ")
	var k int;
	fmt.Scan(&k);
	fmt.Println("Your answer: ",k);    
	/*The Scanln() function is similar to Scan(), but it stops scanning for inputs at a newline*/
	var i,j int

    fmt.Print("Type two numbers: ")
    fmt.Scanln(&i, &j)
    fmt.Println("Your numbers are:", i, "and", j)

	//The Scanf() function receives the inputs and stores them based on the determined formats for its arguments.
	fmt.Print("Type two numbers: ")
    fmt.Scanf("%v %v",&i, &j)
    fmt.Println("Your numbers are:", i, "and", j)
    
	fmt.Println("********** Strings Package***************")
	
	// strings package
	 // 1] conatins function
	 Greetings:= "Hello everyone";
	 fmt.Println(strings.Contains(Greetings,"everyone"))  // return true or false
     // 2] ReplaeAll(string, value to be replaced, new value) function
	 // this function does not actually changes the actual string 
	 fmt.Println(strings.ReplaceAll(Greetings,"Hello","Hiii"))// return  new string with replaced values string

	 // 3] ToUpper(string) return new string with all letters in uppercase
	 fmt.Println(strings.ToUpper(Greetings))

	 // 4] index(string, string or character ) return starting index of the string to be search
	 fmt.Println(strings.Index(Greetings,"ll"));

	 // 5] Split(string, point by which string to be splitted) return new array of strings
	 fmt.Println("Splitted string: " ,strings.Split(Greetings," "))

	 //sort package

	 fmt.Println("*********** Sort Package ***********")
    
	 age := []int {23, 56,78,12, 4, 5,3 ,2,89}

	 age1 := []float32 {23, 56,78,12, 4, 5,3 ,2,89}
	 fmt.Println(age);
	 fmt.Println(age1);
    // sort function 
     sort.Ints(age);
	 fmt.Println("sorted array is: ",age);


	 // get index of array element
	 index:= sort.SearchInts(age,89);
	 fmt.Println(index)

	 //sorting string arra
	 strarr:= []string {"asd","kjh","tye","fghjk"}
    
	  fmt.Println("unsorted string: ",strarr);
	  sort.Strings(strarr);
	  fmt.Println("sorted string: ",strarr);


	  fmt.Println("index of  tye string is : ", sort.SearchStrings(strarr,"tye"))

     fmt.Println("************* CMP package ************")
	  // cmp package
	  //   -1 if x is less than y,
      //    0 if x equals y,
      //    +1 if x is greater than y.
	  userInput1 := ""
	  userInput2 := "some text"

	  fmt.Println(cmp.Or(userInput1, "default"))
	  fmt.Println(cmp.Or(userInput2, "default"))
	  fmt.Println(cmp.Or(userInput1, userInput2, "default"))



   
}
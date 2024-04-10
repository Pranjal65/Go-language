package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	fmt.Println("************* Standard Libraries ******************");
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

      
}
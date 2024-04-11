package main

import (
	"fmt"
	"strings"
)

// func keyword is used to declare the  function  there no return type for function
// function paremeters as (variable name data type) eg (n int)
// return type after function declaration
// we can function as parameter to function

func add(a int, b int) int {
	ans:= a+b;
	fmt.Println(ans);
	return ans
}

func sum(a int, b int) {
	ans:= a+b;
	fmt.Println(ans);
	//return ans
}
// function passed as parametr to another function should not return any value
func  exmpl(f func(int,int)){
	x:=1
	y:=1
	for (y<10){
		 f(x,y);
		y++;
	}
}

// functions with multiple return types

func initials(s string) (string ,string){
     x:=strings.ToUpper(s);
	 a:=strings.Split(x," ");

	 var ans []string
	 for _, val := range a{
		ans=append(ans,val[0:1]);
	 }
	 if(len(ans)>1){
		return ans[0] ,ans[1]
	 }
	 return ans[0], "_"


}
func main(){
	fmt.Println(add(2,3));
    exmpl(sum)
	fn,sn:=initials("pranjal more");
	fmt.Println(fn,sn);
}
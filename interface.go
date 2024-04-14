package main 

import "fmt"
type voice interface{
	Say() string
}

type dog struct {
}
type cat struct{
}

func (c cat) Say() string{
	return "Heyy I am cat";
}
func (d dog) Say() string{
	return "Heyy I am dog";
}

func interfaces(){
	// we use interface as the intermediate between multiple structures or data types
	 fmt.Println("******** interface  *************")
	 var c cat
	 var d dog
	 // here the multiple struct objects are placed in array of interface
	v:=[]voice{c,d}
	for _ ,val:=range v{
		fmt.Println(val.Say());
	}


}



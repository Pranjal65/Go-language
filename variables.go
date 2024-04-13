package main

import "fmt"

func variables(){
	fmt.Println("***************Variables***************")
     // strings
	 var nameone string ="Pranjal"
	 var nametwo  ="Pranjal"
	 fmt.Println(nameone,nametwo)


	 //ints

	 var ageone int =30
	 var agetwo = 20
	 agethree:= 50   //accessed only inside the functions
	 fmt.Println(ageone,agetwo,agethree)
     
	 //long intergs are stored in datatype such as int8, int 26. int 32
	 var longint int8=22
	 fmt.Println(longint);


	 /// array variables

	 var arr [3]int = [3] int {1,2,3};
	 fmt.Println(arr);
	 // dyanamic array
	 var arr2  = [] int {1,2,3};
	 // append function does not  change the original array rather than it return new slice/array with new element appended
	 arr2=append(arr2,5)
	 fmt.Println(arr2[2]);
	 fmt.Println(arr2, " length of arr2 is:  ",len(arr2));

  //slice pof an array2 same as python
  // here the elements added will be from 0-1 excluding 2 and including lower range
    k :=arr2[:2]
   fmt.Println(k);


}

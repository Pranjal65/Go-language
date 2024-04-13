package main

import (
	"fmt"
	//"string"
	//"strings"
)

func maps(){
    fmt.Println("********** Map **********");

	// declaration of map
	menu := map[string]float64{
		"soup":4.999,
		"Samosa":10.888,
		"juice":45.999,
	}
	fmt.Println(menu);
	fmt.Println(menu["juice"]);
	for key,val:= range menu{
		fmt.Printf("\n Price of %v is %v", key, val);
	}

	// int as key type
	phonebook:= map[int]string{
		12345345:"Pranjal",
		56789453:"Pappa",
		10111213:"Didi",
	}
    fmt.Println()
	fmt.Println(phonebook);

	// update value
	phonebook[10111213]="Dii";
	fmt.Println(phonebook);


	// add value
	phonebook[95611325]="Mom";
	fmt.Println(phonebook);

	// delete value
	delete(phonebook,10111213);
	fmt.Println(phonebook);

	// check whether key is present in map or not
	check:=phonebook[10111213] //return 0 if not present else return value of key
	fmt.Println(check);

	phonebook[10111213]="Dii";
	fmt.Println(phonebook);

	check1:=phonebook[810111213] //return 0 if not present else return value of key
	fmt.Println(check1);

	// length of map
	fmt.Println("length of map: ",len(phonebook))

	// initialize empty map
	mp:= map[int]string{}
   str:="A"
	for i:=1;i<=5;i++{
		mp[i]=str;
		str +="A";
	}
	fmt.Println(mp);

	
}

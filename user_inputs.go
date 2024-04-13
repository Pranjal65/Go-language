package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct{
	name string
	entity string
	quantity int
    price int
	total int
}

func createBills() bill{
	var b bill;

	// to read the input we create the object of reader class of bufio package/library and pass os.stdin object to the constructor

	reader:=bufio.NewReader(os.Stdin);
	fmt.Println("enetr the name of customer :");
	name , _:=reader.ReadString('\n')
	name= strings.TrimSpace(name)
	b.name=name;

	fmt.Println("enetr the  entity purchased: ")
	entity,_:= reader.ReadString('\n') 

	b.entity=entity

	fmt.Println("enter the quantity of entity: ")
	var quantity int

	// used to take int inputs
	fmt.Scan(&quantity);
    b.quantity=quantity

	fmt.Println("enter the price of entity: ")
	var price int
	fmt.Scan(&price)
	b.price=price

	b.total=quantity*price;

	return b
} 

func inputs(){

	fmt.Println()
	fmt.Println("********** User Inputs *******")
	b:=createBills();
	fmt.Println(b);
	b.file_manage();
}
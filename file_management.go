package main

import (
	"fmt"
	"os"
)

func format(b *bill) string {
	fb := "***** bill ****"
	// function Sprintf to format the string
	fb += fmt.Sprintf("\nname : %v", b.name)
	fb += fmt.Sprintf("\nentity : %v", b.entity)
	fb += fmt.Sprintf("\nquantity : %v", b.quantity)
	fb += fmt.Sprintf("\nprice : %v", b.price)
	fb += fmt.Sprintf("\ntotal : %v", b.total)
	return fb
}

//function is called in structs
func (b *bill) file_manage() {

	data := []byte(format(b))
	// writefile is used to write data into file
	err := os.WriteFile("./"+b.name+".text", data, 0644)
	if err != nil {
		// panic function is used to break the flow of program when error ocurred
		panic(err)
	}


}

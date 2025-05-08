package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	//scan method 
	
	// fmt.Print("Enter Name : ")
	// var name string 
	// fmt.Scan(&name)
	// fmt.Println(" Hey ,",name)




	// but this can't read a whole string 
	// for that we will use bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter a string :")
	var s string
	s, _ = reader.ReadString('\n')
	fmt.Print("ur string :",s)

}
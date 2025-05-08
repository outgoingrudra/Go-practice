package main

import "fmt"

func main() {
	name := "Rudra Verma"
	age := 20
	val := 23.124
	//Println method
	fmt.Println("Name =",name)
	fmt.Println("age =",age)

	// using printf 
	fmt.Printf("Name : %s , his age : %d\n",name,age)
	fmt.Printf(" value : %.2f",val)

}

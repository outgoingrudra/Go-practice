package main

import "fmt"

func main() {
	fmt.Println("Enter ur age :")
	var age int
	fmt.Scan(&age)

	if age < 0  || age > 150{
		fmt.Println("Some Errors !")
	} else if age < 18 {
		fmt.Println("Child !")
	} else if age < 35 {
		fmt.Println("Young ! ")
	} else {
		fmt.Print("😉")
	}



}
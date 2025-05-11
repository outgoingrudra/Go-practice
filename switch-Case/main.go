package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default :
	  fmt.Print(" Not Sure about day !")
	}
	

	val := 10

	switch val {
	case 10 , 30 , 20  :
		  fmt.Println(" 10 or 20 or 30 ")
	default :
	fmt.Println(" Not those Numbers")

	}
}
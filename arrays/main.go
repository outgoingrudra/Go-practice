package main

import "fmt"

func main() {
	var names [5]string // declaration
	names[0] = "Rudra"
	names[1] = "Virat"
	fmt.Println("Names :",names)

	// declaration and initialization 
	var nums =[10]int{12,23,12,34,11,4,5}
	fmt.Println("numbers = ",nums)
    
	// for array length
	fmt.Println("length of nums =",len(nums))

	// %q use
	var solos [10]string 
	fmt.Printf("auto initialization is go : %q",solos)
}

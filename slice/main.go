package main

import "fmt"

func main() {
	// slices
	numbers := []int{23, 11, 34, 45, 24}
	fmt.Println("Numbers : ",numbers)
	

	//for appending 
	numbers = append(numbers, 100,200,300)
	fmt.Println("After Appending :",numbers)

	// for length 
    fmt.Println("length of nums : ",len(numbers))
	// for capacity
	fmt.Println("capacity of numbers :",cap(numbers))


	//for empty slice
	name :=[]string{}
	fmt.Println(" Empty slice : ",name)

	//creating slice using make function
	elements := make([]int , 3,10)
	elements = append(elements, 23,45)
	fmt.Println("Elements :",elements)
}


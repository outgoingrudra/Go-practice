package main

import "fmt"

func main() {
     // general
	for i := 0; i <= 10; i++ {
		fmt.Println("Number :",i)
	}

	// for while loop
	counter := 0
	for {
		fmt.Println(" While Loop Instance ")
		if counter == 5 {
			break
		}
		counter++
	}

	// for arrays 
	nums := []int{23,14,235,67,234,789}

	for idx , val := range nums{
		fmt.Println("Index=",idx,"Value=",val)
	}



	// for string
	sen := " HEY GUYS !! "
	fmt.Println("------------")

	for idx , c := range sen{
		fmt.Printf("Index= %d and character = %q \n",idx,c)
	}


}
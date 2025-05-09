package main

import "fmt"

func Greet(){
	fmt.Println("Hello ! i am from Greet Function ! ")
}

func sumy(a, b int) int {
	return a + b
}

func DoubleRetrun(a int , b string) (int , string){
	return a*10 , b
}



func main() {

	fmt.Println("sum =",sumy(23,123))

	Greet()

	val , sentence := DoubleRetrun(12 , " Moon-Glancer")

	fmt.Println("value :",val,"String :",sentence)

}

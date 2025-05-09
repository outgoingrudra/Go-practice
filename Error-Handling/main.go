package main

import "fmt"
// using Errorf
func divide(a,b float32) ( float32 , error ){
	if (b==0){
		return 0 , fmt.Errorf("Denominator can't be zero")
	}
	return a/b , nil
}


// u can also return string 
func signum(a int ) (int , string){
	if(a==0){
		return 0 , " Error Occured !"
	}
	return 1 , "nil" 
}


func main(){

	val , err := divide(23,0)
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(" value : ",val)
	//underscore means ignore that variable
	v ,_ := signum(0)
	fmt.Println(v)
}
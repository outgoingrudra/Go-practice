// map = key-value pair storing data structure

package main

import "fmt"

func main() {
	// using make function
	students := make(map[string]int)

	students["Rudra"] = 75
	students["Asker"] = 12
	students["Prince"] = 45
	students["Priya"] = 54
	students["Ravi"] = 12
	students["Kumar"] = 67

	fmt.Println("Marks of Rudra =", students["Rudra"])

	// changing value
	students["Rudra"] = 23
	fmt.Println("Marks of Rudra after changes =", students["Rudra"])

	// delete key
	delete(students, "Asker")
	fmt.Println("Marks of Asker =", students["Asker"]) // Will print 0 since key is deleted

	// print all key-value pairs
	fmt.Println("\nAll student marks:")
	for key, val := range students {
		fmt.Println("Key:", key, "Value:", val)
	}

	// check if key exists
	val, exist := students["Rudra"]
	fmt.Println("\nRudra key exists:", exist, ", its value:", val)

	// another way to create map
	person := map[string]int{
		"abc": 23,
		"xyz": 21,
		"pqr": 72,
	}

	fmt.Println("\nPerson map values:")
	for key, val := range person {
		fmt.Println("Key:", key, "Value:", val)
	}
}

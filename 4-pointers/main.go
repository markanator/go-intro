package main

import "fmt"

func main() {
	age := 32          // regular
	agePointer := &age // pointer to regular variable

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}

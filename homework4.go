package main

import "fmt"

func main() {
	var myVar int = 42
	var myPointer *int = &myVar
	var mySecondPointer **int = &myPointer

	fmt.Printf("myVar: %d\nmyPointer: %d\nmySecondPointer: %d\n\n", myVar, *myPointer, **mySecondPointer)
	fmt.Printf("myVar: %p\nmyPointer: %p\nmySecondPointer: %p\n\n", &myVar, myPointer, mySecondPointer)

	newValue := 99
	*myPointer = newValue

	fmt.Printf("newValue: %d\nmyPointer: %d\nmySecondPointer: %d\n", myVar, *myPointer, **mySecondPointer)

	fmt.Println("----------------------------------------------------")

	myArray := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array:", myArray)

	var myArrayPointer *[5]int = &myArray

	(*myArrayPointer)[2] = 10

	fmt.Println("Array:", myArray)
}

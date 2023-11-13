package main

import "fmt"

func main() {
	fruits := [3]string{"apple", "banana", "cherry"}
	fmt.Println(fruits)
	fmt.Println("---------------")

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i], " ")
	}

	fmt.Println("---------------")

	var arr [3][4]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

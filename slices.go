package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println(fruits)

	fmt.Println("---------------")

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i], " ")
	}

	fmt.Println("---------------")

	slice := make([][]int, 3)

	for i := 0; i < 3; i++ {
		slice[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			slice[i][j] = 0
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d ", slice[i][j])
		}
		fmt.Println()
	}
}

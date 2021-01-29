package main

import "fmt"

func main() {

	fmt.Println(sum([]int{5, 3, 4, -5, 18}))
}

func sum(numbers []int) int {
	total := 0
	for i := range numbers {
		total = numbers[i] + total
	}
	return total
}

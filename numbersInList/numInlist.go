package main

import "fmt"

func main() {

	fmt.Println(numInList([]int{3, 3, 4, 5}, 5))
}

func numInList(list []int, num int) bool {
	for i := range list {
		if list[i] == num {
			return true
		}
	}
	return false
}

package main

import "fmt"

func printValueFizzBuzz(n int) {
	switch {
	case n%3 == 0:
		fmt.Print("Fizz")
	case n%5 == 0:
		fmt.Print("Buzz")
	case n%3 == 0 && n%5 == 0:
		fmt.Print("Fizz Buzz")
	default:
		fmt.Print(n)
	}
}
func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		printValueFizzBuzz(i)
		fmt.Print(", ")
	}
}

func main() {
	fizzBuzz(6)

}

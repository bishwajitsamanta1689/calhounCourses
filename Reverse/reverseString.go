package main

import "fmt"

func main() {
	fmt.Println(reverse("CAT"))
}

//func reverse(word string)string{
//	var res string
//	for i:= len(word)-1;i>=0;i--{
//		res = res+ string(word[i])
//	}
//	return res
//}

func reverse(word string) string {
	var res string
	for i := 0; i < len(word); i++ {
		res = string(word[i]) + res
	}
	return res
}

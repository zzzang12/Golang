package main

import "fmt"

func main() {
	temp := make([]int, 5)
	fmt.Println(len(temp))
	if len(temp) == 0 {
		fmt.Println("비었넹")
	}
}

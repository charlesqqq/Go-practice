package main

import "fmt"

func main() {
	nums := [11]int{}
	for i := range nums {
		nums[i] = i
	}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}

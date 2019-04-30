package main

import "fmt"

func main() {
	person := map[string]int{
		"Alex":  18,
		"Bob":   26,
		"Candy": 16,
	}
	printMap(person)
}

func printMap(people map[string]int) {
	for name, age := range people {
		fmt.Println("name:", name, "age:", age)
	}
}

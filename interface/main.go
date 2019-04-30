package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	printGreeting(englishBot{})
	printGreeting(spanishBot{})
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (bot englishBot) getGreeting() string {
	return "Hello"
}

func (bot spanishBot) getGreeting() string {
	return "Hola"
}

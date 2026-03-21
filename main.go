package main

import "fmt"

func main() {
	Greating("Sasha")
}

func Greating(name string) {
	fmt.Println("Hello" + name + "!")
}

func SayGoodbye(name string) {
	fmt.Println("Goodbye" + name + "!")
}

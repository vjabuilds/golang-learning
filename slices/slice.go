package main

import "fmt"

func main() {
	greeting := "Hello world!"
	fmt.Println("The whole greeting is", greeting)
	fmt.Println("The first part of the greeting is", greeting[:5])
	fmt.Println("The second part of the greeting is", greeting[5:])
	fmt.Println("The middle part of the greeting is", greeting[2:7])
	greeting = greeting[5:] + greeting[:5]
	fmt.Println("The flipped version of the greeting is", greeting)
}
